#!/bin/zsh

appName="bkst-oauth-api"

docker compose -f docker-context/docker-compose.yml -p "${appName}" up -d

sleep 60

docker container cp resources/oauth.cql "${appName}"_oauth_1:/tmp
docker exec -it "${appName}"_oauth_1 cqlsh -f /tmp/oauth.cql

go fmt ./...
go build ./...

./"${appName}"
