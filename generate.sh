#!/bin/sh
# This script will codegen a go package for a server stub or client.


docker run \
    --rm \
    --user $(id -u):$(id -g) \
    -v ${PWD}:/local \
    openapitools/openapi-generator-cli:v5.0.0 generate \
        -i /local/swagger.yml \
        -g go \
        -o /local/models \
        -p enumClassPrefix=true \
        -p skipFormModel=true
rm models/api_*.go \
    models/.travis.yml \
    models/client.go \
    models/configuration.go \
    models/git_push.sh \
    models/go.mod \
    models/go.sum \
    models/response.go \
    models/api/openapi.yaml \
    models/docs/*Api.md
