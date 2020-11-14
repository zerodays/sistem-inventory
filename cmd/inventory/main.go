package main

import (
	"github.com/urfave/cli/v2"
	"github.com/zerodays/sistem-auth/token"
	"github.com/zerodays/sistem-inventory/internal/cmd"
	"github.com/zerodays/sistem-inventory/internal/config"
	"github.com/zerodays/sistem-inventory/internal/database"
	"github.com/zerodays/sistem-inventory/internal/logger"
	"os"
)

func main() {
	// Load configuration.
	config.Load()

	// Initialize logger instance.
	logger.Init()

	// Initialize database.
	database.Init()
	defer database.Close()

	// Loads RSA public signing key for user authentication
	err := token.LoadKey(config.Microservices.UsersUrl() + "/signing_key")

	if err != nil {
		logger.Log.Fatal().Err(err).Send()
		return
	}

	// Create new CLI app.
	app := cli.NewApp()

	// Basic info.
	app.Name = "Sistem inventory microservice"
	app.Authors = []*cli.Author{
		{
			Name:  "Matej Marinko",
			Email: "matejmarinko123@gmail.com",
		},
	}
	app.Version = "0.0.1"

	// Commands for CLI app.
	app.Commands = []*cli.Command{
		cmd.Serve,
	}

	// Run the app.
	err = app.Run(os.Args)
	if err != nil {
		logger.Log.Fatal().Err(err).Send()
	}
}
