// Package swagger Generator based on the swagger API definition
package swagger

//go:generate docker run --rm -v $PWD:/data swaggerapi/swagger-codegen-cli generate -i /data/models/swagger/kato_swagger.json -l go -o /data/pureport/swagger -c /data/models/swagger/config.json
//go:generate gofmt -s -w ../../pureport/swagger
