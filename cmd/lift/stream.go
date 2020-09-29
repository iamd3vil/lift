package main

import (
	"github.com/urfave/cli/v2"
)

func (h *Hub) Streams() *cli.Command {
	return &cli.Command{
		Name: "stream",
		Subcommands: []*cli.Command{
			h.createStreamCommand(),
			h.deleteStreamCommand(),
			h.subscribeCommand(),
			h.publishStreamCommand(),
		},
		Aliases: []string{"s"},
		Usage:   "Commands about creating/deleting/consuming from streams",
	}
}
