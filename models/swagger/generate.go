// Package swagger Generator based on the swagger API definition
package swagger

//go:generate docker run --rm -v $PWD:/data swaggerapi/swagger-codegen-cli generate -i /data/models/swagger/kato_swagger.json -l go -o /data/pureport/swagger
//go:generate gofmt -s -w ../../pureport/swagger
//go:generate mv ../../pureport/swagger/docs ../../docs/swagger
//go:generate mv ../../pureport/swagger/README.md ../../docs/README.md
