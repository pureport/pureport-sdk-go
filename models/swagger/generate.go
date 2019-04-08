// Packager Generator based on the swagger API definition
package swagger

//go:generate docker run --rm -v $PWD:/data swaggerapi/swagger-codegen-cli generate -i /data/models/swagger/swagger_1_8_0.json -l go -o /data/pureport/swagger
//go:generate gofmt -s -w ../../pureport/swagger
