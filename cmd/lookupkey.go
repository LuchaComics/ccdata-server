package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/luchacomics/ccdata-server/internal/session"
	"github.com/luchacomics/ccdata-server/internal/utils"
)

var (
    lkAPIKey string
)

func init() {
	lookupKeyCmd.Flags().StringVarP(&lkAPIKey, "apikey", "k", "", "API Key to process in the command.")
	lookupKeyCmd.MarkFlagRequired("apikey")
	rootCmd.AddCommand(lookupKeyCmd)
}

var lookupKeyCmd = &cobra.Command{
	Use:   "lookupkey",
	Short: "Lookup and print the API key details",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		doLookupKey()
	},
}

func doLookupKey() {
	ctx := context.Background()
    appSignKeyBytes := []byte(applicationSigningKey)
	uuid, err := utils.ProcessJWTToken(appSignKeyBytes, lkAPIKey)
	if err != nil {
		log.Fatal(err)
	}

	// Open up our session handler, powered by redis and let's save the user
	// account with our ID
	sm := session.New()

	user, err := sm.GetUser(ctx, uuid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("\033[H\033[2J")
	fmt.Println("------------\nUSER RECORD\n------------")
	fmt.Println("Id", user.Id)
	fmt.Println("Uuid", user.Uuid)
	fmt.Println("FirstName", user.FirstName)
	fmt.Println("LastName", user.LastName)
	fmt.Println("Email", user.Email)
	fmt.Println("State", user.State)
	fmt.Println("Timezone", user.Timezone)
	fmt.Println("CreatedTime", user.CreatedTime)
	fmt.Println("SessionUuid", user.SessionUuid, "\n")
}
