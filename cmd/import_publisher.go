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
	countryIdString := col[2]
	yearBeganString := col[3]
	yearEndedString := col[4]
	notes := col[5]
	url := col[6]
	brandCountString := col[7]
    indiciaPublisherCountString := col[8]
    seriesCountString := col[9]
    // // created // 10
    // // modified // 11
	issueCountString := col[12]
    deletedString := col[13]
	yearBeganUncertainString := col[14]
	yearEndedUncertainString := col[15]
	yearOverallBeganString := col[16]
	yearOverallBeganUncertainString := col[17]
	yearOverallEndedString := col[18]
	yearOverallEndedUncertainString := col[19]

	// Convert the following.
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		if idString != "id" {
			log.Fatal("ParseUint err | ", err, "at", col)
		}
	}

	countryId, _ := strconv.ParseUint(countryIdString, 10, 64)
	yearBegan, _ := strconv.ParseInt(yearBeganString, 10, 64)
	yearBeganUncertain, _ := strconv.ParseBool(yearBeganUncertainString)
    yearEnded, _ := strconv.ParseInt(yearEndedString, 10, 64)
    yearEndedUncertain, _ := strconv.ParseBool(yearEndedUncertainString)
    brandCount, _ := strconv.ParseInt(brandCountString, 10, 64)
    indiciaPublisherCount, _ := strconv.ParseInt(indiciaPublisherCountString, 10, 64)
    seriesCount, _ := strconv.ParseInt(seriesCountString, 10, 64)
    issueCount, _ := strconv.ParseInt(issueCountString, 10, 64)
    deleted, _ := strconv.ParseBool(deletedString)
	yearOverallBegan, _ := strconv.ParseInt(yearOverallBeganString, 10, 64)
	yearOverallBeganUncertain, _ := strconv.ParseBool(yearOverallBeganUncertainString)
    yearOverallEnded, _ := strconv.ParseInt(yearOverallEndedString, 10, 64)
    yearOverallEndedUncertain, _ := strconv.ParseBool(yearOverallEndedUncertainString)

	if id != 0 {
		m := &models.Publisher{
			Id: id,
			Name: name,
			CountryId: countryId,
			YearBegan: yearBegan,
			YearBeganUncertain: yearBeganUncertain,
			YearEnded: yearEnded,
			YearEndedUncertain: yearEndedUncertain,
			Notes: notes,
			Url: url,
			BrandCount: brandCount,
			IndiciaPublisherCount: indiciaPublisherCount,
			SeriesCount: seriesCount,
			IssueCount: issueCount,
			Deleted: deleted,
			YearOverallBegan: yearOverallBegan,
			YearOverallBeganUncertain: yearOverallBeganUncertain,
			YearOverallEnded: yearOverallEnded,
			YearOverallEndedUncertain: yearOverallEndedUncertain,
		}

		ctx := context.Background()
		err := r.InsertOrUpdate(ctx, m)
		if err != nil {
			log.Panic("InsertOrUpdate | err:", err)
		}
		fmt.Println("Imported ID#", id)
	}
}
