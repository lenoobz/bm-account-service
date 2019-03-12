// Package main entry point of bm account service
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/letrong/bm-account-service/conf"
	"github.com/letrong/bm-account-service/pkg/db"
	dao "github.com/letrong/bm-account-service/pkg/db/mongo"
	"github.com/letrong/bm-account-service/pkg/user"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/spf13/cobra"
)

// main entry point of the whole app
func main() {
	if err := rootCommand().Execute(); err != nil {
		log.Fatal(err)
	}
}

// rootCommand set up root command for the whole app
func rootCommand() *cobra.Command {
	rootCmd := cobra.Command{
		Use: "app",
		Run: run,
	}

	rootCmd.PersistentFlags().StringP("config", "c", "", "An explicit config file to use")
	rootCmd.Flags().IntP("port", "p", 0, "Web server port")

	return &rootCmd
}

func run(cmd *cobra.Command, args []string) {
	// Load config from ENV, config file, or command line
	config, err := conf.LoadConfig(cmd)
	if err != nil {
		log.Fatalf("Failed to load config: %s", err.Error())
		panic(err)
	}

	// Config logger
	logger, err := conf.ConfigureLogging(&config.Logger)
	if err != nil {
		log.Fatalf("Failed to configure logging: %s", err.Error())
		panic(err)
	}

	// Print out to stdout
	logger.Infoln("STARTING...")
	fmtConfig, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		logger.Errorf("Failed to parse config obj: %s", err.Error())
	}
	logger.Infof("%s\n", string(fmtConfig))

	// Init database connection
	var mongodb *mongo.Database
	switch config.DbType {
	case "mongo":
		logger.Infoln("CONNECTING DB...")
		client, err := db.InitMongoClient(&config.Mongo)
		if err != nil {
			log.Fatalf("Failed to connect to mongodb: %s", err.Error())
			panic(err)
		}

		err = client.Connect(context.Background())
		if err != nil {
			log.Fatalf("Failed to connect to mongodb: %s", err.Error())
			panic(err)
		}

		// Collection types can be used to access the database
		mongodb = client.Database(config.Mongo.Dbname)
		defer client.Disconnect(context.Background())
	default:
		panic("Unknown database")
	}

	// Init User Repo
	userRepo := dao.InitUserDao(mongodb)

	// Init User Service
	userService := user.InitUserService(userRepo)

	// Init User Control
	userControl := user.InitUserControl(userService)

	// Init User Router
	engine := user.InitUserRouter(userControl)
	port := fmt.Sprintf(":%d", config.Port)
	engine.Run(port)
}
