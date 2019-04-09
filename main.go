// Package main provides CLI Interface for the Pureport SDK
package main

import (
	"github.com/op/go-logging"
	"github.com/pureport-sdk-go/pureport"
	"github.com/pureport-sdk-go/pureport/credentials"
	ppLog "github.com/pureport-sdk-go/pureport/logging"
)

var log = logging.MustGetLogger("main_logger")

func main() {

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

	config := pureport.NewConfiguration().WithCredentials(cred)

	value, err := config.Credentials.Get()
	if err != nil {
		log.Errorf("Error retrieving credentials: %s", err)
	}

	log.Infof("Credentials: %+v", value)
}
