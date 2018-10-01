#!/bin/bash

docker pull postgres:9.6

docker rm -f test-postgres
docker run --name test-postgres -e POSTGRES_PASSWORD=postgres -d -p 5432:5432 postgres:9.6