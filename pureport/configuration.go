package pureport

import (
	"time"
)

const (
	// ConfFileEnvStr - environment variable for the path to the
	// Pureport configuration file to use.
	ConfFileEnvStr string = "PUREPORT_CONFIG_FILE"
)

// Configuration - Application level configuration data
type Configuration struct {

	// Timeout used for any HTTP Requests made
	Timeout time.Duration

	// The endpoint to request Authentication Tokens
	EndPoint string
}

// NewConfiguration creates a new configuration to application
// level configuration data.
func NewConfiguration(filename string) *Configuration {
	return &Configuration{
		Timeout: (time.Minute * 2),
	}
}

// WithEndPoint - add the Authorization endpoint to the configuration
func (c *Configuration) WithEndPoint(endpoint string) *Configuration {
	c.EndPoint = endpoint
	return c
}
