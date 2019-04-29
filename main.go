// Package main provides CLI Interface for the Pureport SDK
package main

import (
	"github.com/op/go-logging"
	"github.com/pureport/pureport-sdk-go/pureport"
	ppLog "github.com/pureport/pureport-sdk-go/pureport/logging"
	"github.com/pureport/pureport-sdk-go/pureport/session"
	"github.com/pureport/pureport-sdk-go/pureport/swagger"
)

var log = logging.MustGetLogger("main_logger")

func main() {

	cfg := pureport.NewConfiguration()

	logCfg := ppLog.NewLogConfig()
	ppLog.SetupLogger(logCfg)

	s := session.NewSession(cfg)
	ctx := s.GetSessionContext()

	opts := &swagger.FindAllAccountsOpts{}
	sp, r, err := s.Client.AccountsApi.FindAllAccounts(ctx, opts)

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
