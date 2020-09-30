package main

import (
	"path"

	"github.com/urfave/cli/v2"
)

func (h *Hub) MustHaveConfig(fn cli.ActionFunc) cli.ActionFunc {
	return func(c *cli.Context) error {
		var err error

		h.logger.Debugf("config path: %v", configPath)

		if configPath == "" {
			configPath = path.Join(getHomeDir(), ".config", "lift", "config.toml")
		}

		// Initialize config.
		h.cfg, err = initConfig(h.logger, configPath)
		if err != nil {
			h.logger.Fatalf("error while reading config: %v", err)
		}

		return fn(c)
	}
}
