// Package cmd provide support to read setting from environment, config file, command line
package cmd

import (
	"log"

	"github.com/letrong/bm-account-service/conf"
	"github.com/letrong/bm-account-service/service"
	"github.com/letrong/bm-account-service/utils"
	"github.com/spf13/cobra"
)

// RootCommand set up and return root command
func RootCommand() *cobra.Command {
	rootCmd := cobra.Command{
		Use: "app",
		Run: run,
	}

	rootCmd.PersistentFlags().StringP("config", "c", "", "An explicit config file to use")
	rootCmd.Flags().IntP("port", "p", 0, "Web server port")

	return &rootCmd
}

func run(cmd *cobra.Command, args []string) {
	config, err := conf.LoadConfig(cmd)
	if err != nil {
		log.Fatal("Failed to load config: " + err.Error())
	}

	logger, err := conf.ConfigureLogging(&config.Logger)
	if err != nil {
		log.Fatal("Failed to configure logging: " + err.Error())
	}

	prettyConfig, err := utils.PrettyPrint(config)
	if err != nil {
		log.Fatal("Failed to parse config: " + err.Error())
	}

	logger.Infof("STARTING... \n%s", prettyConfig)
	service.StartWebServer(config, logger)
}
