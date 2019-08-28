---
title: Using Telepresence for local Kyma development
---

This document is a general guide to local development with [Telepresence](https://www.telepresence.io/) (tested on version `0.101`).

The Kyma component that you want to debug stores its state in the Kubernetes Custom Resource and, therefore, depends on Kubernetes.    
Mocking the dependency and developing locally are not possible, and manual deployment on every change is a mundane task.  

Telepresence is a tool that connects your local process to a remote Kubernetes cluster through proxy, which lets you easily debug locally.  
When you use Telepresence, it replaces a container in a specified Pod, opens up a new local shell or a pre-configured bash, and proxies the network traffic from the local shell through the Pod. 

Telepresence enables you to make calls such as `curl http://....svc.cluster.local:8081/v1/metadata/services` from your local machine.  
When you run a server in this shell, other Kubernetes services can access it. 

To start developing with Telepresence, follow these steps:

1. [Install Telepresence](https://www.telepresence.io/reference/install).

2. Run your local Kyma or use the cluster. Then, configure your local kubectl to use the desired Kyma cluster. 

3. Check the deployment name to swap, and run: 

	```
	telepresence --namespace {NAMESPACE} --swap-deployment {DEPLOYMENT_NAME}:{CONTAINER_NAME} --run-shell
	```

4. Every Kubernetes Pod has the directory `/var/run/secrets` mounted. The Kubernetes client uses it in the component services. By default, Telepresence copies this directory. It stores the directory path in `$TELEPRESENCE_ROOT`, under the Telepresence shell. The `$TELEPRESENCE_ROOT` variable unwinds to `/tmp/...`. You need to move it to `/var/run/secrets`, where the service expects it. To move it there, create a symlink:
	```
	sudo ln -s $TELEPRESENCE_ROOT/var/run/secrets /var/run/secrets
	```

5. Run `CGO_ENABLED=0 go build ./cmd/{COMPONENT-NAME}` to build the component and give all Kubernetes services that call the component access to this process. The process runs locally on your machine. Use the same command to run various Application Connector services like Application Registry, Proxy or Events.