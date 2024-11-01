package cmd

import (
	"encoding/csv"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
)

// RenameCommand returns a cobra command that renames all files in the given
// directory to "image_<number>.<ext>" where <number> is the 1-indexed sequence
// number of the file and <ext> is the file extension specified by the --ext
// flag (defaulting to ".jpg"). The command also generates a CSV file in the
// same directory containing the original file name and the new file name.
func RenameCommand() *cobra.Command {
	var ext string

	cmd := &cobra.Command{
		Use:   "rename [directory]",
		Short: "Rename images sequentially and export to CSV",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			dir := args[0]
			if err := RenameImages(dir, ext); err != nil {
				log.Fatalf("Error renaming images: %v", err)
			}
		},
	}

	cmd.Flags().StringVar(&ext, "ext", ".jpg", "File extension to rename")
	return cmd
}

// RenameImages renames all files in the given directory to
// "image_<number>.<ext>" where <number> is the 1-indexed sequence number of the
// file and <ext> is the file extension specified by the ext parameter. The
// function also generates a CSV file in the same directory containing the
// original file name and the new file name.
//
// The CSV file has the following columns:
//
// * SourceFile: the original file name
// * ObjectName: the new file name
// * Keywords: always empty
// * CopyrightStatus: always "protected"
// * Marked: always "TRUE"
// * CopyrightNotice: always "All Rights Reserved"
func RenameImages(dir, ext string) error {
	files, err := filepath.Glob(filepath.Join(dir, "*"+ext))
	if err != nil {
		return fmt.Errorf("failed to read directory: %w", err)
	}

	csvFile, err := os.Create(filepath.Join(dir, "renamed_files.csv"))
	if err != nil {
		return fmt.Errorf("failed to create CSV: %w", err)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	writer.Write([]string{"SourceFile", "ObjectName", "Keywords", "CopyrightStatus", "Marked", "CopyrightNotice"})

	for i, file := range files {
		newName := fmt.Sprintf("image_%03d%s", i+1, ext)
		newPath := filepath.Join(dir, newName)

		if err := os.Rename(file, newPath); err != nil {
			log.Printf("Failed to rename %s: %v", file, err)
			continue
		}

		writer.Write([]string{
			newPath, newName, "", "protected", "TRUE", "All Rights Reserved",
		})
		fmt.Printf("Renamed: %s -> %s\n", file, newName)
	}

	return nil
}
