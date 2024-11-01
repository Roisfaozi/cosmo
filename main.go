package main

import (
	"github.com/Roisfaozi/cosmo/cmd"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "cosmo"}

	rootCmd.AddCommand(cmd.RenameCommand())
	rootCmd.AddCommand(cmd.UpdateMetadataCommand())

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
		os.Exit(1)
	}
}
