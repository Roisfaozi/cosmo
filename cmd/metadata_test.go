package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/barasher/go-exiftool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdateMetadataCommandScenarios(t *testing.T) {
	dir := filepath.Join("..", "testdata")

	require.DirExists(t, dir, "Testdata directory should exist")
	t.Run("CompleteMetadataInCSV", func(t *testing.T) {
		imagePath := filepath.Join(dir, "image_010.jpg")

		csvPath := filepath.Join(dir, "metadata.csv")
		file, err := os.Create(csvPath)
		require.NoError(t, err)
		defer file.Close()

		file.WriteString("SourceFile,ObjectName,Keywords,CopyrightStatus,Marked,CopyrightNotice\n")
		file.WriteString(fmt.Sprintf("%s,Sample Image,Keyword1;Keyword2,protected,TRUE,All Rights Reserved\n", imagePath))

		err = UpdateMetadataFromCSV(csvPath)
		require.NoError(t, err)

		e, err := exiftool.NewExiftool()
		require.NoError(t, err)
		defer e.Close()

		meta := e.ExtractMetadata(imagePath)
		require.Len(t, meta, 1)

		title, _ := meta[0].GetString("Title")
		assert.Equal(t, "Sample Image", title, "Title should match the CSV entry")
	})

	t.Run("PartialMetadataInCSV", func(t *testing.T) {
		imagePath := filepath.Join(dir, "image_009.jpg")

		csvPath := filepath.Join(dir, "metadata.csv")
		file, err := os.Create(csvPath)
		require.NoError(t, err)
		defer file.Close()

		file.WriteString("SourceFile,ObjectName,Keywords\n")
		file.WriteString(fmt.Sprintf("%s,Partial Image,Keyword1;Keyword2\n", imagePath))

		err = UpdateMetadataFromCSV(csvPath)
		require.NoError(t, err)

		e, err := exiftool.NewExiftool()
		require.NoError(t, err)
		defer e.Close()

		meta := e.ExtractMetadata(imagePath)
		require.Len(t, meta, 1)

		title, _ := meta[0].GetString("Title")
		assert.Equal(t, "Partial Image", title, "Title should match the CSV entry")
		_, err = meta[0].GetString("CopyrightStatus")
		assert.Error(t, err, "CopyrightStatus should be empty as it was not provided in CSV")
	})

}
