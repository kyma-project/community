# JobManager (Migration Logic)

This PoC is used to come up with a valid design for the new __Migration Logic__, needed to have a fully automated Kyma deploy possible. It will be used to configure the cluster and the components during the deployment of Kyma. The term `Deployment`/`Deploy` is being used in the context of installing Kyma on an empty cluster or to upgrade Kyma from an older to a newer version.
In the following the term `code-snippets` will be used for the Go-Code snippets which are in general the jobs triggered by the new hooks.

To achieve a valid solution for the PoC we need to come up with a design for the following:

- [JobManager (Migration Logic)](#jobmanager-migration-logic)
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
- This mechanism supports code-snippets for two different main use-cases: The __component based__ snippets and the __generic/non-component based__ snippets
  - __Component based__:
    - Check it component is installed on cluster, or if it wants to be newly installed, and only trigger if yes
    - It should be possible to trigger code-snippets before and after an deployment of a component
  - __Generic / Component independent__:
    - Always trigger logic when installing, upgrading, or uninstalling Kyma
    - It should be possible to trigger code-snippets before and after the deployment, deletion of Kyma

### Possible Solution

To fulfill the requirements a new `JobManager` will be introduced, which is capable to register, manage, and trigger certain code-snippets to have a fully automated install/migration/deletion. This struct has two (hash)maps: One for `pre`-jobs and one for `post`-jobs. These maps are  being filled up using the `Register` function. As an input for this function the reference to the job/code-snippet, the refering component name, at which point the job should be triggered (pre/post), and the desired kyma version is needed. As a key for the maps the name of the desired components will be used, which are then pointing to the wanted job/code-snipped. Those two maps will only be filled with the jobs which are needed for the wanted installation type (deploy/uninstall) and for the desired Kyma version. Thus we do not have additional checks if a job is needed, at a later point in the code. Then at the hooks during the deployment/deletion-phase the hook just has to check if the key for the wanted component is present in the pre/post-map, if yes the jobs in the map will be trigged, if no nothing has to be done.

To handle the distinction between deploy and deletion, two JobManager will be used to reduce the amount of complex structures.

To benchmark the jobs, a timer will be used in the pre and post job triggers

<img src="./migration-logic-diagram.png?raw=true">

## Draft for Golang Implementation

```go
var deploymentJobs *JobManager;
var deletionJobs *JobManager;

// Define type for jobs
type job func() err 

// Implement code-snippets/jobs for configuration/migration/etc.
func myfn1() err {
    return errors.New("Sample Error")
}
func myfn2() err {
    return errors.New("Sample Error")
}
func myfn3() err {
    return errors.New("Sample Error")
}

type JobManager struct {
    duration float32
    kymaVersion string
    installationType string
    preJobMap map[string][]job
    postJobMap map[string][]job
}

// Register function for code-snippets/jobs
func (jm *JobManager) Register(f job, component string, exec string, version string) {
    // TODO: Build up Maps with the given code-snippets/jobs
    //       If version == kymaVersion insert pointer to job in corresponding map
}

// Function should be called before compoent is being deployed/upgraded
func (jm * JobManager) ExecutePre(component string) {
  start := time.Now()
  // TODO: Executes the registered functions for given component; using preJobMap
  //       If map for given key(component) is empty, nothing will be done
  t := time.Now()
  jm.Duration += t.Sub(start)
}

// Function should be called after compoent is being deployed/upgraded
func (jm * JobManager) ExecutePost(component string) {
  start := time.Now()
  // TODO: Executes the registered functions for given component; using postJobMap
  //       If map for given key(component) is empty, nothing will be done
  t := time.Now()
  jm.Duration += t.Sub(start)
}

// Initializes JobManager and return pointer
func NewJobManager(installationType string, kymaVersion string) *JobManager {
    jm := &JobManager{}
    jm.kymaVersion = kymaVersion

    //Exmaples for registration
    if ( installationType == "deploy") {
      jm.Register(myfn1, "kiali", "pre", "1.22")
      jm.Register(myfn2, "generic", "post", "1.21")
      // ..
    } else if (installationType == "uninstall") {
      jm.Register(myfn3, "kiali", "post", "1.22")
      // ...
    } else {
      // Error
    }
    return jm
}

// Will be executed once when package is imported
func init() {
  deploymentJobs = NewJobManager("deploy", kymaVersion)
  deletionJobs = NewJobManager("uninstall", kymaVersion)
}

```
Hook for generic jobs in `hydroform/parallel-install/deployment.go`; Pre/Post Generic Jobs - Deploy
```go
import "hydroform/parallel-install/jobs"
func (i *Deployment) deployComponents(ctx context.Context, cancelFunc context.CancelFunc, phase InstallationPhase, eng *engine.Engine, cancelTimeout time.Duration, quitTimeout time.Duration) error {
  ...
  deploymentJobs.ExecutePre("generic")
  statusChan, err := eng.Deploy(ctx)
  ...
  // for-Loop for component install
  ...
  deploymentJobs.ExecutePost("generic") 
}
```

Hook for component jobs in `hydroform/parallel-install/engine.go`; Pre/Post Component Jobs - Deploy and Deletion 
```go
import "hydroform/parallel-install/jobs"
...
  ... // async workers
  func (e *Engine) worker(ctx context.Context, wg *sync.WaitGroup, jobChan <-chan components.KymaComponent, statusChan chan<- components.KymaComponent, installType installationType) {
    ...
    case component, ok := <-jobChan:
    ...
    if ( installationType == "deploy") {
      deploymentJobs.ExecutePre(component.Name)
    }
    else {
      deletionJobs.ExecutePre(component.Name)
    }
    
    component.deploy(ctx)
    if ( installationType == "deploy") {
      deploymentJobs.ExecutePost(component.Name)
    }
    else {
      deletionJobs.ExecutePost(component.Name)
    }
    ...
}
```

Hook for generic jobs in `hydroform/parallel-install-deletion.go`; Pre/Post Generic Jobs - Deletion
```go
import "hydroform/parallel-install/jobs"
...
func (i *Deletion) uninstallComponents(ctx context.Context, cancelFunc context.CancelFunc, phase InstallationPhase, eng *engine.Engine, cancelTimeout time.Duration, quitTimeout time.Duration) error {
	...
  deletionJobs.ExecutePre("generic", )
	statusChan, err := eng.Uninstall(ctx)
  ...
  // for-Loop for component deletion
  ...
  deletionJobs.ExecutePost("generic")
```
Paramter for `Register` function:
| Paramter | Possible Values |Description |
|-----|-----|-----|
| fn | func |Function with job code|
| component | string |Component the job belongs to, or "generic" if it should be run before/after whole Kyma deploy/deletion|
| exec | "pre-deploy", "post-deploy", "pre-deletion", "post-deletion"| Decides when the job should run, before or after the kyma/component deployment |
|version| string, i.e. "1.22"| Kyma version to install/ugrade to, or uninstall |

## Placement of logic and actual jobs

After a short discussion with the included Teams (Goats, Huskies) we decided to implement the logic and jobs for the configuration/migration in the installer library (hydroform repository), to keep it simple and clean at first. When evolving further with the implementation of this PoC, we are going to have a second look if we find a more applicable solution.

The migration logic will be placed as a package inside of the `parallel-install` module.
```
hydroform
│   ...
└───parallel-install 
│     │   ...
│     └───migration
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