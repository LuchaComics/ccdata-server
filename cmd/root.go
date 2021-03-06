package cmd

import (
	"fmt"
	"os"

	// homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	// "github.com/spf13/viper"
)

var (
	databaseHost string
	databasePort string
	databaseUser string
	databasePassword string
	databaseName string
	applicationSigningKey string
)


// Initialize function will be called when every command gets called.
func init() {
	// Get our environment variables which will used to configure our application and save across all the sub-commands.
	rootCmd.PersistentFlags().StringVar(&databaseHost, "dbHost", os.Getenv("CCDATA_DB_HOST"), "The address of database.")
	rootCmd.PersistentFlags().StringVar(&databasePort, "dbPort", os.Getenv("CCDATA_DB_PORT"), "The port of database.")
	rootCmd.PersistentFlags().StringVar(&databaseUser, "dbUser", os.Getenv("CCDATA_DB_USER"), "The database user.")
	rootCmd.PersistentFlags().StringVar(&databasePassword, "dbPassword", os.Getenv("CCDATA_DB_PASSWORD"), "The database password.")
	rootCmd.PersistentFlags().StringVar(&databaseName, "dbName", os.Getenv("CCDATA_DB_NAME"), "The database name.")
	rootCmd.PersistentFlags().StringVar(&applicationSigningKey, "appSignKey", os.Getenv("CCDATA_APP_SIGNING_KEY"), "The signing key.")
}

var rootCmd = &cobra.Command{
	Use:   "ccdata-server",
	Short: "",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Do nothing.
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
