package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/turing-ml/turing-api/app"
)

const (
	addressFlag    = "address-host"
	secretFlag     = "secret"
	dbUserFlag     = "db-user"
	dbPasswordFlag = "db-password"
	dbURLFlag      = "db-url"
	dbNameFlag     = "db-name"
)

var cmdAPI = &cobra.Command{
	Use:   "api",
	Short: "Run the API server",
	Long: `
TuringML API
This app is used to expose the endpoints
for running the TuringML Application
`,
	SilenceUsage: true,
	Run:          runAPI,
}

func runAPI(cmd *cobra.Command, args []string) {

	addr := viper.GetString(addressFlag)
	secret := viper.GetString(secretFlag)
	dbUsername := viper.GetString(dbUserFlag)
	dbPassword := viper.GetString(dbPasswordFlag)
	dbURL := viper.GetString(dbURLFlag)
	dbName := viper.GetString(dbNameFlag)

	a, err := app.NewApp(secret, dbURL, dbUsername, dbPassword, dbName)
	if err != nil {
		log.Panic().Err(err).Msg("could not create app")
	}

	err = a.Serve(addr)
	if err != nil {
		log.Panic().Err(err).Msg("could not serve the application")
	}
}

func init() {
	f := cmdAPI.Flags()

	f.String(addressFlag, ":8000", "server address")
	f.String(secretFlag, "sloth", "authentication secret for jwt")
	f.String(dbUserFlag, "turing", "database username")
	f.String(dbPasswordFlag, "turing", "database password")
	f.String(dbURLFlag, "mongo", "database host and port")
	f.String(dbNameFlag, "turing", "database name")

	viper.BindEnv(addressFlag, "ADDRESS_HOST")
	viper.BindEnv(secretFlag, "SECRET")
	viper.BindEnv(dbURLFlag, "DB_URL")
	viper.BindEnv(dbUserFlag, "DB_USER")
	viper.BindEnv(dbPasswordFlag, "DB_PASSWORD")
	viper.BindEnv(dbNameFlag, "DB_NAME")

	viper.BindPFlags(f)
}

// initConfig sets AutomaticEnv in viper to true.
func initConfig() {
	viper.AutomaticEnv() // read in environment variables that match
}

// Execute will start the application
func Execute() {
	cobra.OnInitialize(initConfig)
	if err := cmdAPI.Execute(); err != nil {
		log.Fatal().Err(err).Msg("")
	}
}
