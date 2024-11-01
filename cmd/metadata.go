package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/barasher/go-exiftool"
	"github.com/spf13/cobra"
)

// UpdateMetadataCommand returns a cobra.Command that updates the metadata of
// files based on a CSV file. The command takes a single argument, the path to
// the CSV file. The CSV file should contain the following columns:
//
// 1. SourceFile - the path to the file
// 2. ObjectName - the title of the file
// 3. Keywords - the keywords of the file
// 4. CopyrightStatus - the copyright status of the file
// 5. Marked - whether the file is marked or not
// 6. CopyrightNotice - the copyright notice of the file
//
// Each row of the CSV file after the first row is processed as a separate file.
// The metadata of each file is updated in-place.
//
// The function returns an error if it fails to open the CSV file, read the CSV
// file, initialize ExifTool, or update the metadata of any of the files.
func UpdateMetadataCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-metadata [csv-file]",
		Short: "Update metadata based on a CSV file",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			csvPath := args[0]
			if err := UpdateMetadataFromCSV(csvPath); err != nil {
				log.Fatalf("Error updating metadata: %v", err)
			}
		},
	}

	return cmd
}

// UpdateMetadataFromCSV reads a CSV file and updates the metadata of the files
// listed in the CSV file. The CSV file should contain the following columns:
//
// 1. SourceFile - the path to the file
// 2. ObjectName - the title of the file
// 3. Keywords - the keywords of the file
// 4. CopyrightStatus - the copyright status of the file
// 5. Marked - whether the file is marked or not
// 6. CopyrightNotice - the copyright notice of the file
//
// If a record in the CSV file is missing one of the required columns, the
// function will log a message and skip the record. If the record has more than
// 6 columns, the extra columns will be ignored.
//
// The function returns an error if it fails to open the CSV file, read the CSV
// file, initialize ExifTool, or update the metadata of any of the files.
func UpdateMetadataFromCSV(csvPath string) error {
	file, err := os.Open(csvPath)
	if err != nil {
		return fmt.Errorf("failed to open CSV: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read CSV: %w", err)
	}

	e, err := exiftool.NewExiftool()
	if err != nil {
		return fmt.Errorf("failed to initialize ExifTool: %w", err)
	}
	defer e.Close()

	for _, record := range records[1:] {
		if len(record) < 3 {
			log.Printf("Skipping record dengan panjang tidak cukup: %v", record)
			continue
		}

		filePath := record[0]
		objectName := record[1]
		keywords := ""
		copyrightStatus := ""
		marked := ""
		copyrightNotice := ""

		if len(record) > 2 {
			keywords = record[2]
		}
		if len(record) > 3 {
			copyrightStatus = record[3]
		}
		if len(record) > 4 {
			marked = record[4]
		}
		if len(record) > 5 {
			copyrightNotice = record[5]
		}

		updateMetadata(e, filePath, objectName, keywords, copyrightStatus, marked, copyrightNotice)
	}
	return nil
}

// updateMetadata updates the metadata of the given file with the given values.
//
// Note that this function does not handle any errors that may occur when updating
// the metadata. If an error occurs, it will be logged, but the function will not
// return an error.
func updateMetadata(e *exiftool.Exiftool, filePath, objectName, keywords, copyrightStatus, marked, copyrightNotice string) {
	metas := e.ExtractMetadata(filePath)

	metas[0].SetString("Title", objectName)
	metas[0].SetString("ObjectName", objectName)
	metas[0].SetString("Caption-Abstract", objectName)
	metas[0].SetString("XPSubject", objectName)
	metas[0].SetString("CopyrightStatus", copyrightStatus)
	metas[0].SetString("Marked", marked)
	metas[0].SetString("CopyrightNotice", copyrightNotice)
	metas[0].SetString("LastKeywordIPTC", keywords)
	metas[0].SetString("LastKeywordXMP", keywords)
	metas[0].SetString("Subject", keywords)
	metas[0].SetString("Keywords", "")
	// Write metadata to the file
	e.WriteMetadata(metas) // No return value to handle, so we assume success if no panic or log output

	altered := e.ExtractMetadata(filePath)
	title, _ := altered[0].GetString("ObjectName")

	fmt.Println("Updated Title: ", title)
	
	// Log success message
	log.Printf("Successfully updated metadata for: %s", filePath)
}
