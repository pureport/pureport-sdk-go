package pureport

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
)

var (
	vip *viper.Viper
)

const (
	endpointEnvStr = "endpoint"
)

func init() {
	vip = viper.New()
	vip.SetConfigName("configuration")
	vip.AddConfigPath("$HOME/.pureport")
	vip.AddConfigPath(".")

	vip.SetEnvPrefix("pureport")
	vip.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	vip.BindEnv(endpointEnvStr)

	vip.SetDefault(endpointEnvStr, "http://api.pureport.com")
}

// Configuration - Application level configuration data
type Configuration struct {

	// Timeout used for any HTTP Requests made
	Timeout time.Duration

	// The endpoint to request Authentication Tokens
	EndPoint string
}

// NewConfiguration creates a new configuration to application
// level configuration data.
func NewConfiguration() *Configuration {

	if err := vip.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Error reading configuration file: %s", err))
	}

	return &Configuration{
		Timeout:  (time.Minute * 2),
		EndPoint: vip.GetString(endpointEnvStr),
	}
}
