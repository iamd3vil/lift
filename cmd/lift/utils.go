package main

import (
	"os"
	"path"
	"runtime"
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
