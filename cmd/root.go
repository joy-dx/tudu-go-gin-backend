package cmd

import (
	"os"

	"github.com/spf13/viper"
	"github.com/symball/go-gin-boilerplate/cmd/db"
	"github.com/symball/go-gin-boilerplate/config"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	cfgFile string
	rootCmd = &cobra.Command{
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			config.ConfigInit(cfgFile)
		},
		Use: "github.com/symball/go-gin-boilerplate",
		Short: `Backend application for managing building reports.

This application uses Cobra CLI wrapper in order to better parse configuration and...
in the future support admin functions such as DB migrate, cleanup, etc.`,
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.PersistentFlags().StringP(
		"db_dsn",
		"d",
		"file::memory:?cache=shared",
		//"postgres://development:development@localhost:5432/development?sslmode=disable",
		"Using BunORM, connect with either SQLite or Postgres",
	)
	viper.BindPFlag("db_dsn", rootCmd.PersistentFlags().Lookup("db_dsn"))

	rootCmd.PersistentFlags().Uint32(
		"auth_argon_salt_length",
		32,
		"Salt length used to help create passwords",
	)
	viper.BindPFlag("auth_argon_salt_length", rootCmd.PersistentFlags().Lookup("auth_argon_salt_length"))

	rootCmd.PersistentFlags().Uint32(
		"auth_argon_memory",
		64,
		"Amount of memory to allocate when handling passwords",
	)
	viper.BindPFlag("auth_argon_memory", rootCmd.PersistentFlags().Lookup("auth_argon_memory"))

	rootCmd.PersistentFlags().Uint32(
		"auth_argon_iterations",
		3,
		"Number of passes over the memory",
	)
	viper.BindPFlag("auth_argon_iterations", rootCmd.PersistentFlags().Lookup("auth_argon_iterations"))

	rootCmd.PersistentFlags().Uint8(
		"auth_argon_parallelism",
		2,
		"The number of threads (or lanes) used by the algorithm.",
	)
	viper.BindPFlag("auth_argon_parallelism", rootCmd.PersistentFlags().Lookup("auth_argon_parallelism"))

	rootCmd.PersistentFlags().Uint32(
		"auth_argon_key_length",
		32,
		"Length of the generated key (or password hash). 16 bytes or more is recommended.",
	)
	viper.BindPFlag("auth_argon_key_length", rootCmd.PersistentFlags().Lookup("auth_argon_key_length"))

	rootCmd.AddCommand(db.Cmd)
}
