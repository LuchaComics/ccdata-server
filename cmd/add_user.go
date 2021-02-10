package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
	"github.com/google/uuid"

	"github.com/luchacomics/ccdata-server/internal/models"
	repo "github.com/luchacomics/ccdata-server/internal/repositories"
	"github.com/luchacomics/ccdata-server/internal/utils"
)

var (
	auFirstName string
	auLastName string
	auEmail string
	auPassword string
	auState int
)

func init() {
	addUserCmd.Flags().StringVarP(&auFirstName, "fname", "f", "", "First name of the user account")
	addUserCmd.MarkFlagRequired("fname")
	addUserCmd.Flags().StringVarP(&auLastName, "lname", "l", "", "Last name of the user account")
	addUserCmd.MarkFlagRequired("lname")
	addUserCmd.Flags().StringVarP(&auEmail, "email", "e", "", "Email of the user account")
	addUserCmd.MarkFlagRequired("email")
	addUserCmd.Flags().StringVarP(&auPassword, "password", "p", "", "Password of the user account")
	addUserCmd.MarkFlagRequired("password")
	addUserCmd.Flags().IntVarP(&auState, "state", "s", 0, "State of the user account")
	addUserCmd.MarkFlagRequired("state")
	rootCmd.AddCommand(addUserCmd)
}

var addUserCmd = &cobra.Command{
	Use:   "add_user",
	Short: "Add user account",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		runAddUser()
	},
}

func runAddUser() {
	ctx := context.Background()

	// Load up our database.
	db, err := utils.ConnectDB(databaseHost, databasePort, databaseUser, databasePassword, databaseName)
	if err != nil {
	    log.Fatal(err)
	}
	defer db.Close()

	// Load up our repositories.
	r := repo.NewUserRepo(db)

	// Check to see the user account already exists.
	userFound, _ := r.GetByEmail(ctx, auEmail)
	if userFound != nil {
		log.Fatal("Email already exists.")
	}

	passwordHash, err := utils.HashPassword(auPassword)
	if err != nil {
		log.Fatal(err)
	}

	m := &models.User{
		Uuid: uuid.NewString(),
		FirstName: auFirstName,
		LastName: auLastName,
		Email: auEmail,
		PasswordHash: passwordHash,
		State: int8(auState),
		Timezone: "utc",
		CreatedTime: time.Now(),
		SessionUuid: uuid.NewString(),
	}

	err = r.InsertOrUpdate(ctx, m)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("\033[H\033[2J")
	fmt.Println("User created.")
}
