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

	cfg := pureport.NewConfiguration("")
	cfg = cfg.WithEndPoint("https://dev1-api.pureportdev.com")

	logCfg := ppLog.NewLogConfig()
	ppLog.SetupLogger(logCfg)

	s := session.NewSession(cfg)

	s.SomeRequest()
}
