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
)

// initLogger initializes logger
func initLogger(verbose bool) *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	// Set logger level
	if verbose {
		logger.SetLevel(logrus.DebugLevel)
		logger.Debug("verbose logging enabled")
	} else {
		logger.SetLevel(logrus.InfoLevel)
	}

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
			Name:  "verbose",
			Usage: "Enable verbose logging",
		},
	}

	logger := initLogger(true)

	hub, err := NewHub(logger)
	if err != nil {
		logger.Fatalf("error initializing hub: %v", err)
	}

	// Register commands.
	app.Commands = []*cli.Command{
		hub.Streams(),
	}

	// Run the app.

	if err = app.Run(os.Args); err != nil {
		logger.Errorf("OOPS: %s", err)
	}
}
