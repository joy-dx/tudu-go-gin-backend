/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package db

import (
	"fmt"
	"github.com/symball/go-gin-boilerplate/migrations"
	"github.com/symball/go-gin-boilerplate/storage"
	"github.com/uptrace/bun/migrate"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

// newMigrationCmd represents the newMigration command
var newMigrationCmd = &cobra.Command{
	Use:   "newMigration",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1)),
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		migrator := migrate.NewMigrator(storage.DBGet(), migrations.Migrations)
		name := strings.Join(args, "_")
		files, err := migrator.CreateSQLMigrations(cmd.Context(), name)
		if err != nil {
			log.Fatal(err)
		}

		for _, mf := range files {
			fmt.Printf("created migration %s (%s)\n", mf.Name, mf.Path)
		}
	},
}

func init() {
	Cmd.AddCommand(newMigrationCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newMigrationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newMigrationCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
