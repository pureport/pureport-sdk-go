// Package main provides CLI Interface for the Pureport SDK
package main

import (
	"github.com/pureport-sdk-go/pureport/credentials"
	"log"
)

func main() {

	providers := []credentials.Provider{
		&credentials.EnvironmentProvider{},
		&credentials.FileProvider{
			Filename: "",
			Profile:  "",
		},
	}

	cred := credentials.NewChainCredentials(providers)

	value, err := cred.Get()
	if err != nil {
		log.Fatalf("Error retrieving credentials: %s", err)
	}

	log.Printf("Credentials: %+v", value)

}
