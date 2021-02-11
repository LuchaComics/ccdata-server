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
	issueFilePath string
)

func init() {
	importIssueCmd.Flags().StringVarP(&issueFilePath, "filepath", "f", "", "Path to the issue csv file.")
	importIssueCmd.MarkFlagRequired("filepath")
	rootCmd.AddCommand(importIssueCmd)
}

var importIssueCmd = &cobra.Command{
	Use:   "import_issue",
	Short: "Populates the database issue table.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("\033[H\033[2J") // Clear screen
		doRunImportIssue()
	},
}

func doRunImportIssue() {
	// Load up our database.
	db, err := sqldb.ConnectDB(databaseHost, databasePort, databaseUser, databasePassword, databaseName)
	if err != nil {
	    log.Fatal(err)
	}
	defer db.Close()

	// Load up our repositories.
	r := repo.NewIssueRepo(db)

	f, err := os.Open(issueFilePath)
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

		saveIssueRowInDb(r, line)
	}
}

func saveIssueRowInDb(r *repo.IssueRepo, col []string) {
	// Extract the row.
	idString := col[0]
	number := col[1]

	// id	number	volume	no_volume	display_volume_with_number	series_id	indicia_publisher_id	indicia_pub_not_printed	brand_id	no_brand	publication_date	key_date	sort_code	price	page_count	page_count_uncertain	indicia_frequency	no_indicia_frequency	editing	no_editing	notes	created	modified	deleted	is_indexed	isbn	valid_isbn	no_isbn	variant_of_id	variant_name	barcode	no_barcode	title	no_title	on_sale_date	on_sale_date_uncertain	rating	no_rating	volume_not_printed	no_indicia_printer																																							


	// Convert the following.
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		if idString != "id" {
			log.Fatal("ParseUint err | ", err, "at", col)
		}
	}

	// countryId, _ := strconv.ParseUint(countryIdString, 10, 64)
	// yearBegan, _ := strconv.ParseInt(yearBeganString, 10, 64)
	// yearBeganUncertain, _ := strconv.ParseBool(yearBeganUncertainString)
    // yearEnded, _ := strconv.ParseInt(yearEndedString, 10, 64)
    // yearEndedUncertain, _ := strconv.ParseBool(yearEndedUncertainString)
    // brandCount, _ := strconv.ParseInt(brandCountString, 10, 64)
    // indiciaIssueCount, _ := strconv.ParseInt(indiciaIssueCountString, 10, 64)
    // seriesCount, _ := strconv.ParseInt(seriesCountString, 10, 64)
    // issueCount, _ := strconv.ParseInt(issueCountString, 10, 64)
    // deleted, _ := strconv.ParseBool(deletedString)
	// yearOverallBegan, _ := strconv.ParseInt(yearOverallBeganString, 10, 64)
	// yearOverallBeganUncertain, _ := strconv.ParseBool(yearOverallBeganUncertainString)
    // yearOverallEnded, _ := strconv.ParseInt(yearOverallEndedString, 10, 64)
    // yearOverallEndedUncertain, _ := strconv.ParseBool(yearOverallEndedUncertainString)

	if id != 0 {
		m := &models.Issue{
			Id: id,
			Number: number,
		}

		ctx := context.Background()
		err := r.InsertOrUpdate(ctx, m)
		if err != nil {
			log.Panic("InsertOrUpdate | err:", err)
		}
		fmt.Println("Imported ID#", id)
	}
}
