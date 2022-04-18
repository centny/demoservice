#!/bin/bash
docker rm -f wmservice-srv
docker run -d \
    --name wmservice-srv \
    --link wmservice-postgres:postgres \
    --link wmservice-redis:redis \
    -p 3741:3741 \
    -v /data/wmservice/conf:/srv/conf \
    -v /data/wmservice/upload:/srv/upload \
    wmservice:$1
