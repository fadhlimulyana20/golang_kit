package migration

import (
	"fmt"
	"log"
	"template/database"

	"github.com/pressly/goose/v3"
	"github.com/spf13/cobra"
)

// Create Migration File
func migrate() {
	db := database.Connection()
	if err := goose.Up(db, migrationPath); err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close()
}

// Rollback migration with step
func rollback() {
	db := database.Connection()
	if err := goose.Down(db, migrationPath); err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close()
}

// Show migration version
func version() {
	db := database.Connection()
	if err := goose.Version(db, migrationPath); err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close()
}

var MigrationCmd = &cobra.Command{
	Use:   "migration [COMMANDS]",
	Short: "Migration tool",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("you must specify a comand like create, migrate, etc")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var MigrateCmd = &cobra.Command{
	Use:                   "migrate",
	Short:                 "Migrate all migration file",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		migrate()
	},
}

var RollbackCmd = &cobra.Command{
	Use:   "rollback",
	Short: "Rollback migration",
	Run: func(cmd *cobra.Command, args []string) {
		step, _ := cmd.Flags().GetInt("step")
		for i := 0; i < step; i++ {
			rollback()
		}
	},
}

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show migration version",
	Run: func(cmd *cobra.Command, args []string) {
		version()
	},
}

func init() {
	RollbackCmd.PersistentFlags().Int("step", 1, "step for rolling back migration")

	MigrationCmd.AddCommand(CreateMigrationCmd)
	MigrationCmd.AddCommand(MigrateCmd)
	MigrationCmd.AddCommand(RollbackCmd)
	MigrationCmd.AddCommand(VersionCmd)
}
