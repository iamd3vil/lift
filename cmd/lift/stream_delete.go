package main

import "github.com/urfave/cli/v2"

func (h *Hub) deleteStreamCommand() *cli.Command {
	return &cli.Command{
		Name:     "delete",
		Aliases:  []string{"d"},
		Action:   h.MustHaveConfig(h.deleteStream),
		Usage:    "Deletes a stream",
		Category: "stream",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "stream",
				Aliases:  []string{"s"},
				Required: true,
				Usage:    "Name of the stream",
			},
		},
	}
}

func (h *Hub) deleteStream(ctx *cli.Context) error {
	stream := ctx.String("stream")
	c, err := getClient(h.cfg)
	if err != nil {
		return err
	}
	defer c.Close()

	if err := c.DeleteStream(ctx.Context, stream); err != nil {
		return err
	}
	h.logger.Infof("Stream %s deleted", stream)
	return nil
}
