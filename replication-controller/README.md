# KUBERNETES-REPLICATION-CONTROLLERS #

This Replication Controller project for practice Kubernetes. The go-hostname image container is a simple app that running on port 5000

All of configuration will be following port 5000 for the application

## Setup

* Setup Virtual Box <https://virtualbox.org/>
* Setup Docker <https://www.docker.io/>
* Setup Minikube <https://kubernetes.io/docs/tasks/tools/install-minikube/>

## Run Minikube
Running minikube with this command :
```
$ minikube start
```
If any issue with network, use virtualbox driver :
```
$ minikube start --vm-driver virtualbox
```

## Create
Create replication :
```
$ kubectl create -f [REPLICATION CONTROLLER YAML FILE]
$ kubectl create -f go-hostname-rc.yaml
```

## Check
Check replication :
```
$ kubectl get replicationcontrollers
or
$ kubectl get replicationcontroller
or
$ kubectl get rc
```
Describe replication set :
```
$ kubectl describe rc/[REPLICATION CONTROLLER NAME]
$ kubectl describe rc/go-hostname-rc
```
Check pod :
```
$ kubectl get pod
```
Check pod more detail :
```
$ kubectl get pod -o wide
```
Describe pod :
```
$ kubectl describe pod [POD NAME]
$ kubectl describe pod go-hostname-{UNIQUE-ID}
```

## Delete
Delete pod :
```
$ kubectl delete pod [POD NAME]
$ kubectl delete pod go-hostname-{UNIQUE-ID}
```
Delete replication :
```
$ kubectl delete rc [REPLICATION CONTROLLER NAME]
$ kubectl delete rc go-hostname-rc 
```
Delete replication without pod :
```
$ kubectl delete rc [REPLICATION CONTROLLER NAME] --cascade=false
$ kubectl delete rc go-hostname-rc --cascade=false
```

## License

MIT