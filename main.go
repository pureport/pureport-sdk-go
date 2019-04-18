// Package main provides CLI Interface for the Pureport SDK
package main

import (
	"github.com/op/go-logging"
	"github.com/pureport/pureport-sdk-go/pureport"
	ppLog "github.com/pureport/pureport-sdk-go/pureport/logging"
	"github.com/pureport/pureport-sdk-go/pureport/session"
)

var log = logging.MustGetLogger("main_logger")

func main() {

	cfg := pureport.NewConfiguration()

	logCfg := ppLog.NewLogConfig()
	ppLog.SetupLogger(logCfg)

	s := session.NewSession(cfg)
	ctx := s.GetSessionContext()
	sp, r, err := s.Client.SupportedConnectionsApi.GetSupportedConnections(ctx)

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
