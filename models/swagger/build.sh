#!/bin/sh

docker run --rm -v $PWD:/data openapitools/openapi-generator-cli:v4.2.0 generate -i /data/models/swagger/pureport_swagger.json -g go -o /data/pureport/client -c /data/models/swagger/config.json --enable-post-process-file -p packageName='pureport' -p packageVersion=1.2.1
