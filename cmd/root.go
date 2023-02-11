package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"template/cmd/http"
	"template/cmd/migration"
	"template/cmd/stub"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "go-kit",
		Short: "Go server application",
		Long:  `CLI to build server application`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	ctx, cancel := context.WithCancel(context.Background())

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-quit
		cancel()
	}()

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	viper.SetDefault("license", "apache")

	comands := []*cobra.Command{
		{
			Use:   "serve",
			Short: "Start HTTP server",
			Long:  "Start HTTP Server",
			Run: func(cmd *cobra.Command, args []string) {
				http.StartServer(ctx)
			},
		},
		{
			Use:   "stubilize",
			Short: "Make stubs file from source code",
			Long:  "Make stubs file from source code",
			Run: func(cmd *cobra.Command, args []string) {
				stub.MakeStub()
			},
		},
		{
			Use:   "unstubilize [args]",
			Short: "Make source code from stub files",
			Long:  "Make source code from stub files",
			Args:  cobra.MinimumNArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				stub.GenerateFromStub(args[0])
			},
		},
		{
			Use:   "generate [type] [name]",
			Short: "Make source code from template stub files",
			Long:  "Make source code from template stub files",
			Args:  cobra.MinimumNArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				stub.GenerateFromTemplateStub(args[0], "default", args[1])
			},
		},
	}
	rootCmd.AddCommand(comands...)
	rootCmd.AddCommand(migration.MigrationCmd)
	rootCmd.AddCommand(stub.TemplateCmd)
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cobra")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
