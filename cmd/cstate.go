package cmd

import (
	"context"
	"fmt"
	"log"
	// "time"

	"github.com/spf13/cobra"

	repo "github.com/luchacomics/ccdata-server/internal/repositories"
	"github.com/luchacomics/ccdata-server/internal/session"
	"github.com/luchacomics/ccdata-server/internal/utils"
)

var (
	csEmail string
	csState int
)

func init() {
	csCmd.Flags().StringVarP(&genkeyEmail, "email", "f", "", "Email of the user account")
	csCmd.MarkFlagRequired("email")
	csCmd.Flags().IntVarP(&csState, "state", "s", 0, "State of the user account")
	csCmd.MarkFlagRequired("state")
	rootCmd.AddCommand(csCmd)
}

var csCmd = &cobra.Command{
	Use:   "cstate",
	Short: "Change user state",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		doRunChangeUserState()
	},
}

func doRunChangeUserState() {
	ctx := context.Background()

	// Load up our database.
	db, err := utils.ConnectDB(databaseHost, databasePort, databaseUser, databasePassword, databaseName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Load up our data access later
	r := repo.NewUserRepo(db)

	user, err := r.GetByEmail(ctx, genkeyEmail)
	if err != nil {
		log.Fatal(err)
	}

	// Update state in our database.
	user.State = int8(csState)
	err = r.Update(ctx, user)
	if err != nil {
		log.Fatal(err)
	}

	// Open up our session handler, powered by redis and let's save the user.
	sm := session.New()
	err = sm.SaveUser(ctx, user.SessionUuid, user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("\033[H\033[2J")
	fmt.Println("User state modified.")
}
