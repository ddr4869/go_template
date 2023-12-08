#!/bin/bash

action=$1
tag=$2

case $action in
    "build")
        docker build -t go-board:$tag .
        ;;
    "tag")
        docker tag go-board:$tag spspid/go-board:$tag
        ;;
    "push")
        docker push spspid/go-board:$tag
        ;;
    "all")
        docker build -t go-board:$tag .
        docker tag go-board:$tag spspid/go-board:$tag
        docker push spspid/go-board:$tag
        ;;
    *)
        echo "Choose one in [build|tag|push|all]"
        exit 1
        ;;
esac