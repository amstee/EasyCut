# Infrastructure

## Build

The build directory contains every script, template or configuration file needed for our services
deployment 

* The docker directory contains the template of our services Dockerfile.
* The scripts directory contains the files helping us to build and release our code

#### Dockerfile

```dockerfile
FROM alpine:3.8
ARG service
ENV BINARY=$service

RUN adduser -D -g '' appuser

RUN mkdir -p /go/bin/

COPY release/${service} /go/bin/

WORKDIR /go/bin/

USER appuser

ENTRYPOINT /go/bin/$BINARY
```

#### Makefile

```
os = linux
TAG ?= latest
# SERVICE := default

guard-%:
	@ if [ "${${*}}" = "" ]; then \
	echo "Environment variable $* not set"; \
	exit 1; \
	fi

all:	guard-SERVICE release build

.PHONY: release
release:	guard-SERVICE
	mkdir -p release
	CGO_ENABLED=0 GOOS=$(os) GOARCH=amd64 go build -o release/$(SERVICE)

clean:
	rm -rf ./release

build:	guard-SERVICE
	docker build --build-arg service=$(SERVICE) -t amstee/easy-cut-$(SERVICE)":"$(TAG) .
```


## Deployment

This project deployment is scripted for a simple docker environment through docker-compose, 
a docker swarm cluster and a kubernetes cluster.

* The local deployment is really convenient for testing and development
* The docker swarm cluster deployment is really convenient for a fast deployment on premise or on a small cluster hoster locally or on a cheap cloud provider
* The kubernetes cluster deployment is perfect for a production environment where the services need performence and automatic scalability

### Local

The local environment is really convenient if you want to develop and test or use the easy cut api for your own projects.
Its a very lightweight environment.

This deployment is based on `docker-compose` allowing us to deploy and make our services communicate easily.
The environment **TAG** variable is necessary before running the services to indicate docker with image to use.

1. Install docker
2. Install docker-compose
3. Run the command `cd /infra/deploy/local && TAG=**WANTED PROJECT TAG** docker-compose up`
4. Go to `http://localhost/auth/status` and check the service status


### Swarm

The swarm environment is really convenient if you want to stress test the application, make sure nginx run correctly as a reverse-proxy.
Its also perfect is you have a small vps where you can deploy a swarm cluster and run easy cut server.

This deployment is based on `docker swarm` allowing us to deploy, isolate and scale our services easily

If you want to deploy your swarm cluster locally follow this tutorial : https://docs.docker.com/get-started/part4/

If you run swarm on a vm, make sure to run the command below on each cluster node :
`sudo sysctl -w vm.max_map_count=262144`

1. install docker
2. install docker-compose
3. Run `docker swarm init` on your master and read the output
4. Add your nodes to the cluster `docker swarm join --token <token> ...`
5. **If needed :** Configure docker to talk to your master. Example (docker-machine): `eval $(docker-machine env swarm-master)`
6. Add the docker-compose.yml file on the master node. Example (docker-machine): `docker-machine scp -r ./docker-compose.yml swarm-master:/home/docker/docker-compose.yml`
7. Add the nginx configuration file on the master node. Example (docker-machine): `docker-machine scp -r ./nginx.conf swarm-master:/home/docker/nginx.conf`
8. Set the environment variables **API_CLIENT_ID** and **API_CLIENT_SECRET**
9. SSH to your master machine and deploy : `TAG=dev docker stack deploy --compose-file docker-compose.yml easy-cut-dev`


### Kubernetes

The kubernetes environment is perfect for production, it allows us to easily scale our services,
load balance the traffic and safely deploy new services.

The project environment are separated inside the kubernetes cluster user namespaces.
This allow us to concurrently run different project environments on the same cluster.

The first file `**ENV_NAME**.yml` in the k8 directory define the namespace associated to the environment name

There are two configmap definitions which allow us to safeley store the environment variables and configuration
files inside our kubernetes cluster

The `load-balancer.yml` file define our first load balancer service which direct the egress trafic to our nginx pods.

Each EasyCut service has its corresponding service inside the k8 cluster allowing us to load balance the traffic and scale the pods.