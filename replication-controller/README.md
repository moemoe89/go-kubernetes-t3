# KUBERNETES-REPLICATION-CONTROLLERS #

This Replication Controller project for practice Kubernetes

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
Create pod :
```
$ kubectl create -f go-hostname.yaml
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
$ kubectl describe pod go-hostname-{UNIQUE-ID}
```

## Delete
Delete pod :
```
$ kubectl delete pod go-hostname-{UNIQUE-ID}
```
Delete replication :
```
$ kubectl delete rc go-hostname-rc 
```
Delete replication without pod :
```
$ kubectl delete rc go-hostname-rc --cascade=false
```

## License

MIT