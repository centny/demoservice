#!/bin/bash
docker rm -f demoservice-postgres
docker run -d \
    --name demoservice-postgres \
    -e POSTGRES_DB=demoservice \
    -e POSTGRES_USER=demoservice \
    -e POSTGRES_PASSWORD=123 \
    -e PGDATA=/var/lib/postgresql/data/pgdata \
    -e POSTGRES_HOST_AUTH_METHOD=md5 \
    -v /data/demoservice/postgres/:/var/lib/postgresql/data \
    postgres:latest