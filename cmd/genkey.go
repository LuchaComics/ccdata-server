package cmd

import (
	"context"
	"fmt"
	"log"
	"time"
	"os"

	"github.com/spf13/cobra"

	repo "github.com/luchacomics/ccdata-server/internal/repositories"
	"github.com/luchacomics/ccdata-server/internal/session"
	"github.com/luchacomics/ccdata-server/internal/utils"
)

var (
	genkeyEmail string
)

func init() {
	genkeyCmd.Flags().StringVarP(&genkeyEmail, "email", "f", "", "Email of the user account")
	genkeyCmd.MarkFlagRequired("email")
	rootCmd.AddCommand(genkeyCmd)
}

var genkeyCmd = &cobra.Command{
	Use:   "genkey",
	Short: "Generate and print the API key of user",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		doRunGenerateAPIKey()
	},
}

func doRunGenerateAPIKey() {
	ctx := context.Background()

	// Load up our database.
	db, err := utils.ConnectDB(databaseHost, databasePort, databaseUser, databasePassword, databaseName)
	if err != nil {
		log.Fatal("ConnectDB:", err)
	}
	defer db.Close()

	// Load up our data access later
	r := repo.NewUserRepo(db)

	user, err := r.GetByEmail(ctx, genkeyEmail)
	if err != nil {
		log.Fatal(err)
	}

    appSignKeyBytes := []byte(applicationSigningKey)

    // Note:
	// (1) 1 hour x 24 = 1 day
	// (2) Today is : Tuesday, February 9, 2021.
	//     The date after 100000 days is : Sunday, November 25, 2294
	//     SOURCE: https://whatisthedatetoday.com/100000-days-from-today.html
	afterManyYears := time.Hour * 24 * 100000

	// Open up our session handler, powered by redis and let's save the user
	// account with our ID
	sm := session.New()
	err = sm.SaveUser(ctx, user.SessionUuid, user)
	if err != nil {
		log.Fatal("SaveUser:", err)
	}

    // Generate our one-time token
	accessToken, _, err := utils.GenerateJWTTokenPair(appSignKeyBytes, user.SessionUuid, afterManyYears)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(accessToken, "\n")

	// Save to the user's environment.
	os.Setenv("CCDATA_APP_ACCESS_TOKEN", accessToken)

	// Output message.
	fmt.Printf("First run in your console:\n\nexport CCDATA_APP_ACCESS_TOKEN=%s\n\n", accessToken)
	fmt.Printf("Then run in your console:\n\nhttp get 127.0.0.1:5000/api/v1/version \"Authorization:Bearer $CCDATA_APP_ACCESS_TOKEN\"\n\n")
}
