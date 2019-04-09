package endpoint

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pureport-sdk-go/pureport"
	"github.com/pureport-sdk-go/pureport/credentials"
)

const providerName = "EndpointCredentialsProvider"

// Provider for endpoint based credential requests
type Provider struct {
	credentials.Expiry
	credentials.Provider

	// https://golang.org/pkg/net/http/
	Client *http.Client

	// HTTP Endpoint to query credentials from
	EndPoint string

	// ExpiryWindow allow the refresh of the authorization token
	// to be refreshed prior to expiration to ensure that we have
	// a valid token.
	ExpiryWindow time.Duration
}

func newEndpointProvider(cfg pureport.Configuration, endpoint string, provider credentials.Provider) credentials.Provider {
	return &Provider{
		Client: &http.Client{
			Timeout: cfg.Timeout,
		},
		EndPoint: endpoint,
		Provider: provider,
	}
}

// NewEndPointCredentials creates a new credentials for requesting and updating
// credentials from a remote endpoint.
func NewEndPointCredentials(cfg pureport.Configuration, endpoint string, provider credentials.Provider) *credentials.Credentials {
	return credentials.NewCredentials(newEndpointProvider(cfg, endpoint, provider))
}

// IsExpired - see Provider.IsExpired()
func (p *Provider) IsExpired() bool {
	return p.Expiry.IsExpired()
}

// Retrieve - see Provider.Retrieve()
func (p *Provider) Retrieve() (credentials.Value, error) {

	local, err := p.Provider.Retrieve()
	if err != nil {
		return credentials.Value{ProviderName: providerName}, err
	}

	// Create the body of the request
	values := map[string]string{"key": local.APIKey, "secret": local.Secret}
	jsonValue, err := json.Marshal(values)
	if err != nil {
		return credentials.Value{ProviderName: providerName}, fmt.Errorf("Error creating credential body")
	}

	buf := bytes.NewBuffer(jsonValue)

	// Create the HTTP Request
	resp, err := p.Client.Post(fmt.Sprintf("%s/login", p.EndPoint), "application/json", buf)
	if err != nil {
		return credentials.Value{ProviderName: providerName}, fmt.Errorf("Error creating credentials login request")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return credentials.Value{ProviderName: providerName}, fmt.Errorf("Error reading credential request body")
	}

	var data map[string]string
	if err = json.Unmarshal(body, &data); err != nil {
		return credentials.Value{ProviderName: providerName}, fmt.Errorf("Error reading credential response")
	}

	// Initialize the Expiry
	expiresIn, err := time.ParseDuration(fmt.Sprintf("%ss", data["expires_in"]))
	if err != nil {
		return credentials.Value{ProviderName: providerName}, fmt.Errorf("Error converting expiry time")
	}

	p.Expiry.SetExpiration(time.Now().Add(expiresIn), p.ExpiryWindow)

	return credentials.Value{
		APIKey:       local.APIKey,
		Secret:       local.Secret,
		SessionToken: data["access_token"],
	}, nil
}
