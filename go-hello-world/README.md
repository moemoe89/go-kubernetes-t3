# GO-HELLO-WORLD #

This Go project for practice Kubernetes and will be built in Docker image for storing in Docker Hub

## Setup

* Setup Golang <https://golang.org/>
* Setup Docker <https://www.docker.io/>
* Setup Account in Docker Hub <https://hub.docker.com/>

## Build Image
Make image :
```
$ docker build -t momo89/go-hello-world .
```

## Push to Docker Hub
Login to Docker :
```
$ docker login
```
Push to Docker Hub :
```
$ docker push momo89/go-hello-world:latest
```

## Note
momo89 is my account, please change the account if you want to use your own


## License

MIT