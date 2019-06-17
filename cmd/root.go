package cmd

import (
	"github.com/op/go-logging"
	"github.com/pureport/pureport-sdk-go/pureport"
	"github.com/pureport/pureport-sdk-go/pureport/client"
	ppLog "github.com/pureport/pureport-sdk-go/pureport/logging"
	"github.com/pureport/pureport-sdk-go/pureport/session"
	"github.com/spf13/cobra"
)

var log = logging.MustGetLogger("main_logger")
var logCfg = ppLog.NewLogConfig()

var cliSession *session.Session

var rootCmd = &cobra.Command{
	Use:   "pp-cli",
	Short: "Pureport CLI interface",
	Long:  `Command line interface for the Pureport Cloud Fabric Service.`,
}

var cmdViewAccount = &cobra.Command{
	Use:   "account",
	Short: "View your account information",
	Long:  "Command to view your current account information",
	Run: func(cmd *cobra.Command, args []string) {

		ctx := cliSession.GetSessionContext()

		opts := &client.FindAllAccountsOpts{}
		sp, r, err := cliSession.Client.AccountsApi.FindAllAccounts(ctx, opts)

		if err != nil {
			log.Info(err)
			log.Fatalf("Error while querying SupportedConnections.")
		}

		if r.StatusCode != 200 {
			log.Error(r)
			log.Error(sp)
		} else {
			log.Info(r)
			log.Info(sp)
		}
	},
}

func init() {

	cobra.OnInitialize(initConfig)

	cfg := pureport.NewConfiguration()
	cliSession = session.NewSession(cfg)
	cliSession.Credentials.Get()

	rootCmd.AddCommand(cmdViewAccount)
	rootCmd.AddCommand(cmdSwagger)
}

func initConfig() {

	logCfg.Level = "debug"
	ppLog.SetupLogger(logCfg)

}

// Execute is the main interface for the CLI Command interpreter
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Panic(err)
	}
}
