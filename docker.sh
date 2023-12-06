#!/bin/bash

action=$1
tag=$2

if [ "$action" == "build" ]; then
    docker build -t go-board:$tag .
fi
if [ "$action" == "tag" ]; then
    docker tag go-board:$tag spspid/go-board:$tag
fi
if [ "$action" == "push" ]; then
    docker push spspid/go-board:$tag
fi