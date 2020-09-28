package main

import (
	"fmt"

	lift "github.com/liftbridge-io/go-liftbridge"
	"github.com/urfave/cli/v2"
)

func (h *Hub) Streams() *cli.Command {
	return &cli.Command{
		Name: "stream",
		Subcommands: []*cli.Command{
			{
				Name:      "create",
				Aliases:   []string{"c"},
				Action:    h.MustHaveConfig(h.createStream),
				ArgsUsage: "[STREAMNAME]",
				Usage:     "Create a stream",
				Category:  "stream",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "stream",
						Aliases:  []string{"s"},
						Required: true,
						Usage:    "Name of the stream",
					},
					&cli.StringFlag{
						Name:     "subject",
						Aliases:  []string{"ns"},
						Required: true,
						Usage:    "Nats Subject",
					},
				},
			},
		},
		Aliases: []string{"s"},
		Usage:   "Commands about creating/deleting/consuming from streams",
	}
}

func (h *Hub) createStream(ctx *cli.Context) error {
	streamName := ctx.String("stream")
	c, err := lift.Connect(h.cfg.Server.Addresses)
	if err != nil {
		return err
	}
	err = c.CreateStream(ctx.Context, ctx.String("subject"), ctx.String("stream"))
	if err != nil {
		if err == lift.ErrStreamExists {
			return fmt.Errorf("error while creating a stream: stream %v already exists", streamName)
		}
		return err
	}
	h.logger.Infof("Stream \"%s\" with \"%s\" NATS subject created", streamName, ctx.String("subject"))
	return nil
}
