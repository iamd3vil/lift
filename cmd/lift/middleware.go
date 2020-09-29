package main

import (
	"github.com/urfave/cli/v2"
)

func (h *Hub) MustHaveConfig(fn cli.ActionFunc) cli.ActionFunc {
	return func(c *cli.Context) error {
		var err error

		// Initialize config.
		h.cfg, err = initConfig(h.logger)
		if err != nil {
			h.logger.Fatalf("error while reading config: %v", err)
		}

		return fn(c)
	}
}
