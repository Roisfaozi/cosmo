package cmd_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Roisfaozi/cosmo/cmd"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRenameCommandScenarios(t *testing.T) {
	dir := filepath.Join("..", "testdata")

	require.DirExists(t, dir, "Testdata directory should exist")

	t.Run("NoFilesInDirectory", func(t *testing.T) {
		csvPath := filepath.Join(dir, "renamed_files.csv")

		err := cmd.RenameImages(dir, ".jpeg")
		require.NoError(t, err)

		_, err = os.Stat(csvPath)
		assert.NoError(t, err, "CSV file should be generated even if no files are renamed")
	})

}
