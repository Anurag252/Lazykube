# `kubectl-command-cli`

A CLI tool that allows users to quickly find `kubectl` commands and their documentation.

## How to use

After installing the tool, use the following command to search for a `kubectl` command:

`kubectl-command-cli kubectl-command`


This command will search for all the `kubectl` commands that match the given input and display their descriptions.

## Code structure

The codebase is divided into the following files:

- `main.go`: The entry point of the application.
- `kubesearch.go`: Contains the search command and related functions.
- `readcheatsheet.go`: Contains the function to read the YAML file that contains the `kubectl` commands and their descriptions.
- `root.go`: The root command of the application.

## Dependencies

- `cobra`: A library for creating powerful modern CLI applications in Go.
- `yaml.v2`: A library for parsing and encoding YAML files in Go.

## Installation

To install the `kubectl-command-cli` tool, run the following command:

`go install anurag252/kubectl-cli/cmd/kubectl-command-cli`


This will install the tool in your `$GOBIN` directory.

## Usage

To use the tool, run the following command:

`kubectl-command-cli <command>`

Replace `<command>` with the search term you want to use to find `kubectl` commands.

output

```
kubectl-cli kubectl-search get
kubectl config view --> Get the configuration of the cluster
kubectl get all --all-namespaces --> List everything
kubectl get daemonset --> List one or more daemonsets
kubectl get deployment --> List one or more deployments
kubectl get events --> List recent events for all resources in the system
kubectl get events --field-selector type=Warning --> List Warnings only
kubectl get events --field-selector involvedObject.kind!=Pod --> List events but exclude Pod events
kubectl get events --field-selector involvedObject.kind=Node, involvedObject.name=<node_name> --> Pull events for a single node with a specific name
kubectl get events --field-selector type!=Normal --> Filter out normal events from a list of events
kubectl logs --tail=20 <pod_name> --> Get the most recent 20 lines of logs
kubectl logs -f <service_name> [-c <$container>] --> Get logs from a service and optionally select which container
kubectl get namespace <namespace_name> --> List one or more namespaces
kubectl get node --> List one or more nodes
kubectl get pods -o wide | grep <node_name> --> Pods running on a node
kubectl get pod --> List one or more pods
kubectl exec -it <pod_name> /bin/sh --> Get interactive shell on a a single-container pod
kubectl get rc --> List the replication controllers
kubectl get rc --namespace=”<namespace_name>” --> List the replication controllers by namespace
kubectl get replicasets --> List ReplicaSets
kubectl get secrets --> List secrets
kubectl get services --> List one or more services
kubectl get serviceaccounts --> List service accounts
kubectl get statefulset --> List StatefulSet
```
