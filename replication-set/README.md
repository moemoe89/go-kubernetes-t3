# KUBERNETES-REPLICATION-SETS #

This Replication Set project for practice Kubernetes. The go-hostname image container is a simple app that running on port 5000
                                                    
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
$ kubectl apply -f [REPLICATION SET YAML FILE]
$ kubectl apply -f go-hostname-rs.yaml
```

## Check
Check replication set :
```
$ kubectl get rs
```
Describe replication set :
```
$ kubectl describe rs/[REPLICATION SET NAME]
$ kubectl describe rs/go-hostname-rs
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
$ kubectl delete rs [REPLICATION SET NAME]
$ kubectl delete rs go-hostname-rs
```
Delete replication without pod :
```
$ kubectl delete rs [REPLICATION SET NAME] --cascade=false
$ kubectl delete rs go-hostname-rs --cascade=false
```

## License

MIT