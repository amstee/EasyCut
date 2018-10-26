#!/usr/bin/env bash

services=( appointment auth barber rating salon user )
k8conf=( "configmap/auth0.yml" "configmap/nginx.yml" "configmap/ssl.yml" )
k8=( "auth.yml" "user.yml" "nginx.yml" "load-balancer.yml" )

if [ $1 = "install" ]; then
    go get github.com/gorilla/mux &
    go get github.com/urfave/negroni &
    go get github.com/rs/cors &
    go get github.com/auth0/go-jwt-middleware &
    go get github.com/spf13/viper &
    go get github.com/dgrijalva/jwt-go &
    go get github.com/olivere/elastic &
    exit
fi

if [ $1 = "dependencies" ]; then
    echo "github.com/gorilla/mux"
    echo "github.com/urfave/negroni"
    echo "github.com/rs/cors"
    echo "github.com/auth0/go-jwt-middleware"
    echo "github.com/spf13/viper"
    echo "github.com/dgrijalva/jwt-go"
    echo "github.com/olivere/elastic"
    exit
fi

if [[ -z "${TAG}" ]]; then
    TAG="latest"
    echo "TAG not found in environment --> setting to default value : latest"
fi

if [ $1 = "deploy" ]; then
    if [ $2 = "local" ]; then
        cd "infra/deploy/local"
        TAG=${TAG} docker-compose up
    elif [ $2 = "swarm" ]; then
    	cd "infra/deploy/swarm"
    	TAG=${TAG} docker stack deploy --compose-file docker-compose.yml easy-cut-${TAG}
    elif [ $2 = "k8" ]; then
        cd "infra/deploy/k8/"
        if [ "$TAG" == "latest" ] || [ "$TAG" == "dev" ]; then
            kubctl apply -f dev.yml
            cd "dev"
        elif [ "$TAG" == "prod" ]; then
            kubctl apply -f prod.yml
            cd "prod"
        else
            echo "Deployment associated to $TAG not found"
            exit
        fi
        for c in "${k8conf[@]}"
        do
            kubectl apply -f ${c}
        done
        for k in "${k8[@]}"
        do
            kubectl apply -f ${k}
        done
    else
        echo "Command deploy ${2} unknown"
        echo "Possible deployments : local, swarm, k8"
    fi
fi

if [ $1 = "clean" ]; then
    cd services
    for i in "${services[@]}"
    do
        cd ${i}/
        make clean
    done
    echo "Cleaning done"
fi

if [ $1 = "release" ]; then
    cd services
    for i in "${services[@]}"
    do
        cd ${i}/
        make release SERVICE=${i}
    done
    echo "Release of all services done"
fi

if [ $1 = "build" ]; then
    cd services
    for i in "${services[@]}"
    do
        cd ${i}/
        make build SERVICE=${i} TAG="${TAG}"
    done
    echo "Build of all services done"
fi

if [ $1 = "activate_swarm" ]; then
    eval $(docker-machine env -u)
fi

if [ $1 = "deactivate_swarm" ]; then
    eval eval $(docker-machine env -u)
fi