package main

import (
	"os"
	"path"
	"runtime"

	lift "github.com/liftbridge-io/go-liftbridge/v2"
)

func getHomeDir() string {
	var (
		home = os.Getenv("HOME")
	)
	if runtime.GOOS == "linux" {
		home = os.Getenv("XDG_CONFIG_HOME")
		if home == "" {
			home = os.Getenv("HOME")
		}
	}

	return path.Clean(home)
}

func getClient(cfg Config) (lift.Client, error) {
	return lift.Connect(cfg.Server.Addresses)
}
