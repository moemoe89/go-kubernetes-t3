# KUBERNETES-PODS #

This Pod project for practice Kubernetes

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
$ kubectl create -f go-hello-world.yaml
```

## Check
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
$ kubectl describe pod go-hello-world
```

## Access
Access pod :
```
$ kubectl port-forward go-hello-world 8782:8782
```
Navigate browser to this address :
```
localhost:8782
```

## Delete
Delete pod :
```
$ kubectl delete pod go-hello-world --now 
```

## License

MIT