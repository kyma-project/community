# Extend Kyma CLI

Kyma cli has been used to provision a cluster and install kyma into. We should look into extend the current functionality set such that we can provide functionality to access various features of kyma.

The extension of functionality would enable people with no or less kubernetes knowledge can use kyma without any issue.

## Security

Kyma cli should support RBAC. This would be realised by downloading the requisite kubeconfig file. The admins should be able to use more infrastructure commands and developers should be able to use the kyma cluster with restricted rights.

## Architecture

  We should follow the modular approach with each component having its own library. This library can be used inside the kyma cli for acessing various components. The library contains all the implementation. The kyma cli should invoke this library. We can take inspiration from hydroform and implement it in similar way. Currently we have following modules in kyma:

  1. Lambdas
  2. Service Catalog
  3. Events
  4. Application Connector
  5. APIs

## Command structure

  The command structure used by cloud providers is following:

  ```bash
  <cliName> <functionality> <action>
  #example aws s3 ls
  ```

  The current kyma-cli also uses the same pattern eg.

  ```bash
  kyma provision
  or
  kyma install
  ```

### Lambdas

For the serverless component we should start with supporting following commands

#### Local workspace setup

1. Creation of project locally
  `kyma lambda init <lambda-name> --runtime <runtime-name> -n <namespace>` or `kyma lambda init <lambda-name> --runtime <runtime-name> -n <namespace> -p <path>`
  This should create directoy with following content based on the runtime:

      ```bash
        <lambda-name>
        ├── .vscode
        │   └── launch.json
        ├── deployment
        │   └── deployment.yaml
        ├── local
        │   └── index.js
        └── code
          ├── config.yaml
          ├── handler.js
          └── package.json
      ```

    > When path `-p` is passed then it should create in desitnation path else it use the workspace set in the yaml file.

    * launch.json consists of settings for the debugging
    * deployment.yaml consists of the yaml that has been deployed on the kyma cluster.
    * `index.js` consists of the server code with handler pointing to `handler.js`
    * config.yaml can be used as info file with current configurations like:
      - name
      - namespace
      - events subscribed
      - environment variables
      - Api name to expose with
      - replicas
      - min/max memory
      - min/max cpu
      - runtime
      - debug command (to be used with telepresence. This can help if we want to have different runtimes)
    * handler.js where the developer can write his logic to test
    * package.json file with dependencies.

    >The templates generated should be placed a directory in the kyma cli. In the future if we are supporting more runtimes we can move it to separate repo.

2. Testing/debugging it locally

We can introduce commands to test the lambda locally and also if the developer wants to debug the code. This is useful in case when we have some event payloads that are not easy to recreate.

`kyma lambda debug <lambda-name>`

> We can use telepresence for local debugging. We can either use --swap-deployment or --new-deployment option

#### Working with deployed kyma cluster

We should allow deployment of the lambdas on a k8s cluster. For this we should use the service account role. Following commands should be supported

##### CRUD

`kyma lambda create <lambda-name> -n <namespace> --runtime <runtime> --min-mem <mb> --max-mem <mb> --min-cpu <cpu> --max-cpu <cpu> --min-replicas 1 --max-replicas 1 --src-code <path-to-src-code> --dependecies <path-to-dep-file>`

`kyma lambda create -p <path-to-lambda-dir>`
> This would create a lambda after reading the deployment.yaml from the path provided and also would read `config.yaml` to read the config like events, api etc.

`kyma lambda update <lambda-name> -n <namespace> --runtime <runtime> --min-mem <mb> --max-mem <mb> --min-cpu <cpu> --max-cpu <cpu> --min-replicas 1 --max-replicas 1 --src-code <path-to-src-code> --dependecies <path-to-dep-file>`

`kyma lambda update -p <path-to-lambda-dir>`
 > Update an existing lambda with the deployment.yaml passed.

`kyma lambda delete <lambda-name> -n <namespace>`

`kyma lambda get <lambda-name> -n <namespace>`

##### Labels

`kyma lambda label <lambda-name> ["foo=bar",..] -n <namespace>`

##### Expose

`kyma lambda trigger <lambda-name> --http --secure(optional) -n <namespace>`

##### Service Catalog

`kyma lambda bind <lambda-name> --binding-instaces [<binding-instance-name>,...] -n <namespace>`
`kyma lambda unbind <lambda-name> --binding-instaces [<binding-instance-name>,...] -n <namespace>`

`kyma lambda bind <lambda-name> --binding-usage [<binding-usage-name>,...] -n <namespace>`
`kyma lambda unbind <lambda-name> --binding-usage [<binding-usage-name>,...] -n <namespace>`


##### Bind to events

`kyma lambda trigger <lambda-name> --events [<event-name>,..] -n <namespace>`
`kyma lambda trigger <lambda-name> --events [<event-name>,..] -n <namespace>`

##### Bind to environment variables

`kyma lambda bind <lambda-name> --env [<foo=bar>,..] -n <namespace>`
`kyma lambda unbind <lambda-name> --env [<foo=bar>,..] -n <namespace>`

##### List all functions

`kyma lambda list -n <namespace>`

> A detailed table view with lambda name, runtime, age, replicas and state

##### Runtimes

`kyma lambda get runtimes`

##### Logs

`kyma lambda <lambda-name> logs -n <namespace>`

##### Lambda Status

`kyma lambda status <lambda-name> -n <namespace>`

Should display status

```bash
  Name: <lambda-name>
  Runtime: <my-runtime>
  State: Running/Error/Deploying/...
  Replicas: 1
  Min-replicas: 1
  Max-replicas: 2
  API: <url>
  Secure: yes/no
  Bindings:
    Events: <list of events>
    Servicebindings: <list of service bindings>
  labels: foo:bar
```

`kyma lambda show <lambda-name> trigger`
`kyma lambda show <lambda-name> trigger --events`
`kyma lambda show <lambda-name> trigger --http`
`kyma lambda show <lambda-name> service-instances`
`kyma lambda show <lambda-name> labels`
`kyma lambda show <lambda-name> replicas`

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