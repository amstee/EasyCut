#!/usr/bin/env bash

services=( appointment auth barber rating salon user )

if [[ -z "${TAG}" ]]; then
    TAG="latest"
fi

if [ $1 = "deploy" ]; then
    if [ $2 = "local" ]; then
        cd "infra/deploy/local"
        TAG=${TAG} docker-compose up
    else
        echo "Commad deploy ${2} unknown"
    fi
fi

if [ $1 = "clean" ]; then
    cd services
    for i in "${services[@]}"
    do
        cd $i/
        make clean
        cd ..
    done
    echo "Cleaning done"
fi

if [ $1 = "release" ]; then
    cd services
    for i in "${services[@]}"
    do
        cd $i/
        make release SERVICE=$i
        cd ..
    done
    echo "Release of all services done"
fi

if [ $1 = "build" ]; then
    cd services
    for i in "${services[@]}"
    do
        cd $i/
        make build SERVICE=$i TAG="${TAG}"
        cd ..
    done
    echo "Build of all services done"
fi
