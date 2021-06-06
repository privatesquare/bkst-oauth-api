#!/bin/sh

docker container cp oauth.cql docker-context_oauth_1:/tmp
docker exec -it docker-context_oauth_1 cqlsh -f /tmp/oauth.cql
