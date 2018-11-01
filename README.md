# Easy Cut

## Project organization

This project is composed of two applications. 
A server and an IOS application.
This Document is gonna present you the technical details of the server.

You can find the IOS app repo at `https://github.com/Mthomas3/easycuts`

## Fast Deploy

**Local :**

1. Install docker and docker-compose 
2. Run the following command `TAG=dev ./easy-cut.sh deploy local`

And you're done, the easy cut server is now listening over port 80 / 443 on your device !

**Kubernetes :**

1. Make sure your kubectl environment is setted to deploy on your cluster
2. Run the command `TAG=**YOUR TAG** ./easy-cut.sh deploy k8`

To deploy the project on swarm, look [here](infra/README.md)

## Server

This repo is composed of two main directories : infra and services.

To develop this project we use a set of libraries and tools to develop faster and more efficiently.

`Auth0` allow us to externalize the users authentication and management.

`Golang` is the language use to develop our services.

`Docker` is used for the project deployment and portability.

`Alpine3.8` is the base image used for our containers.

`Kubernetes` is used for the project deployment and hosting.

`Swarm` is used as a kubernetes alternative.

`Mux` is the package used for our services routes.

`Negroni` for the middleware injection in our services.

`ElasticSearch` is used for data storage.

`Elastic` is used as an elastic search client.

### [Services](services/README.md)

The project is developed respecting the Microservices development technique.
To resume this technique quickly, each project functionality is implemented by one an only one service.
For example the user session and permissions check is implemented by the service auth and the user creation is implemented by the user service.

Each service is located in its own directory matching its name.
The Docker images matches the services names.

Example:
* Service name : `auth` <--> Image name : `easy-cut-auth`
* Service name : `user` <--> Image name : `easy-cut-user`

### [Infrastructure](infra/README.md)

The infrastructure directory contains the logic implemented for the project deployment.

The project infrastructure is being developed and maintained in order to be easily deployed
on a kubernetes cluster, docker swarm cluster and locally.


### Src directory

The src directory contains common code to all services, you can consider it as a library


### Postman directory

the postman directory contains a postman configuration to quickly test easy cut server endpoints