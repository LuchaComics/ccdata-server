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
	seriesFilePath string
)

func init() {
	importSeriesCmd.Flags().StringVarP(&seriesFilePath, "filepath", "f", "", "Path to the series csv file.")
	importSeriesCmd.MarkFlagRequired("filepath")
	rootCmd.AddCommand(importSeriesCmd)
}

var importSeriesCmd = &cobra.Command{
	Use:   "import_series",
	Short: "Populates the database series table.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("\033[H\033[2J") // Clear screen
		doRunImportSeries()
	},
}

func doRunImportSeries() {
	// Load up our database.
	db, err := sqldb.ConnectDB(databaseHost, databasePort, databaseUser, databasePassword, databaseName)
	if err != nil {
	    log.Fatal(err)
	}
	defer db.Close()

	// Load up our repositories.
	r := repo.NewSeriesRepo(db)

	f, err := os.Open(seriesFilePath)
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

		saveSeriesRowInDb(r, line)
	}
}

func saveSeriesRowInDb(r *repo.SeriesRepo, col []string) {
	// Extract the row.
	idString := col[0]
	name := col[1]
	sortName := col[2]
	format := col[3]
	yearBeganString := col[4]
	yearBeganUncertainString := col[5]
	yearEndedString := col[6]
	yearEndedUncertainString := col[7]
    publicationDates := col[8]
    firstIssueIdString := col[9]
    lastIssueIdString := col[10]
    isCurrentString := col[11]
    publisherIdString := col[12]
    countryIdString := col[13]
    languageIdString := col[14]
	trackingNotes := col[15]
	notes := col[16]
	hasGalleryString := col[17]
	issueCountString := col[18]
	// created // 19
    // modified // 20
	deletedString := col[21]
	hasIndiciaFrequencyString := col[22]
	hasIsbnString := col[23]
	hasBarcodeString := col[24]
	hasIssueTitleString := col[25]
	hasVolumeString := col[26]

	isComicsPublicationString := col[27]
	color := col[28]
	dimensions := col[29]
	paperStock := col[30]
	binding := col[31]
	publishingFormat := col[32]
	hasRatingString := col[33]
	publicationTypeIdString := col[34]
	isSingletonString := col[35]
	hasAboutComicsString := col[36]
	hasIndiciaPrinterString := col[37]

	// Convert the following.
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		if idString != "id" {
			log.Fatal("ParseUint err | ", err, "at", col)
		}
	}

	yearBegan, _ := strconv.ParseInt(yearBeganString, 10, 64)
	yearBeganUncertain, _ := strconv.ParseBool(yearBeganUncertainString)
    yearEnded, _ := strconv.ParseInt(yearEndedString, 10, 64)
    yearEndedUncertain, _ := strconv.ParseBool(yearEndedUncertainString)
	firstIssueId, _ := strconv.ParseUint(firstIssueIdString, 10, 64)
	lastIssueId, _ := strconv.ParseUint(lastIssueIdString, 10, 64)
	isCurrent, _ := strconv.ParseBool(isCurrentString)
	publisherId, _ := strconv.ParseUint(publisherIdString, 10, 64)
	countryId, _ := strconv.ParseUint(countryIdString, 10, 64)
	languageId, _ := strconv.ParseUint(languageIdString, 10, 64)
	hasGallery, _ := strconv.ParseBool(hasGalleryString)
	issueCount, _ := strconv.ParseInt(issueCountString, 10, 64)
	deleted, _ := strconv.ParseBool(deletedString)
	hasIndiciaFrequency, _ := strconv.ParseBool(hasIndiciaFrequencyString)
	hasIsbn, _ := strconv.ParseBool(hasIsbnString)
	hasBarcode, _ := strconv.ParseBool(hasBarcodeString)
	hasIssueTitle, _ := strconv.ParseBool(hasIssueTitleString)
	hasVolume, _ := strconv.ParseBool(hasVolumeString)

	isComicsPublication, _ := strconv.ParseBool(isComicsPublicationString)
	hasRating, _ := strconv.ParseBool(hasRatingString)
	publicationTypeId, _ := strconv.ParseUint(publicationTypeIdString, 10, 64)
	isSingleton, _ := strconv.ParseBool(isSingletonString)
	hasAboutComics, _ := strconv.ParseBool(hasAboutComicsString)
	hasIndiciaPrinter, _ := strconv.ParseBool(hasIndiciaPrinterString)

	if id != 0 {
		m := &models.Series{
			Id: id,
			Name: name,
			SortName: sortName,
			Format: format,
			YearBegan: yearBegan,
			YearBeganUncertain: yearBeganUncertain,
			YearEnded: yearEnded,
			YearEndedUncertain: yearEndedUncertain,
			PublicationDates: publicationDates,
			FirstIssueId: firstIssueId,
			LastIssueId: lastIssueId,
			IsCurrent: isCurrent,
			PublisherId: publisherId,
			CountryId: countryId,
			LanguageId: languageId,
			TrackingNotes: trackingNotes,
			Notes: notes,
			HasGallery: hasGallery,
			IssueCount: issueCount,
			Deleted: deleted,
			HasIndiciaFrequency: hasIndiciaFrequency,
			HasIsbn: hasIsbn,
			HasBarcode: hasBarcode,
			HasIssueTitle: hasIssueTitle,
			HasVolume: hasVolume,
			IsComicsPublication: isComicsPublication,
			Color: color,
			Dimensions: dimensions,
			PaperStock: paperStock,
			Binding: binding,
			PublishingFormat: publishingFormat,
			HasRating: hasRating,
			PublicationTypeId: publicationTypeId,
			IsSingleton: isSingleton,
			HasAboutComics: hasAboutComics,
			HasIndiciaPrinter: hasIndiciaPrinter,
		}

		// log.Println("Id:",m.Id)
		// log.Println("Name:",m.Name)
		// log.Println("SortName:", m.SortName)
		// log.Println("PublisherId:", m.PublisherId)
		// log.Println()

		ctx := context.Background()
		err := r.InsertOrUpdate(ctx, m)
		if err != nil {
			log.Println("Import Skipped | InsertOrUpdate | err:", err)
		} else {
			fmt.Println("Imported ID#", id)
		}
	}
}
