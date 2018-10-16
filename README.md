# Easy Cut

## What's Easy Cut

Easy cut is a personal project that i'm developing with M3thomas.
It's purpose is to reference, rank and rate barbers anywhere around the world.

## Project organization

This project is composed of two applications. 
A server and an IOS application.
This Document is gonna present you the technical details about the server.

You can find the IOS app repo at `https://github.com/Mthomas3/easycuts`

## Server

This repo is composed of two main directories : infra and services

To develop this project we use a set of libraries and tools to make development faster and 
have a more stable product.

`Auth0` allow us to externalize the users authentication and management

`Golang` is the language use to develop our services

`Docker` is used for the project deployment and portability

`Alpine3.8` is the base image used for our containers

`Kubernetes` is used for the project deployment and hosting

`Mux` is the package used for our services routes

`Negroni` for the middleware injection in our services

### [Services](services/README.md)

The project is developed respecting the Microservices development technique.
To resume this technique quickly, each project functionality is implemented by one an only one service.
For example the user session and permissions check is implemented by the service auth.

Each service is located in its own directory matching its name.
The Docker images matches the services names.

Example:
    Service name : `auth` <--> Image name : `easy-cut-auth`

### [Infrastructure](infra/README.md)

The infrastructure directory contains the logic implemented for the project deployment.

The project infrastructure is being developed and maintained in order to be easily deployed
on kubernetes and docker swarm.
