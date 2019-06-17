package cmd

import (
	"io"
	"net/http"
	"os"

	"github.com/pureport/pureport-sdk-go/pureport/client"
	"github.com/spf13/cobra"
)

var cmdSwagger = &cobra.Command{
	Use:   "swagger",
	Short: "Download Swagger",
	Long:  "Command to download the Swagger Schema for the REST API",
	Run: func(cmd *cobra.Command, args []string) {

		// Create the file
		out, err := os.Create("./pureport_swagger.json")
		if err != nil {
			log.Panicf("Error creating local file: %s", err)
		}
		defer out.Close()

		ctx := cliSession.GetSessionContext()
		cl := &http.Client{}

		req, err := http.NewRequest("GET", cliSession.Configuration.EndPoint+"/swagger.json", nil)

		// AccessToken Authentication
		if auth, ok := ctx.Value(client.ContextAccessToken).(string); ok {
			req.Header.Add("Authorization", "Bearer "+auth)
		}

		resp, err := cl.Do(req)
		if err != nil {
			log.Panicf("Error downloading Swagger Schema: %s", err)
		}
		defer resp.Body.Close()

		if _, err = io.Copy(out, resp.Body); err != nil {
			log.Panicf("Error reading Swagger Schema: %s", err)
		}

	},
}
