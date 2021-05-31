# JobManager (Migration Logic)

This PoC investigates a valid design for the new __JobManager__, which is needed to enable a fully-automated Kyma deploy. It will be used to configure the cluster and the components during the deployment of Kyma. The terms "Deployment" and "Deploy" are used in the context of installing Kyma on an empty cluster, or to upgrade Kyma from an older to a newer version.

To achieve a valid solution for the PoC we need to come up with a design for the following:

- [WorkFlow / General mechanism](#workflow--general-mechanism)
  - [Requirements](#requirements)
  - [Possible Solution](#possible-solution)
- [Draft for Golang Implementation](#draft-for-golang-implementation)
- [Placement of logic and actual jobs](#placement-of-logic-and-actual-jobs)
- [Additions](#additions)

## Workflow / General mechanism

### Requirements

- Only support single linear upgrade: A &#8594; B && B &#8594; C; NOT A &#8594; C. This is due to the fact that Kyma only supports single linear upgrades.
- The mechanism should only trigger for one specific Kyma version, this could be: The Kyma version that you want to install on an empty cluster, the Kyma version you want to upgrade to, or the Kyma version you want to uninstall from the cluster.
- This mechanism supports jobs for two different use cases: The __component-based__ jobs and the __global/component-independent__ jobs
  - __Component-based__:
    - Check whether the component is installed on the cluster or must be newly installed; and only trigger if it must be installed.
    - It should be possible to trigger jobs before and after a deployment of a component.
  - __Global / Component-independent__:
    - Always trigger the logic when installing, upgrading, or uninstalling Kyma.
    - It should be possible to trigger jobs before and after the deployment or deletion of Kyma.
    - Call component-independent jobs `global` jobs to stick to the naming convention of our helm charts.

### Possible Solution

To fulfill the requirements, a new package, called `JobManager`, is introduced, which registers, manages, and triggers certain jobs to have a fully-automated installation, migration, or deletion. This package has four (hash)maps to manage the workload: Two for `pre`-jobs (deploy and deletion) and two for `post`-jobs (deploy and deletion). In the (hash)maps, the key is the name of the component the jobs belong to, and the value is a slice of the jobs.
Furthermore, the `JobManager`package has a `duration` variable for benchmarking, and a `targetVersion` variable to know which jobs should be triggered at a certain deploy.

Jobs are implemented within the `JobManager` package in `go`-files, one for each component, using the specific `job`-interface. Then, the implemented interface is registered using `register(job)` in the same file. This function queues the jobs into a slice, because until then the value of the targetVersion is unknown. 

The JobManager is used by the `deployment` package in the `deployment.go` and the `deletion.go` file. Within the `NewDeployment` or `NewDeletion` functions, the `SetKymaVersion` function of the JobManager is called to set the targetVersion, and thus to build the needed maps. Doing it in this way, we only have the required jobs registered in the maps and save some checks later on (&#8594; everything pre-calculated). Then, at the hooks, during the deployment or deletion phase, each hook only has to check if the key for the wanted component is present in the pre/post-map. If it's present, the jobs in the map are trigged, if not, nothing must be done.

To benchmark the jobs, a timer is used in the pre- and post-job triggers.

Retries for the jobs are not handled by the JobManager. Retries should be implemented by the jobs themselves, because it's more flexible and the interface is easy to manage.

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

// Register job using implemented interface
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

Pre- and post-jobs will be executed before and after each Kyma component. In this way, only the components that will be installed or upgraded are considered by the JobManager &#8594; others are not executed. Furthermore, since the prerequisites of components are deployed using the `worker` function of an `engine` as well, the JobManager handles them automatically.

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

After a short discussion with the included Teams (Goats, Huskies), we decided to implement the logic and jobs for the jobManager in the installer library (hydroform repository) as a package, to keep it simple, clean, and easy to access.

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

- To have a consistent output, we will use the Unified Logging library.
