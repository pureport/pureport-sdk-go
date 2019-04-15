package session

import (
	"context"
	"os"

	"github.com/op/go-logging"
	"github.com/pureport/pureport-sdk-go/pureport"
	"github.com/pureport/pureport-sdk-go/pureport/credentials"
	"github.com/pureport/pureport-sdk-go/pureport/credentials/endpoint"
	"github.com/pureport/pureport-sdk-go/pureport/swagger"
)

var log = logging.MustGetLogger("main_logger")

// Session contains the data for a particular request session
type Session struct {
	*credentials.Credentials
	*pureport.Configuration

	Client *swagger.APIClient
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

	cred := credentials.NewViperCredentials("")

	return endpoint.NewEndPointCredentials(*cfg, cfg.EndPoint, cred)
}

// NewSession creates a new request session
func NewSession(cfg *pureport.Configuration) *Session {

	return &Session{
		Credentials:   createCredentials(cfg),
		Configuration: cfg,
		Client:        createClient(cfg),
	}
}

// GetSessionContext gathers the context information need to
// for communicating with the Pureport API
func (s *Session) GetSessionContext() context.Context {

	value, err := s.Credentials.Get()
	if err != nil {
		log.Fatalf("Retrieving access credentials failed: %s", err)
	}

	return context.WithValue(context.Background(), swagger.ContextAccessToken, value.SessionToken)
}
