package main

import (
	"os"
	"path"

	"github.com/urfave/cli/v2"
)

func (h *Hub) initCommand() *cli.Command {
	return &cli.Command{
		Name:    "init",
		Aliases: []string{"i"},
		Usage:   "Initialize a sample config",
		Action:  h.initConfig,
	}
}

func (h *Hub) initConfig(ctx *cli.Context) error {
	cfgPath := path.Join(getHomeDir(), ".config", "lift", "config.toml")
	_, err := os.Stat(cfgPath)
	if err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(path.Join(getHomeDir(), ".config", "lift"), 0750)
			f, err := os.Create(cfgPath)
			if err != nil {
				return err
			}
			defer f.Close()
			if _, err := f.Write([]byte(defaultConfig)); err != nil {
				return err
			}
			return nil
		}
		return err
	}
	h.logger.Infof("Config at %v already exists", cfgPath)
	return nil
}
