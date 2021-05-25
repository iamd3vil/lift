package main

import "github.com/urfave/cli/v2"

func (h *Hub) pauseStreamCommand() *cli.Command {
	return &cli.Command{
		Name:     "pause",
		Aliases:  []string{"ps"},
		Action:   h.MustHaveConfig(h.pauseStream),
		Usage:    "Pauses a stream",
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

func (h *Hub) pauseStream(ctx *cli.Context) error {
	stream := ctx.String("stream")
	c, err := getClient(h.cfg)
	if err != nil {
		return err
	}

	if err := c.PauseStream(ctx.Context, stream); err != nil {
		return err
	}

	h.logger.Infof("Stream %s paused", stream)
	return nil
}
