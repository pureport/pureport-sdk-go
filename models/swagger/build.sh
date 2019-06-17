#!/bin/sh

docker run --rm -v $PWD:/data swaggerapi/swagger-codegen-cli:2.4.5 generate -i /data/models/swagger/kato_swagger.json -l go -o /data/pureport/client -c /data/models/swagger/config.json --invoker-package 'pureport'
