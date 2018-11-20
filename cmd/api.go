package cmd

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/turing-ml/turing-api/app"
	"os"
)

var cmdApi = &cobra.Command{
	Use:   "server",
	Short: "Run the API server",
	Long: `
TuringML API
This app is used to expose the endpoints
for running the TuringML Application
`,
	SilenceUsage: true,
	Run:          runApi,
}

func runApi(cmd *cobra.Command, args []string) {

	addr, _ := cmd.Flags().GetString("addr")
	secret, _ := cmd.Flags().GetString("secret")
	dbUsername, _ := cmd.Flags().GetString("dbUsername")
	dbPassword, _ := cmd.Flags().GetString("dbPassword")
	dbUrl, _ := cmd.Flags().GetString("dbUrl")
	dbName, _ := cmd.Flags().GetString("dbName")

	a, err := app.NewApp(secret, dbUrl, dbUsername, dbPassword, dbName)
	if err != nil {
		log.Panic().Err(err).Msg("could not create app")
	}

	err = a.Serve(addr)
	if err != nil {
		log.Panic().Err(err).Msg("could not serve the application")
	}
}

func init() {
	cmdApi.Flags().StringP("addr", "a", ":8000", "server address")
	cmdApi.Flags().StringP("secret", "s", "sloth", "authentication secret for jwt")
	cmdApi.Flags().StringP("dbUsername", "u", "greeny", "database username")
	cmdApi.Flags().StringP("dbPassword", "p", "greeny", "database password")
	cmdApi.Flags().StringP("dbUrl", "l", "192.168.99.100:3306", "database host and port")
	cmdApi.Flags().StringP("dbName", "n", "greeny", "database name")
}

func Execute() {
	if err := cmdApi.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}