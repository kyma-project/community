# Dynamic Kyma-Install-Configuration (Migration Logic)


This PoC is used to come up with a valid design for the new __Migration Logic__, needed to have a fully automated Kyma deploy possible. It will be used to configure the cluster and the components during the deployment of Kyma. The term `Deployment`/`Deploy` is being used in the context of installing Kyma on an empty cluster or to upgrade Kyma from an older to a newer version.

To achieve a valid solution for the PoC we need to come up with a design for the following:

- Go-Implementation-Draft for the migration hook
- Placement of logic code-snippets
- Work-Flow how and when logic should be triggered

## Golang Implementation

```
// Define type for go-code-snippets
type fn func() err 

// Implement code-snippets for configuration/migration/etc.
func myfn1() err {
    return errors.New("Sample Error")
}
func myfn2() err {
    return errors.New("Sample Error")
}


type funcExecPair struct {
    f fn,
    exec string,
}

var componentsList map[string][]funcExecPair

// Register function for code-snippets
func registerConfiguration(f fn, component string, exec string) {
    // TODO
}

//Example usage
func main() {
    registerConfiguration(myfn1, "kiali", before)
    registerConfiguration(myfn2, "", after)

}
```

| Paramter | Possible Values |Description |
|-----|-----|-----|
| fn | func |Function with migration code-snippet|
| component | string |Componenet the code-snippet belongs to, if empty then its generic|
| exec | "before", "after"| Decides when the code-snippet should run, before or after the kyma/component deployment |
|||
|||



## Placement of logic code-snippets

After a short discussion with the inlcuded Teams (Goats, Huskies) we decided to implement the code-snippets for the configuration/migration in the installer library (Hydroform repository), to keep it simple and clean at first. When evolving further with the implementation of this PoC, we are going to have a second look if we find a more applicable solution.

```
hydroform
│   ...
└───migration
│   │   main.go // Register functions
│   │   logic.go // Main logic
│   └───config
│       │   component1.go
│       │   component2.go
│       │   ...
│      ...
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
