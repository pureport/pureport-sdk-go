package pureport

import (
	"fmt"
	"os"
	"os/user"
	"time"
)

const (
	// ConfFileEnvStr - environment variable for the path to the
	// Pureport configuration file to use.
	ConfFileEnvStr string = "PUREPORT_CONFIG_FILE"
)

// Configuration - Application level configuration data
type Configuration struct {
	Filename string

	// Timeout used for any HTTP Requests made
	Timeout time.Duration

	// The endpoint to request Authentication Tokens
	EndPoint string
}

// NewConfiguration creates a new configuration to application
// level configuration data.
func NewConfiguration(filename string) *Configuration {
	return &Configuration{
		Filename: filename,
		Timeout:  (time.Minute * 2),
		EndPoint: "http://api.pureport.com",
	}
}

func (c *Configuration) filename() (string, error) {

	if len(c.Filename) != 0 {
		return c.Filename, nil
	}

	if c.Filename = os.Getenv(ConfFileEnvStr); len(c.Filename) != 0 {
		return c.Filename, nil
	}

	usr, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("Users' home directory not found. Can't initialize configuration")
	}

	defaultPath := usr.HomeDir + "/.pureport/configuration.yml"
	if _, err = os.Stat(defaultPath); err != nil {
		return "", fmt.Errorf("Configuration file not found")
	}

	return defaultPath, nil
}

// Load the configuration data
func (c *Configuration) Load() error {
	return nil
}

// WithEndPoint - add the Authorization endpoint to the configuration
func (c *Configuration) WithEndPoint(endpoint string) *Configuration {
	c.EndPoint = endpoint
	return c
}
