package pureport

import (
	"github.com/pureport-sdk-go/pureport/credentials"
)

const (
	// ConfFileEnvStr - environment variable for the path to the
	// Pureport configuration file to use.
	ConfFileEnvStr string = "PUREPORT_CONFIG_FILE"
)

// Configuration - Application level configuration data
type Configuration struct {
	Credentials *credentials.Credentials `desc:"The credentials for authenticating with Pureport control plane"`
}

// NewConfiguration creates a new configuration to application
// level configuration data.
func NewConfiguration() *Configuration {
	return &Configuration{}
}

// WithCredentials set the credentials provider to use
func (c *Configuration) WithCredentials(credentials *credentials.Credentials) *Configuration {
	c.Credentials = credentials
	return c
}

// WithConfigFile load configuration data from the specified file
func (c *Configuration) WithConfigFile(filename string) *Configuration {
	return c
}
