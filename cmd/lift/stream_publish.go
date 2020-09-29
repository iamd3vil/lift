package main

import (
	"fmt"

	"github.com/liftbridge-io/go-liftbridge"
	cli "github.com/urfave/cli/v2"
)

func (h *Hub) publishStreamCommand() *cli.Command {
	return &cli.Command{
		Name:     "publish",
		Aliases:  []string{"p"},
		Action:   h.MustHaveConfig(h.publishStream),
		Usage:    "Publish to a stream",
		Category: "stream",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "stream",
				Aliases:  []string{"s"},
				Required: true,
				Usage:    "Name of the stream",
			},
			&cli.IntFlag{
				Name:    "partition",
				Aliases: []string{"p"},
				Usage:   "Partition to publish to",
			},
		},
	}
}

func (h *Hub) publishStream(ctx *cli.Context) error {
	stream := ctx.String("stream")
	partition := int32(ctx.Int("partition"))
	c, err := getClient(h.cfg)
	if err != nil {
		return err
	}
	defer c.Close()

	var (
		val string
	)

	for {
		if _, err := fmt.Scanln(&val); err != nil {
			h.logger.Errorf("error while reading input: %v", err)
			break
		}
		if _, err := c.Publish(ctx.Context, stream, []byte(val), liftbridge.ToPartition(partition)); err != nil {
			return err
		}
	}

	return nil
}
