# JobManager (Migration Logic)

This PoC is used to come up with a valid design for the new __JobManager__, needed to have a fully automated Kyma deploy possible. It will be used to configure the cluster and the components during the deployment of Kyma. The term `Deployment`/`Deploy` is being used in the context of installing Kyma on an empty cluster or to upgrade Kyma from an older to a newer version.

To achieve a valid solution for the PoC we need to come up with a design for the following:

- [Work-Flow / General mechanism](#work-flow--general-mechanism)
  - [Requirements](#requirements)
  - [Possible Solution](#possible-solution)
- [Draft for Golang Implementation](#draft-for-golang-implementation)
- [Placement of logic and actual jobs](#placement-of-logic-and-actual-jobs)
- [Additions](#additions)

## Work-Flow / General mechanism

### Requirements

- Only support single linear upgrade: A &#8594; B && B &#8594; C; NOT A &#8594; C. This is due to the fact that Kyma only supports single linear upgrades.
- Mechanism should only trigger for one specific Kyma version, this could be: The Kyma version which wants to be installed on an empty cluster, Kyma version to upgrade to, or the Kyma version which should be uninstalled from the cluster.
- This mechanism supports jobs for two different main use-cases: The __component based__ jobs and the __global/non-component based__ jobs
  - __Component based__:
    - Check if component is installed on cluster, or if it wants to be newly installed, and only trigger if yes
    - It should be possible to trigger jobs before and after an deployment of a component
  - __global / Component independent__:
    - Always trigger logic when installing, upgrading, or uninstalling Kyma
    - It should be possible to trigger jobs before and after the deployment, deletion of Kyma
    - We call them `global` jobs to stick to the naming convention of our helm charts

### Possible Solution

To fulfill the requirements a new package, called `JobManager`, will be introduced, which is capable to register, manage, and trigger certain jobs to have a fully automated install/migration/deletion. This package has four (hash)maps to manage the main workload: Two for `pre`-jobs (deploy and deletion) and two for `post`-jobs (deploy and deletion). Key of the maps are the names of the component they belong to, and the value will be a slice of jobs.
Furthermore, it has a `duration` var for benchmarking, and a `targetersion` var to know which jobs should be triggered at a certain deploy.

Jobs will be implemented inside of the `JobManager` package in `go`-files, one for each component, using the specific `job`-interface. Then the implemented interface will be registered using `register(job)` in the same file. This function queues the jobs into a slice, this is due to the reason that we do not know what the value of the targetVersion is until now. 

The JobManager will be used by the `deployment` package in the `deployment.go` and the `deletion.go` file. Inside of the `NewDeployment` or `NewDeletion` functions, the `SetKymaVersion` function of the JobManager will be called to set the targetVersion and thus to build the needed maps. Doing it in this way, we only have the requireed jobs registered in the maps and save some checks later on (&#8594; everything pre-calculated).Then at the hooks, during the deployment/deletion-phase, each hook just has to check if the key for the wanted component is present in the pre/post-map, if yes the jobs in the map will be trigged, if no nothing has to be done.

To benchmark the jobs, a timer will be used in the pre and post job triggers.

Retries for the jobs will not be handled by the JobManager. Retries should be implemented by the jobs themself. This is due to the reason to have a better flexability and a simple to manage interface.

<img src="./migration-logic-diagram.png?raw=true">

## Draft for Golang Implementation

### core.go with main logic 

```go
package jobManager

import (
	"sync"
	"time"
)

type component string
type targetVersion string
type installationType string

type executionTime int

const (
	Pre executionTime = iota
	Post
)

const (
	Deploy    installationType = "deploy"
	Uninstall installationType = "uninstall"
)

var duration time.Duration = 0.00
var kymaVersion targetVersion

var preDeployJobMap map[component][]job
var postDeployJobMap map[component][]job

var preDeletionJobMap map[component][]job
var postDeletionJobMap map[component][]job

var jobs []jobs

// Define type for jobs
type job interface {
	execute() error
	when() (component, targetVersion, executionTime, installationType)
}

func initializeMaps() error {
	for _, job := range jobs {
		// TODO: Add job to corresponding map
	}
}

// Register job
func register(j job) {
	jobs = append(jobs, job)
}

// Gets called in deletion.go and deployment.go when `NewDeletion`/`NewDeployment` is being called
func SetKymaVersion(version targetVersion) {
	kymaVersion = version
	if err := initializeMaps(); err != nil {
		// TODO: handle error
	}
}

// Function should be called before component is being deployed/upgraded
func ExecutePre(component string, it installationType) {
	start := time.Now()
	// TODO: Executes the registered functions for given component; using maps
	//       If map for given key(aka component) is empty, nothing will be done
	//       Check installationType, to know which map should be used
	t := time.Now()
	duration += t.Sub(start)
}

// Function should be called after compoent is being deployed/upgraded
func ExecutePost(component string, it installationType) {
	start := time.Now()
	// TODO: Executes the registered functions for given component; using maps
	//       If map for given key(aka component) is empty, nothing will be done
	//       Check installationType, to know which map should be used
	t := time.Now()
	duration += t.Sub(start)
}

// Returns duration of all jobs for benchmarking
func GetDuration() time.Duration {
	return duration
}

```

### component1.go - Exmaple component file draft

```go
package jobManager

// Register job using implemented Interface
register(job1)
type job1 struct{}

func (j job1) execute() {
	// Do something
  ...
  return nil
}

func (j job1) when() {
	return ("kiali", "1.22", Pre, Deploy)
}

```

### Hook for global jobs in `hydroform/parallel-install/deployment.go`; Pre/Post global Jobs - Deploy

Pre- and post-jobs will be executed before and after Kyma deploy.

```go
import "hydroform/parallel-install/jobs"
func (i *Deployment) deployComponents(ctx context.Context, cancelFunc context.CancelFunc, phase InstallationPhase, eng *engine.Engine, cancelTimeout time.Duration, quitTimeout time.Duration) error {
  ...
  deploymentJobs.ExecutePre("global", jobManager.Deploy)
  statusChan, err := eng.Deploy(ctx)
  ...
  // for-Loop for component install
  ...
  deploymentJobs.ExecutePost("global", jobManager.Deploy) 
}
```

### Hook for component jobs in `hydroform/parallel-install/engine.go`; Pre/Post Component Jobs - Deploy and Deletion 

Pre- and post-jobs will be executed before and after each Kyma component. In this way only the componentns which will be installed/upgraded will be considered by the JobManager &#8594; Other will just not be executed. Furthermore, since the pre-requisites of components are deployed using the `worker` function of an `engine` as well, the JobManager handles them automatically.

```go
import "hydroform/parallel-install/jobs"
...
  ... // async workers
  func (e *Engine) worker(ctx context.Context, wg *sync.WaitGroup, jobChan <-chan components.KymaComponent, statusChan chan<- components.KymaComponent, installType installationType) {
    ...
    case component, ok := <-jobChan:
    ...
    jobManager.ExecutePre(component.Name, installationType)
    
    component.deploy(ctx)
    
    jobManager.ExecutePost(component.Name, installationType)
    ...
}
```

### Hook for global jobs in `hydroform/parallel-install-deletion.go`; Pre/Post global Jobs - Deletion

Pre- and post-jobs will be executed before and after Kyma deletion.
```go
import "hydroform/parallel-install/jobs"
...
func (i *Deletion) uninstallComponents(ctx context.Context, cancelFunc context.CancelFunc, phase InstallationPhase, eng *engine.Engine, cancelTimeout time.Duration, quitTimeout time.Duration) error {
	...
  jobManager.ExecutePre("global", jobManager.Uninstall)
	statusChan, err := eng.Uninstall(ctx)
  ...
  // for-Loop for component deletion
  ...
  jobManager.ExecutePost("global", jobManager.Uninstall)
```

## Placement of logic and actual jobs

After a short discussion with the included Teams (Goats, Huskies) we decided to implement the logic and jobs for the jobManager in the installer library (hydroform repository) as a package, to keep it simple, clean, and easy to access.

The jobManager will be placed as a package inside of the `parallel-install` module.
```
hydroform
│   ...
└───parallel-install 
│     │   ...
│     └───jobManager
│     │     │   core.go // Register functions
│     │     │   logic.go // Outsource main logic, if needed
│     │     └───jobs
│     │           │   component1.go
│     │           │   component2.go
│     │           │   ...
│    ...         ...
...
```

## Additions

- To have a consistent output, we will use the Unified Logging library