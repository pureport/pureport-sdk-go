// Package swagger Generator based on the swagger API definition
package swagger

//go:generate docker run --rm -v $PWD:/data swaggerapi/swagger-codegen-cli:2.4.9 generate -i /data/models/swagger/pureport_swagger.json -l go -o /data/pureport/client -c /data/models/swagger/config.json --invoker-package pureport\client
//go:generate gofmt -s -w ../../pureport/client
