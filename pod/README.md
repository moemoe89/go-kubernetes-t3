# KUBERNETES-PODS #

This Pod project for practice Kubernetes. The go-hostname image container is a simple app that running on port 5000

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
Create pod :
```
$ kubectl apply -f [POD YAML FILE]
$ kubectl apply -f go-hostname-pod.yaml
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
$ kubectl describe pod [POD NAME]
$ kubectl describe pod go-hostname-pod
```

## Access
Access pod:
```
$ kubectl port-forward [POD NAME] [ACCESS PORT]:[TARGET PORT]
$ kubectl port-forward go-hostname-pod 5000:5000
```
Navigate browser to this address :
```
localhost:5000
```

## Delete
Delete pod :
```
$ kubectl delete pod [POD NAME] --now
$ kubectl delete pod go-hostname-pod --now
```

## License

MIT