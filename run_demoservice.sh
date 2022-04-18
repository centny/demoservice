#!/bin/bash
docker rm -f demoservice-srv
docker run -d \
    --name demoservice-srv \
    --link demoservice-postgres:postgres \
    --link demoservice-redis:redis \
    -p 10808:8080 \
    -v /data/demoservice/conf:/srv/conf \
    demoservice:$1
