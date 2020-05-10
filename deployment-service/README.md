# KUBERNETES-DEPLOYMENT-SERVICES #

This Deployment Service project for practice Kubernetes. The go-hostname image container is a simple app that running on port 5000

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

## Apply
Apply deployment :
```
$ kubectl apply -f [DEPLOYMENT YAML FILE]
$ kubectl apply -f go-hostname-deployment.yaml
```
Apply service :
```
$ kubectl apply -f [SERVICE YAML FILE]
$ kubectl apply -f go-hostname-service.yaml
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
$ kubectl describe pod go-hostname-{UNIQUE-ID}
```
Check deployment :
```
$ kubectl get deployment
```
Describe deployment :
```
$ kubectl describe deployment/[DEPLOYMENT NAME]
$ kubectl describe deployment/go-hostname-deployment
```
Check service :
```
$ kubectl get service
```
Check service more detail :
```
$ kubectl get service -o wide
```
Describe service :
```
$ kubectl describe service [SERVICE NAME]
$ kubectl describe service go-hostname-service
```

## Access
Get minikube ip :
```
minikube ip
```
We will get ip let say:
```
192.168.99.100
```
Navigate browser to this address :
```
[MINI KUBE IP]:[NODE PORT]
192.168.99.100:31515
```

## Delete
Delete pod :
```
$ kubectl delete pod [POD NAME]
$ kubectl delete pod go-hostname-{UNIQUE-ID}
```
Delete deployment :
```
$ kubectl delete deployment [DEPLOYMENT NAME]
$ kubectl delete deployment go-hostname-deployment
```
Delete service :
```
$ kubectl delete svc [SERVICE NAME]
$ kubectl delete svc go-hostname-service
```

## License

MIT