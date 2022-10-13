package main

import (
	"fmt"
	"time"

	lift "github.com/liftbridge-io/go-liftbridge/v2"
	"github.com/urfave/cli/v2"
)

func (h *Hub) createStreamCommand() *cli.Command {
	return &cli.Command{
		Name:     "create",
		Aliases:  []string{"c"},
		Action:   h.MustHaveConfig(h.createStream),
		Usage:    "Create a stream",
		Category: "stream",
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
			&cli.IntFlag{
				Name:    "partitions",
				Aliases: []string{"p"},
				Usage:   "Number of partitions",
			},
			&cli.DurationFlag{
				Name:    "retention",
				Aliases: []string{"r"},
				Usage:   "Retention for the stream",
				Value:   7 * 24 * time.Hour,
			},
		},
	}
}

func (h *Hub) createStream(ctx *cli.Context) error {
	streamName := ctx.String("stream")
	partitions := int32(ctx.Int("partitions"))
	retention := ctx.Duration("retention")
	c, err := getClient(h.cfg)
	if err != nil {
		return err
	}
	defer c.Close()

	opts := []lift.StreamOption{
		lift.Partitions(partitions),
		lift.RetentionMaxAge(retention),
	}

	err = c.CreateStream(ctx.Context, ctx.String("subject"),
		ctx.String("stream"), opts...)
	if err != nil {
		if err == lift.ErrStreamExists {
			return fmt.Errorf("error while creating a stream: stream %v already exists", streamName)
		}
		return err
	}
	h.logger.Infof("Stream \"%s\" with \"%s\" NATS subject created", streamName, ctx.String("subject"))
	return nil
}
