package migration

import (
	"errors"
	"fmt"
	"log"
	"template/database"

	"github.com/pressly/goose/v3"
	"github.com/spf13/cobra"
)

const seedPath = "./database/seeder"

func createSeeder(name string) {
	db := database.Connection()
	if err := goose.Create(db, seedPath, name, "sql"); err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close()
}

func seed() {
	db := database.Connection()
	goose.SetTableName("seeder_version")
	if err := goose.Up(db, seedPath); err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close()

	log.Println("database is seeded successfully")
}

var SeederCmd = &cobra.Command{
	Use:   "seeder [COMMANDS]",
	Short: "Seeder tool",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("you must specify a comand like make or seed")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var CreateSeederCmd = &cobra.Command{
	Use:                   "make [ARG]",
	Short:                 "Generate seeder file",
	Long:                  "Generate seeder file",
	DisableFlagsInUseLine: true,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one arg")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		createSeeder(args[0])
	},
}

var SeedDBCmd = &cobra.Command{
	Use:                   "seed",
	Short:                 "Seeding DB",
	Long:                  "Seeding DB",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		seed()
	},
}

func init() {
	SeederCmd.AddCommand(SeedDBCmd)
	SeederCmd.AddCommand(CreateSeederCmd)
}
