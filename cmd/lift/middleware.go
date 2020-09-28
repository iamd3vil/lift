package main

import (
	"sync"

	"github.com/urfave/cli/v2"
)

var once sync.Once

func (h *Hub) MustHaveConfig(f cli.ActionFunc) cli.ActionFunc {
	once.Do(func() {
		cfg, err := initConfig(h.logger)
		if err != nil {
			h.logger.Fatalf("error while reading config: %v", err)
		}
		h.cfg = cfg
	})
	return f
}
