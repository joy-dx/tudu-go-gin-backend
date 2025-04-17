package db

import (
	"github.com/spf13/cobra"
	"github.com/symball/go-gin-boilerplate/storage"
	"github.com/symball/go-gin-boilerplate/todos"
	"github.com/symball/go-gin-boilerplate/users"
	"github.com/uptrace/bun/dbfixture"
	"log"
	"os"
)

// migrateCmd represents the migrate command
var fixturesCmd = &cobra.Command{
	Use:   "fixtures",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		db := storage.DBGet()

		// Todos fixtures
		db.RegisterModel((*users.User)(nil), (*todos.Todo)(nil))

		fixture := dbfixture.New(db, dbfixture.WithTruncateTables())
		err := fixture.Load(cmd.Context(), os.DirFS("fixtures"), "testdata.yml")
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	Cmd.AddCommand(fixturesCmd)
}
