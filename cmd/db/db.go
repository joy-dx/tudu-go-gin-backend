/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package db

import (
	"github.com/spf13/cobra"
	"github.com/symball/go-gin-boilerplate/config"
	"github.com/symball/go-gin-boilerplate/migrations"
	"github.com/symball/go-gin-boilerplate/storage"
	"github.com/uptrace/bun/migrate"
)

var migrator *migrate.Migrator
var Cmd = &cobra.Command{
	Use:   "db",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the need
		cmd.Ced files
to quickly create a Cobra application.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		cfgFile, _ := cmd.Flags().GetString("config")
		config.ConfigInit(cfgFile)
		storage.DBInit()
		migrator = migrate.NewMigrator(storage.DBGet(), migrations.Migrations)
		migrator.Init(cmd.Context())
	},
}

func init() {
}
