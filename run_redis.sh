#!/bin/bash
docker rm -f demoservice-redis
docker run -d \
    --name demoservice-redis \
    redis
