#!/usr/bin/env bash

services=( appointment auth barber rating salon user )

if [[ -z "${TAG}" ]]; then
    TAG="latest"
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
