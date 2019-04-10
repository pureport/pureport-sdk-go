// Package main provides CLI Interface for the Pureport SDK
package main

import (
	"github.com/op/go-logging"
	"github.com/pureport-sdk-go/pureport"
	"github.com/pureport-sdk-go/pureport/credentials"
	"github.com/pureport-sdk-go/pureport/credentials/endpoint"
	ppLog "github.com/pureport-sdk-go/pureport/logging"
)

var log = logging.MustGetLogger("main_logger")

func main() {

	config := pureport.NewConfiguration()
	config = config.WithEndPoint("https://dev1-api.pureportdev.com/login")

	logConfig := ppLog.NewLogConfig()
	ppLog.SetupLogger(logConfig)

	providers := []credentials.Provider{
		&credentials.EnvironmentProvider{},
		&credentials.FileProvider{
			Filename: "",
			Profile:  "",
		},
	}
	cred := credentials.NewChainCredentials(providers)

	endpointCred := endpoint.NewEndPointCredentials(*config, config.EndPoint, cred)

	value, err := endpointCred.Get()
	if err != nil {
		log.Errorf("Error retrieving credentials: %s", err)
	}

	log.Infof("Credentials: %+v", value)
}
