#!/usr/bin/env bash

services=( appointment auth barber rating salon user perms )
k8conf=( "configmap/auth0.yml" "configmap/nginx.yml" "configmap/ssl.yml" )
k8=( "auth.yml" "user.yml" "nginx.yml" "load-balancer.yml" )

if [[ -z "${1}" ]] || [ $1 = "--help" ]; then
    echo -e "Commands :"
    echo -e "1. install\t --> install go dependencies"
    echo -e "2. dependencies\t --> display go dependencies"
    echo -e "3. deploy\t --> Deploy the project on the specified platform"
    echo -e "  3.1 deploy local\t --> deploy the project locally"
    echo -e "  3.2 deploy swarm\t --> deploy the project on your swarm cluster"
    echo -e "  3.3 deploy k8   \t --> deploy the project on your kubernetes cluster"
    echo -e "4. clean\t --> clean the services build"
    echo -e "5. release\t --> build the services"
    echo -e "6. build\t --> build the docker images corresponding to the services"
    exit
fi

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

if [ -z "${TAG}" ]; then
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
        cd ..
    done
    echo "Cleaning done"
fi

if [ $1 = "release" ]; then
    cd services
    for i in "${services[@]}"
    do
        cd ${i}/
        SERVICE=${i} make release
        cd ..
    done
    echo "Release of all services done"
fi

if [ $1 = "build" ]; then
    cd services
    for i in "${services[@]}"
    do
        cd ${i}/
        TAG="${TAG}" SERVICE=${i} make build
        cd ..
    done
    echo "Build of all services done"
fi

if [ $1 = "push" ]; then
    cd services
    for i in "${services[@]}"
    do
        cd ${i}/
        make push SERVICE=${i} TAG="${TAG}"
        cd ..
    done
    echo "Push of all services done"
fi

if [ $1 = "activate_swarm" ]; then
    eval $(docker-machine env -u)
fi

if [ $1 = "deactivate_swarm" ]; then
    eval eval $(docker-machine env -u)
fi