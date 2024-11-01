package pkg

import (
	"encoding/csv"
	"fmt"
	"github.com/barasher/go-exiftool"
	"log"
	"os"
)

func OldMetadata() {
	// Ganti dengan path file CSV Anda
	csvPath := "E:\\NGetik\\python\\bot-metadata\\new_image\\renamed_files.csv"

	// Buka file CSV
	file, err := os.Open(csvPath)
	if err != nil {
		log.Fatalf("Error opening CSV: %v", err)
	}
	defer file.Close()

	// Baca CSV
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading CSV: %v", err)
	}

	// Inisialisasi ExifTool
	e, err := exiftool.NewExiftool()
	if err != nil {
		log.Fatalf("Error initializing ExifTool: %v", err)
	}
	defer e.Close()

	// Proses setiap baris CSV (mulai dari baris ke-2)
	for _, record := range records[1:] {
		sourceFile := record[0]
		objectName := record[1]
		//captionAbstract := record[2]
		keywords := record[3]
		copyrightStatus := record[4]
		marked := record[5]
		copyrightNotice := record[6]

		// Salin file untuk pengujian
		if err != nil {
			log.Fatalf("Error creating temp directory: %v", err)
		}

		if err != nil {
			log.Fatalf("Error copying file: %v", err)
		}

		// Ekstrak metadata asli
		originals := e.ExtractMetadata(sourceFile)

		// Tampilkan Title asli
		title, _ := originals[0].GetString("Title")
		fmt.Println("Original Title: " + title)
		//
		//// Atur metadata baru
		originals[0].SetString("Title", objectName)
		originals[0].SetString("ObjectName", objectName)
		originals[0].SetString("Caption-Abstract", objectName)
		originals[0].SetString("XPSubject", objectName)
		originals[0].SetString("CopyrightStatus", copyrightStatus)
		originals[0].SetString("Marked", marked)
		originals[0].SetString("CopyrightNotice", copyrightNotice)
		originals[0].SetString("LastKeywordIPTC", keywords)
		originals[0].SetString("LastKeywordXMP", keywords)
		originals[0].SetString("Subject", keywords)
		originals[0].SetString("Keywords", "")

		// Tulis metadata baru ke file

		e.WriteMetadata(originals)

		// Ekstrak metadata setelah pembaruan
		altered := e.ExtractMetadata(sourceFile)

		title, err = altered[0].GetString("Title")
		if err != nil {
			log.Fatalf("Error getting title: %v", err)
		}
		key, _ := originals[0].GetString("Keywords")
		fmt.Println("Keywords Update: ", key)
		fmt.Println("Updated Title: ", title)
	}
}
