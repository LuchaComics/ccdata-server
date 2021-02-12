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
	volume := col[2]
	noVolumeString := col[3]
	displayVolumeWithNumberString := col[4]
    seriesIdString := col[5]
	indiciaPublisherIdString := col[6]
	indiciaPubNotPrintedString := col[7]
	brandIdString := col[8]
    noBrandString := col[9]
    publicationDate := col[10]
    keyDate := col[11]
    sortCode := col[12]
    price := col[13]
// 14 page_count
// 15 page_count_uncertain
// 16 indicia_frequency
// 17 no_indicia_frequency
// 18 editing
// 19 no_editing
// 20 notes
// 21 created
// 22 modified
// 23 deleted
// 24 is_indexed
// 25 isbn
// 26 valid_isbn
// 27 no_isbn
// 28 variant_of_id
// 29 variant_name
// 30 barcode
// 31 no_barcode
// 32 title
// 33 no_title
// 34 on_sale_date
// 35 on_sale_date_uncertain
// 36 rating
// 37 no_rating
// 38 volume_not_printed
// 39 no_indicia_printer

	fmt.Println(14, col[14])
	fmt.Println(15, col[15])
	fmt.Println(16, col[16])
	fmt.Println(17, col[17])
	fmt.Println(18, col[18])
	fmt.Println(19, col[19])
	fmt.Println(20, col[20])
	fmt.Println(21, col[21])
	fmt.Println(22, col[22])
	fmt.Println(23, col[23])
	fmt.Println(24, col[24])
	fmt.Println(25, col[25])
	fmt.Println(26, col[26])
	fmt.Println(27, col[27])
	fmt.Println(28, col[28])
	fmt.Println(29, col[29])
	fmt.Println(30, col[30])
	fmt.Println(31, col[31])
	fmt.Println(32, col[32])
	fmt.Println(33, col[33])
	fmt.Println(34, col[34])
	fmt.Println(35, col[35])
	fmt.Println(36, col[36])
	fmt.Println(37, col[37])
	fmt.Println(38, col[38])
	fmt.Println(39, col[39])
    fmt.Println()

	// Convert the following.
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		if idString != "id" {
			log.Fatal("ParseUint err | ", err, "at", col)
		}
	}

    noVolume, _ := strconv.ParseBool(noVolumeString)
	displayVolumeWithNumber, _ := strconv.ParseBool(displayVolumeWithNumberString)
	seriesId, _ := strconv.ParseUint(seriesIdString, 10, 64)
	indiciaPublisherId, _ := strconv.ParseUint(indiciaPublisherIdString, 10, 64)
	indiciaPubNotPrinted, _ := strconv.ParseBool(indiciaPubNotPrintedString)
	brandId, _ := strconv.ParseUint(brandIdString, 10, 64)
	noBrand, _ := strconv.ParseBool(noBrandString)

	// yearBegan, _ := strconv.ParseInt(yearBeganString, 10, 64)
	// yearBeganUncertain, _ := strconv.ParseBool(yearBeganUncertainString)
    // volume, _ := strconv.ParseInt(volumeString, 10, 64)
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
			Volume: volume,
			NoVolume: noVolume,
			DisplayVolumeWithNumber: displayVolumeWithNumber,
			SeriesId: seriesId,
			IndiciaPublisherId: indiciaPublisherId,
			IndiciaPubNotPrinted: indiciaPubNotPrinted,
			BrandId: brandId,
			NoBrand: noBrand,
			PublicationDate: publicationDate,
			KeyDate: keyDate,
			SortCode: sortCode,
			Price: price,
		}

		ctx := context.Background()
		err := r.InsertOrUpdate(ctx, m)
		if err != nil {
			fmt.Println("Skipping ID#", id, " b/c err:", err)
			// panic(err)
		}
		fmt.Println("Imported ID#", id)
	}
}
