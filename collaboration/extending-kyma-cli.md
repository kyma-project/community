# Managing Kyma components using Kyma CLI

Kyma cli has been used for cluster management. We should look into how we can introduce support for management of components on the kyma cluster. This would bring in following benefits:

* Enable users who have limited or no kuberenetes knowledge to use kyma.
* Provide an abstraction on complex kubectl command chains.
* Could be easily used into CI/CD pipelines where they can have automated way of configuring kyma cluster to have functions deployed via CI/CD pipeline

## Architecture

We should follow the modular approach with each component having its own library. This library can be used inside the kyma cli for acessing various components. The library contains all the implementation. The kyma cli should invoke this library. We can take inspiration from hydroform and implement it in similar way. Currently we have following modules in kyma:

  1. Functions
  2. Service Catalog
  3. Events
  4. Application Connector
  5. APIs

## Command structure

The command structure used by cloud providers is following:

```bash
<cliName> <functionality> <action>
#example aws s3 ls

# Functions
kyma function <action>

# Events
kyma events <action>

# Application Connector
kyma applications <events>

# Apis
kyma apis <events>
```

### Functions

For the serverless component we should start with supporting following commands

#### Local workspace setup

For the local workspace setup, a developer should be provided with a project containing basic working function example. This should be deployable on to an existing kyma cluster using the kyma cli. It should also be possible to test this code locally by means of make files. User should be provded with some readme where he can understand the project structure and also the configuration examples. It should also contain commands to deploy and debug the function on to kyma cluster and also how to run locally.

##### Creation of project locally (approach 1)

This could be used for the user who is new to kyma and kubernetes in general. He should be able to have an example function created easily. We can have following commands.

```bash
  kyma function init <function-name> --runtime <runtime-name>

  kyma function init <function-name> --runtime <runtime-name> -n <namespace> -p <path>
```

> The namespace should be optional. When not passed it can be set to `default`. When the user has path to the function code (residing in git) we should generate the scaffolding around it (if its not existing already).
When path `-p` is passed then it should create the project in the designated path, otherwise, it uses the current directory.
The runtime corresponds to the language in which the function code would be written.

This should create directoy structure with following content based on the runtime passed (assuming vscode as the IDE).:

```bash
  <function-name>
  ├── .vscode
  │   └── launch.json
  ├── resources
  │   └── function.yaml
  |   └── api.yaml
  |   └── subscription.yaml
  ├── local
  │   └── <server-code.extension>
  └── Readme.md
  └── src
    ├── config.yaml
    ├── <function-code.extension>
    └── <dependencies.extension>
```

* launch.json consists of settings for the local debugger on vscode
* resource folder consists of the yaml of the function cr that has been deployed on the kyma cluster. This would be generated before function is being deployed. It would be generated on every deployment.
* `<server-code.extension>` consists of the server code with handler pointing to <function-code.extension>. The `extension` of file and `server-code` is with repsect to the runtime passed in init command
* config.yaml can be used as info file with current configurations like:
  * name
  * namespace
  * events subscribed
  * environment variables
  * Api name to expose with
  * replicas
  * min/max memory
  * min/max cpu
  * runtime
  * debug command (to be used with telepresence. This can help if we want to have different runtimes)
* <function-code.extension> where the developer can write his logic to test. The `extension` is runtime specific.
* `dependecies.extension` file with list of dependencies. eg. `package.json` for node and `go.mod and go.sum` for golang dependencies

>The generated templates should be placed in a directory in the library for the functions. We should have templates for each runtime that can be supported. In the future if we are supporting more runtime we can move it to separate repo.

##### config.yaml

This file basically enables developers to configure their functions. With init command this file is created and it basically has some default values. The readme files provide information about various options present in this file

Developer can modify this file to deploy or update a function on to the kyma cluster. This would enable the developer not only tune the function (by modifying replica, memory cpu), but also provide an overview of the current triggers for the function.

Usually one does not change the config so often during development compared to code. So our aim should be to make deployment and testing the function simpler. So having a config file taking care of configuration would make deployment much simpler.

Having such a config file would help developers with limited or no kubernetes knowledge to get themselves aquainted with various configuration options for function and concentrate on the business logic for writing the functions.

Additionally it can be used for automated deployments of functions. As this file can be read by kyma cli to deloy function with required configurations.

Example structure of config.yaml

```yaml
name: <function_name>
namespace: <namespace> #default by default
events:
  - event1
  - event2
environment-vars:
  - foo:bar
api:
  url: <url>
  actions:
    - GET
    - POST
  security:
    - JWT
replicas: 1
memory:
  min:
  max:
cpu:
  min:
  max:
runtime:
debug-string:
```

##### Template files

###### <function-code.extension>

  Out of the box from the init command we should give an easy example which can be extended by the developer later. The example should be possible to be deployed directly on the kyma cluster. Here below there is a node example

  ```js
    module.exports = {
      main: function (event, context) {
        return 'Hello World';
      }
    };
  ```

###### <dependecies.extension>

We should generate an empty package.json file which can be edited by developer as and when needed.

###### index.js

We should provide `index.js`(which would start the server and serve the handler function) out of the box which can be used by the developer to test his code locally.

###### readme

We should also provide a readme.md which would consist of following instructions:

* How he can deploy the function into the cluster.
* How he can run the function locally.
* How to use the `config.yaml` file
* Folder structure.
* How to test debug it locally like using telepresence.

##### Testing/Debugging

*Local*

We can introduce commands to test the function locally and also if the developer wants to debug the code. These can be achieved through make file.

*Using Events from remote cluster*

Sometimes the event payloads can be complex to be created locally. For such cases its nice to have the events on remote cluster triggering the code on the workstation. This would enable the developer to debug his code on the workstation using events being sent to the remote cluster. We can introduce command for example below:

`kyma function debug <function-name> -n <namespace>`

> We can use telepresence for local debugging. We can either use --swap-deployment or --new-deployment option.

Telepresence would create a proxy in the kyma cluster and redirect the calls to local codebase. One can attach the debugger to his code base along with breakpoints to debug the code.

##### Deploying and running the project

The function generated by `init` command should be deployable onto kyma cluster. For deploying the file onto kyma cluster it would read the values from `config.yaml`. If the api name is not set then we should show some warning message and deploy.

For running/debugging it locally we should provide makefile. The running/debugging commands should be part of readme.

##### Creation of project locally (approach 2)

The other approach would include creation of the local project structure as shown before. But it would not contain any `config.yaml`. Rather after creation of the project structure the function could be deployed first time via the command `kyma function apply -p <path_to_function_dir>`. But for rest of the changes like modifying the mem/cpu or api rules (like verbs or security) or like events and service binding should be done via commands for each module as listerd [expanded commands with switches](#expanded_command_with_switches)

##### Pros and Cons approach 1 vs approach 2

*Pros Apprach 1*

* `config.yaml` helps the user with no kuberenertes knowledge to easily on board kyma. One can write configuration and can deploy it. It provides a summary for the current configuration of the application

* Since during the developement of function we dont change the config so many times `kyma function apply` is an easy command to use. Also if the config needs to be changed and re-applied then `kyma function apply` is stil small command to use.

* It could be eventually used as a recipie for deployment of functions on to the production cluster using CI/CD pipeline

* Special runtime specific info like `debug` flags when using with `telepresence` could be passed as field in the `config.yaml`

*Cons Approach 1*

* Once the function has been deployed and if the function has been modified on the cluster using commands listed at [expanded commands with switches](#expanded_command_with_switches), then the local `config.yaml` is out of sync. Re applying via `kyma function apply` can result in confusion.

* The `config.yaml` might at the end look like a CRD.

*Pros Approach 2*

* We dont have a `config.yaml` so we dont need to sync states between local and remote on deployed functions.

*Cons Approach 2*

* One has to use lengthy commands to configure functions.
* Difficult to have an overview of current state.
* Need to figure out how to use debug commands (when using telepresence).

#### Working with deployed kyma cluster

We should allow deployment of the functions on a k8s cluster. Following commands should be supported

##### CRUD

##### Commands that use config

`kyma function apply -p <path-to-function-dir>`
> This would create/update a function after reading the `config.yaml` from the path provided and to read the config like events, api etc. For deletion we should use the delete method.

##### Expanded command with switches

```bash
  kyma function create <function-name>
     -n <namespace>
     --runtime <runtime>  
     --src-code <path-to-src-code>
     --dependecies <path-to-dep-file>
     [--min-mem <mb>] [--max-mem <mb>]
     [--min-cpu <cpu>] [--max-cpu <cpu>]
     [--min-replicas 1] [--max-replicas 1]
```

```bash
kyma function update <function-name>
     -n <namespace>
     --runtime <runtime>  
     --src-code <path-to-src-code>
     --dependecies <path-to-dep-file>
     [--min-mem <mb>] [--max-mem <mb>]
     [--min-cpu <cpu>] [--max-cpu <cpu>]
     [--min-replicas 1] [--max-replicas 1]
```

`kyma function delete <function-name> -n <namespace>`

`kyma function get <function-name> -n <namespace>`

*Labels*

`kyma function label <function-name> ["foo=bar",..] -n <namespace>`

*Expose*

`kyma function expose <function-name> [--secure] [--actions <GET/POST>] -n <namespace>`

*Service Catalog*

`kyma function bind <function-name> --binding-instaces [<binding-instance-name>,...] -n <namespace>`
`kyma function unbind <function-name> --binding-instaces [<binding-instance-name>,...] -n <namespace>`

`kyma function bind <function-name> --binding-usage [<binding-usage-name>,...] -n <namespace>`
`kyma function unbind <function-name> --binding-usage [<binding-usage-name>,...] -n <namespace>`

*Bind to events*

`kyma function bind <function-name> --events [<event-name>,..] -n <namespace>`
`kyma function bind <function-name> --events [<event-name>,..] -n <namespace>`

*Bind to environment variables*

`kyma function bind <function-name> --env [<foo=bar>,..] -n <namespace>`
`kyma function unbind <function-name> --env [<foo=bar>,..] -n <namespace>`

*List all functions*

`kyma function list -n <namespace> [--labels foo=bar]`

> A detailed table view with function name, runtime, age, replicas and state

*Runtimes*

`kyma function get runtimes`

*Logs*

`kyma function logs <function-name> -n <namespace>`

*function Status*

`kyma function status <function-name> -n <namespace>`

Should display status

```bash
  Name: <function-name>
  Runtime: <my-runtime>
  State: Running/Error/Deploying/...
  Replicas: 1
  Min-replicas: 1
  Max-replicas: 2
  API: <url>
    Secure: yes/no
    Actions: GET, POST
  Bindings:
    Events: <list of events>
    Servicebindings: <list of service bindings>
  labels: foo:bar
```

`kyma function show <function-name> events`
`kyma function show <function-name> apis`
`kyma function show <function-name> service-binding-usages`
`kyma function show <function-name> labels`
`kyma function show <function-name> replicas`

#### Typical User flow (approach 1)

This section details a typical development flow for the functions

1. New User creating function from scratch or from existing github project as shown [here](#creation-of-project-locally). This would generate a function project for vscode. User is also provided with readme file to explain the structure of project and also the config.yaml. It should also explain how user can test it locally and also debugging.

2. Once the project is initialized user can deploy the function onto an exisiting kyma cluster using command [here](#commands-that-use-config). It would use the `config.yaml` to deploy all resources.

3. User can change the code and test it locally by running a server locally on his workstation. User can make use of make file provided.

4. For debugging purpose user can either run it locally and attach the debuger or for in-cluster debugging he can use the `kyma function debug` command which would use telepresence to redirect the event triggers for example to the local code and he can attach the debugger to debug.

#### Typical User flow (approach 2)

1. New User creating function from scratch or from existing github project as shown [here](#creation-of-project-locally). This would generate a function project for vscode. User is also provided with readme file to explain the structure of project and also the config.yaml. It should also explain how user can test it locally and also debugging.

2. Once the project is initialized user can deploy the function onto an exisiting kyma cluster using command [here](#commands-that-use-config)

3. Use the commands listed at [expanded commands with switches](#expanded_command_with_switches) to configure the functions

4. Debugging ?

### Service catalog

We should enable kyma cli to be able to configure service catalog too. We should support following commands. The implementation of the commands should be abstracted in a separate library:

`kyma service-instance list`

#### Service Bindings

##### CRUD

`kyma service-binding create <service-binding-name> -n <namespace> --service-instance <service-instance-name>`

`kyma service-binding delete <service-binding-name> -n <namespace>`

`kyma service-binding get <service-binding-name> -n <namespace>`

#### Service Binding Usage

##### CRUD

`kyma service-binding-usage create <sbu-name> -n namespace --service-binding <service-binding-name> --function <function-name>`

`kyma service-binding-usage create <sbu-name> -n namespace --service-binding <service-binding-name> --microservice <microservice-name>`

`kyma service-binding-usage update <sbu-name> -n namespace --service-binding <service-binding-name> --function <function-name>`

`kyma service-binding-usage update <sbu-name> -n namespace --service-binding <service-binding-name> --microservice <microservice-name>`

`kyma service-binding-usage delete <sbu-name> -n namespace`

`kyma service-binding-usage delete <sbu-name> -n namespace`

`kyma service-binding-usage get <sbu-name> -n namespace`

`kyma service-binding-usage get <sbu-name> -n namespace`

### Events

We should enable support for events in kyma cli too. We should have support for following commands and implementation abstracted as a separate library:

`kyma event list -n <namespace>`

`kyma event create -n <namespace> subscriptions --topic {"foo":"bar"}`

`kyma event update -n <namespace> subscriptions --topic {"foo":"bar"}`

`kyma event list -n <namespace> subscriptions`

`kyma event trigger --topic {"foo":"bar"} --data {"foo": "bar"}`

### Aplication Connector

Application connector should also have a separate library invoked inside the kyma cli. It should also support following commands: