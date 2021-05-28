# Dynamic Kyma-Install-Configuration (Migration Logic)

This PoC is used to come up with a valid design for the new __Migration Logic__, needed to have a fully automated Kyma deploy possible. It will be used to configure the cluster and the components during the deployment of Kyma. The term `Deployment`/`Deploy` is being used in the context of installing Kyma on an empty cluster or to upgrade Kyma from an older to a newer version.
To achieve a valid solution for the PoC we need to come up with a design for the following:
- Go-Implementation-Draft for the migration hook
- Placement of logic code-snippets
- Work-Flow how and when logic should be triggered
## Golang Implementation
```go
var Snippets *SnippetManager;
// Define type for go-code-snippets
type snippet func() err 
// Implement code-snippets for configuration/migration/etc.
func myfn1() err {
    return errors.New("Sample Error")
}
func myfn2() err {
    return errors.New("Sample Error")
}
type SnippetManager struct {
    preSnippetMap map[string][]snippet
    postSnippetMap map[string][]snippet
}
// Register function for code-snippets
func (sm *SnippetManager) Register(f snippet, component string, exec string) {
    // TODO: Build up Maps with the given code-snippets
}
// Function should be called before compoent is being deployed/upgraded
func (sm * SnippetManager) ExecutePre(component string) {
  // TODO: Executes the registered functions for given component
}
// Function should be called after compoent is being deployed/upgraded
func (sm * SnippetManager) ExecutePost(component string) {
  // TODO: Executes the registered functions for given component
}
// Initializes SnippetManager and return pointer
func NewSnippetManager() *SnippetManager {
    sm := &SnippetManager{}
    
    sm.Register(myfn1, "kiali", "pre")
    sm.Register(myfn2, "", "post")
    return sm
}

func init() {
  Snippets = NewSnippetManager()
}

```
Hook for generic snippets in `hydroform/parallel-install/deployment.go`
```go
import "hydroform/parallel-install/snippets"
func (i *Deployment) deployComponents(ctx context.Context, cancelFunc context.CancelFunc, phase InstallationPhase, eng *engine.Engine, cancelTimeout time.Duration, quitTimeout time.Duration) error {
  ...
  Snippets.ExecutePre("")
  statusChan, err := eng.Deploy(ctx)
  Snippets.ExecutePost("") 
  ...
}
```

Hook for component snippets in `hydroform/parallel-install/engine.go`
```go
import "hydroform/parallel-install/snippets"
...
  ... // async workers
  func (e *Engine) worker(ctx context.Context, wg *sync.WaitGroup, jobChan <-chan components.KymaComponent, statusChan chan<- components.KymaComponent, installType installationType) {
    ...
    case component, ok := <-jobChan:
    ...
    Snippets.ExecutePre(component.Name)
    component.deploy(ctx)
    Snippets.ExecutePost(component.Name)
    ...
}
```
| Paramter | Possible Values |Description |
|-----|-----|-----|
| fn | func |Function with migration code-snippet|
| component | string |Componenet the code-snippet belongs to, if empty then its generic|
| exec | "pre", "post"| Decides when the code-snippet should run, before or after the kyma/component deployment |
|||
|||

## Placement of logic code-snippets
After a short discussion with the inlcuded Teams (Goats, Huskies) we decided to implement the code-snippets for the configuration/migration in the installer library (Hydroform repository), to keep it simple and clean at first. When evolving further with the implementation of this PoC, we are going to have a second look if we find a more applicable solution.
```
hydroform
│   ...
└───parallel-install 
│     │   ...
│     └───migration
│     │     │   core.go // Register functions
│     │     │   logic.go // Outsource main logic, if needed
│     │     └───snippets
│     │           │   component1.go
│     │           │   component2.go
│     │           │   ...
│    ...         ...
...
```

## Work-Flow
<img style="float: right;" src="./migration-logic-diagram.png?raw=true">

- Only support single linear upgrade: A &#8594; B && B &#8594; C; NOT A &#8594; C
- __Component based__:
  - Check it component is installed on cluster, or if it wants to be newly installed
  - Migration should be done __before__ installation of component
    - Trigger confiuration logic
    - Install component
  - Migration should be done __after__ installation of component
    - Trigger configuration logic
- __Generic / Component independent__:
  - Always trigger logic
  - Before or after kyma installation
  - 
## Additions
- Log migration time for benchmarking (use go-timer)
- Use unified logging as output
- Everything should be precalculated --> Means code should know in advances which migration steps need to be done
- Rename Migration; since this components is not only for Migration, but also for "dirty" shells scripts