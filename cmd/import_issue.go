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
    pageCount := col[14]
    pageCountUncertainString := col[15]
    indiciaFrequency := col[16]
    noIndiciaFrequencyString := col[17]
    editing := col[18]
    noEditingString := col[19]
    notes := col[20]
// 21 created
// 22 modified
    deletedString := col[23]
    isIndexedString := col[24]
    isbn := col[25]
    validIsbnString := col[26]
    noIsbnString := col[27]
    variantOfIdString := col[28]
    variantName := col[29]
    barcode := col[30]
    noBarcodeString := col[31]
    title := col[32]
    noTitleString := col[33]
    onSaleDate := col[34]
    onSaleDateUncertainString := col[35]
    rating := col[36]
    noRatingString := col[37]
    volumeNotPrintedString := col[38]
    noIndiciaPrinterString := col[39]

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
	pageCountUncertain, _ := strconv.ParseBool(pageCountUncertainString)
	noEditing, _ := strconv.ParseBool(noEditingString)
	noIndiciaFrequency, _ := strconv.ParseBool(noIndiciaFrequencyString)
	deleted, _ := strconv.ParseBool(deletedString)
	isIndexed, _ := strconv.ParseBool(isIndexedString)
	validIsbn, _ := strconv.ParseBool(validIsbnString)
	noIsbn, _ := strconv.ParseBool(noIsbnString)
	variantOfId, _ := strconv.ParseUint(variantOfIdString, 10, 64)
	noBarcode, _ := strconv.ParseBool(noBarcodeString)
	noTitle, _ := strconv.ParseBool(noTitleString)
	onSaleDateUncertain, _ := strconv.ParseBool(onSaleDateUncertainString)
    noRating, _ := strconv.ParseBool(noRatingString)
    volumeNotPrinted, _ := strconv.ParseBool(volumeNotPrintedString)
    noIndiciaPrinter, _ := strconv.ParseBool(noIndiciaPrinterString)

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
			PageCount: pageCount,
			PageCountUncertain: pageCountUncertain,
			IndiciaFrequency: indiciaFrequency,
			NoIndiciaFrequency: noIndiciaFrequency,
			Editing: editing,
			NoEditing: noEditing,
			Notes: notes,
			Deleted: deleted,
			IsIndexed: isIndexed,
			Isbn: isbn,
			ValidIsbn: validIsbn,
			NoIsbn: noIsbn,
			VariantOfId: variantOfId,
			VariantName: variantName,
			Barcode: barcode,
			NoBarcode: noBarcode,
			Title: title,
			NoTitle: noTitle,
			OnSaleDate: onSaleDate,
			OnSaleDateUncertain: onSaleDateUncertain,
			Rating: rating,
			NoRating: noRating,
			VolumeNotPrinted: volumeNotPrinted,
			NoIndiciaPrinter: noIndiciaPrinter,
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
