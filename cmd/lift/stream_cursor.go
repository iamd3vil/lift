package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func (h *Hub) setCursorCommand() *cli.Command {
	return &cli.Command{
		Name:     "set-cursor",
		Aliases:  []string{"sc"},
		Action:   h.MustHaveConfig(h.setCursor),
		Usage:    "Sets a cursor position for a stream",
		Category: "stream",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "id",
				Aliases:  []string{"i"},
				Required: true,
				Usage:    "ID of the cursor, generally consumer group id",
			},
			&cli.StringFlag{
				Name:     "stream",
				Aliases:  []string{"s"},
				Required: true,
				Usage:    "Name of the stream",
			},
			&cli.IntFlag{
				Name:    "partition",
				Aliases: []string{"p"},
				Value:   0,
				Usage:   "Partition",
			},
			&cli.Int64Flag{
				Name:     "offset",
				Aliases:  []string{"o"},
				Required: true,
				Usage:    "Offset to set the cursor",
			},
		},
	}
}

func (h *Hub) setCursor(ctx *cli.Context) error {
	stream := ctx.String("stream")
	partition := ctx.Int("partition")
	offset := ctx.Int64("offset")
	id := ctx.String("id")
	c, err := getClient(h.cfg)
	if err != nil {
		return err
	}
	defer c.Close()

	if err := c.SetCursor(ctx.Context, id, stream, int32(partition), offset); err != nil {
		return fmt.Errorf("error setting cursor: %v", err)
	}
	return nil
}
