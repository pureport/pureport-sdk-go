package session

import (
	"context"
	"os"

	"github.com/op/go-logging"
	"github.com/pureport-sdk-go/pureport"
	"github.com/pureport-sdk-go/pureport/credentials"
	"github.com/pureport-sdk-go/pureport/credentials/endpoint"
	"github.com/pureport-sdk-go/pureport/swagger"
)

var log = logging.MustGetLogger("main_logger")

// Session contains the data for a particular request session
type Session struct {
	*credentials.Credentials
	*pureport.Configuration

	client *swagger.APIClient
}

func createClient(cfg *pureport.Configuration) *swagger.APIClient {
	c := swagger.NewConfiguration()
	c.BasePath = cfg.EndPoint
	//c.AddDefaultHeader()

	if hostname, err := os.Hostname(); err != nil {
		c.Host = ""
	} else {
		c.Host = hostname
	}

	return swagger.NewAPIClient(c)
}

func createCredentials(cfg *pureport.Configuration) *credentials.Credentials {

	providers := []credentials.Provider{
		&credentials.EnvironmentProvider{},
		&credentials.FileProvider{
			Filename: "",
			Profile:  "",
		},
	}
	cred := credentials.NewChainCredentials(providers)

	return endpoint.NewEndPointCredentials(*cfg, cfg.EndPoint, cred)
}

// NewSession creates a new request session
func NewSession(cfg *pureport.Configuration) *Session {

	return &Session{
		Credentials:   createCredentials(cfg),
		Configuration: cfg,
		client:        createClient(cfg),
	}
}

// SomeRequest - Stub request function
func (s *Session) SomeRequest() {

	value, err := s.Credentials.Get()
	if err != nil {
		log.Fatalf("Retrieving access credentials failed")
	}

	ctx := context.WithValue(context.Background(), swagger.ContextAccessToken, value.SessionToken)

	sp, r, err := s.client.SupportedConnectionsApi.GetSupportedConnections(ctx)

	if err != nil {
		log.Info(r)
		log.Fatalf("Error while querying SupportedConnections.")
	}

	if r.StatusCode != 200 {
		log.Error(r)
		log.Error(sp)
	} else {
		log.Info(r)
		log.Info(sp)
	}
}
