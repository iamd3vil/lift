package main

import "github.com/urfave/cli/v2"

func (h *Hub) MustHaveConfig(f cli.ActionFunc) cli.ActionFunc {
	cfg, err := initConfig(h.logger)
	if err != nil {
		h.logger.Fatalf("error while reading config: %v", err)
	}
	h.cfg = cfg
	return f
}
