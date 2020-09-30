package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	cli "github.com/urfave/cli/v2"
)

var (
	// Version and date of the build. This is injected at build-time.
	buildVersion = "unknown"
	buildDate    = "unknown"
	configPath   string
	verbose      bool
)

// initLogger initializes logger
func initLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	// Set logger level
	logger.SetLevel(logrus.InfoLevel)

	return logger
}

func main() {
	// Intialize new CLI app
	app := cli.NewApp()
	app.Name = "lift"
	app.Usage = "A CLI Client for Liftbridge"
	app.Version = fmt.Sprintf("%s, %s", buildVersion, buildDate)
	app.Authors = []*cli.Author{
		{
			Name: "Sarat Chandra <me@saratchandra.in>",
		},
	}
	// Register command line args.
	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:        "verbose",
			Usage:       "Enable verbose logging",
			Destination: &verbose,
		},
		&cli.StringFlag{
			Name:        "config",
			Aliases:     []string{"c"},
			Usage:       "Configuration Path",
			Destination: &configPath,
		},
	}

	logger := initLogger()

	hub, err := NewHub(logger)
	if err != nil {
		logger.Fatalf("error initializing hub: %v", err)
	}

	app.Before = func(ctx *cli.Context) error {
		if verbose {
			hub.logger.SetLevel(logrus.DebugLevel)
			hub.logger.Debug("verbose logging enabled")
		}
		return nil
	}

	// Register commands.
	app.Commands = []*cli.Command{
		hub.Streams(),
		hub.initCommand(),
	}

	// Run the app.

	if err = app.Run(os.Args); err != nil {
		logger.Errorf("OOPS: %s", err)
	}
}
