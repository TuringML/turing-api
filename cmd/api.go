package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/turing-ml/turing-api/api"
)

const (
	addressFlag    = "address-host"
	secretFlag     = "secret"
	dbUserFlag     = "db-user"
	dbPasswordFlag = "db-password"
	dbURLFlag      = "db-url"
	dbNameFlag     = "db-name"
	vaultTokenFlag = "vault-token"
	vaultAddrFlag  = "vault-address"
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
	vaultToken := viper.GetString(vaultTokenFlag)
	vaultAddr := viper.GetString(vaultAddrFlag)

	err := api.Serve(addr, secret, dbUsername, dbPassword, dbURL, dbName, vaultToken, vaultAddr)
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
	f.String(dbURLFlag, "192.168.99.100:3306", "database host and port")
	f.String(dbNameFlag, "turing", "database name")
	f.String(vaultTokenFlag, "sloths-are-nice", "vault token used for the calls")
	f.String(vaultAddrFlag, "192.168.99.100", "url of the vault service")

	viper.BindEnv(addressFlag, "ADDRESS_HOST")
	viper.BindEnv(secretFlag, "SECRET")
	viper.BindEnv(dbURLFlag, "DB_URL")
	viper.BindEnv(dbUserFlag, "DB_USER")
	viper.BindEnv(dbPasswordFlag, "DB_PASSWORD")
	viper.BindEnv(dbNameFlag, "DB_NAME")
	viper.BindEnv(vaultTokenFlag, "VAULT_TOKEN")
	viper.BindEnv(vaultAddrFlag, "VAULT_ADDR")

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
