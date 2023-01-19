package migration

import (
	"log"
	"template/database"

	"github.com/pressly/goose/v3"
	"github.com/spf13/cobra"
)

func migrate() {
	db := database.Connection()
	if err := goose.Up(db, migrationPath); err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close()
}

var MigrateCmd = &cobra.Command{
	Use:                   "migrate",
	Short:                 "Migrate all migration file",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		migrate()
	},
}
