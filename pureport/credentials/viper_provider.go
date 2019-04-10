package credentials

import (
	"fmt"
	"github.com/spf13/viper"
)

const (
	// ViperProviderName is the name of this provider
	ViperProviderName = "ViperProvider"

	configAPIKeyStr    = "api_key"
	configAPISecretStr = "api_secret"
)

var (
	vip *viper.Viper
)

// ViperProvider is a credentials provider using Viper as the configuration source.
type ViperProvider struct {
	retrieved bool

	// Filename - path to the configuration file (Optional)
	Filename string

	// Profile - the profile to use from the configuration file (Optional)
	Profile string
}

func init() {
	vip = viper.New()
	vip.SetConfigName("credentials")
	vip.AddConfigPath("$HOME/.pureport")
	vip.AddConfigPath(".")
	vip.SetEnvPrefix("PUREPORT")
}

// NewViperCredentials creates a new credentials provider using Viper
func NewViperCredentials(profile string) *Credentials {

	if profile == "" {
		profile = "default"
	}

	return NewCredentials(&ViperProvider{
		retrieved: false,
		Profile:   profile,
	})
}

// Retrieve Provider.Retrieve()
func (p *ViperProvider) Retrieve() (Value, error) {
	p.retrieved = false

	if err := vip.ReadInConfig(); err != nil {
		return Value{ProviderName: ViperProviderName}, fmt.Errorf("Error reading in configuration file: %s", err)
	}

	// Check environment first
	key := vip.GetString(configAPIKeyStr)
	secret := vip.GetString(configAPISecretStr)

	// Read from configuration file
	if key == "" || secret == "" {
		key = vip.GetString(fmt.Sprintf("profiles.%s.%s", p.Profile, configAPIKeyStr))
		secret = vip.GetString(fmt.Sprintf("profiles.%s.%s", p.Profile, configAPISecretStr))
	}

	if key == "" || secret == "" {
		return Value{ProviderName: ViperProviderName}, fmt.Errorf("API Key and/or Secret not found")
	}

	p.retrieved = true
	return Value{
		ProviderName: ViperProviderName,
		APIKey:       key,
		Secret:       secret,
	}, nil
}

// IsExpired Provider.IsExpired()
func (p *ViperProvider) IsExpired() bool {
	return !p.retrieved
}
