package cmd_test

import (
	"fmt"
	"github.com/Roisfaozi/cosmo/cmd"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRenameCommandScenarios(t *testing.T) {
	// 1. Tes untuk direktori tanpa file
	t.Run("NoFilesInDirectory", func(t *testing.T) {
		dir := "E:\\NGetik\\golang\\garap\\exif-editor\\testdata"
		csvPath := filepath.Join(dir, "renamed_files.csv")

		err := cmd.RenameImages(dir, ".jpg")
		require.NoError(t, err)

		fmt.Println("Path", csvPath)
		_, err = os.Stat(csvPath)
		assert.NoError(t, err, "CSV file should be generated even if no files are renamed")
	})

}
