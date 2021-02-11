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
	publisherFilePath string
)

func init() {
	importPublisherCmd.Flags().StringVarP(&publisherFilePath, "filepath", "f", "", "Path to the publisher csv file.")
	importPublisherCmd.MarkFlagRequired("filepath")
	rootCmd.AddCommand(importPublisherCmd)
}

var importPublisherCmd = &cobra.Command{
	Use:   "import_publisher",
	Short: "Populates the database publisher table.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("\033[H\033[2J") // Clear screen
		doRunImportPublisher()
	},
}

func doRunImportPublisher() {
	// Load up our database.
	db, err := sqldb.ConnectDB(databaseHost, databasePort, databaseUser, databasePassword, databaseName)
	if err != nil {
	    log.Fatal(err)
	}
	defer db.Close()

	// Load up our repositories.
	r := repo.NewPublisherRepo(db)

	f, err := os.Open(publisherFilePath)
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
			continue // If error then skip this record and proceed to next record.
		}

		savePublisherRowInDb(r, line)
	}
}

func savePublisherRowInDb(r *repo.PublisherRepo, col []string) {
	// Extract the row.
	idString := col[0]
	name := col[1]

	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		if idString != "id" {
			log.Fatal("ParseUint err | ", err, "at", col)
		}
	}
	if id != 0 {
		m := &models.Publisher{
			Id: id,
			Name: name,
		}

		ctx := context.Background()
		err := r.InsertOrUpdate(ctx, m)
		if err != nil {
			log.Panic("InsertOrUpdate | err:", err)
		}
		fmt.Println("Imported ID#", id)
	}
}
