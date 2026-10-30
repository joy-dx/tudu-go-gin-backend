/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package db

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/symball/go-gin-boilerplate/migrations"
	"github.com/symball/go-gin-boilerplate/storage"
	"github.com/uptrace/bun/migrate"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		bunMigrations, getMigrationsErr := migrations.GetForDatabase()
		if getMigrationsErr != nil {
			return fmt.Errorf("could not get migrations: %w", getMigrationsErr)
		}

		migrator := migrate.NewMigrator(
			storage.DBGet(),
			bunMigrations,
		)

		if err := migrator.Init(cmd.Context()); err != nil {
			return fmt.Errorf("initialize migrator: %w", err)
		}

		ms, err := migrator.MigrationsWithStatus(cmd.Context())
		if err != nil {
			return fmt.Errorf("could not get migration status: %w", err)
		}
		fmt.Printf("migrations: %s\n", ms)
		fmt.Printf("unapplied migrations: %s\n", ms.Unapplied())
		fmt.Printf("last migration group: %s\n", ms.LastGroup())
		return nil
	},
}

func init() {
	Cmd.AddCommand(statusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
