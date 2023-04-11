#!/bin/bash

set -e

docker pull nats:2.9-alpine
docker pull postgres:15-alpine
docker pull redis:7-alpine

docker run \
    -d \
    --name nats \
    -p 4222:4222 \
    nats:2.9-alpine

docker run \
    -d \
    --name postgres \
    -e POSTGRES_HOST_AUTH_METHOD=trust \
    -p 5432:5432 \
    -v $PWD/.devcontainer/sql:/docker-entrypoint-initdb.d \
    postgres:15-alpine

docker run -d \
    --name redis \
    -p 6379:6379 \
    redis:7-alpine


echo 'PS1="$ "' >> ~/.bashrc
