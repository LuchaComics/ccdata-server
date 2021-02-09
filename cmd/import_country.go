package cmd

import (
	"bufio"
	"context"
	"fmt"
	"encoding/csv"
	"os"
	"log"
	"io"
	"strconv"

	"github.com/spf13/cobra"

    "github.com/luchacomics/ccdata-server/internal/models"
	repo "github.com/luchacomics/ccdata-server/internal/repositories"
	sqldb "github.com/luchacomics/ccdata-server/internal/utils"
)

var (
	countryFilePath string
)

func init() {
	importCountryCmd.Flags().StringVarP(&countryFilePath, "filepath", "f", "", "Path to the country csv file.")
	importCountryCmd.MarkFlagRequired("filepath")
	rootCmd.AddCommand(importCountryCmd)
}

var importCountryCmd = &cobra.Command{
	Use:   "import_country",
	Short: "Populates the database country table.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("\033[H\033[2J") // Clear screen
		doRunImportCountry()
	},
}

func doRunImportCountry() {
	// Load up our database.
	db, err := sqldb.ConnectDB(databaseHost, databasePort, databaseUser, databasePassword, databaseName)
	if err != nil {
	    log.Fatal(err)
	}
	defer db.Close()

	// Load up our repositories.
	cr := repo.NewCountryRepo(db)

	f, err := os.Open(countryFilePath)
	if err != nil {
		log.Fatal(err)
	}

	// defer the closing of our `f` so that we can parse it later on
	defer f.Close()

	reader := csv.NewReader(bufio.NewReader(f))

	for {
		// Read line by line until no more lines left.
		line, error := reader.Read()
		if error == io.EOF {
			break
        } else if error != nil {
			log.Fatal(error)
		}

		saveCountryRowInDb(cr, line)
	}
}

func saveCountryRowInDb(cr *repo.CountryRepo, col []string) {
	// Extract the row.
	idString := col[0]
	code := col[1]
	name := col[2]

	id, _ := strconv.ParseUint(idString, 10, 64)
	if id != 0 {
		m := &models.Country{
			Id: id,
			Code: code,
			Name: name,
		}
		ctx := context.Background()
		err := cr.InsertOrUpdate(ctx, m)
		if err != nil {
			log.Panic(err)
		}
		fmt.Println("Imported ID#", id)
	}
}
