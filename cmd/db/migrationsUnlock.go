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

// rollbackCmd represents the rollback command
var migrationsUnlockCmd = &cobra.Command{
	Use:   "migrationsUnlock",
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

		migrator.Unlock(cmd.Context())
		return nil
	},
}

func init() {
	Cmd.AddCommand(migrationsUnlockCmd)
}
