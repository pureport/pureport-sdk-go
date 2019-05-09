package commandline

import (
	"github.com/pureport/pureport-sdk-go/pureport/logging"
)

// CmdLineParams the data that can be initialized from the command line
type CmdLineParams struct {
	ConfigFile      string
	CredentialsFile string
	Profile         string
	LogConfig       *logging.LogConfig
}

// NewCommandLineParams creates a new Structure to hold the command line parameters
func NewCommandLineParams() *CmdLineParams {
	return &CmdLineParams{
		ConfigFile: "",
		Profile:    "",
		LogConfig:  logging.NewLogConfig(),
	}
}
