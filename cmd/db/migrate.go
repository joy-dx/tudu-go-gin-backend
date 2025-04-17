/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package db

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := migrator.Lock(cmd.Context()); err != nil {
			log.Panic(err)
		}
		defer migrator.Unlock(cmd.Context())

		group, err := migrator.Migrate(cmd.Context())
		if err != nil {
			log.Panic(err)
		}
		if group.IsZero() {
			fmt.Printf("there are no new migrations to run (database is up to date)\n")
			return
		}
		fmt.Printf("migrated to %s\n", group)
	},
}

func init() {
	Cmd.AddCommand(migrateCmd)
}
