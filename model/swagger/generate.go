// Packager Generator based on the swagger API definition
package swagger

//go:generate docker run -v $(pwd):/data swaggerapi/swagger-codegen-cli generate -i /data/swagger_1_8_0.json -l go -o ../../pureport/swagger
//go:generate gofmt -s -w ../../pureport/swagger
