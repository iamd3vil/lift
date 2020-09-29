package main

import (
	"github.com/liftbridge-io/go-liftbridge"
	cli "github.com/urfave/cli/v2"
)

func (h *Hub) subscribeCommand() *cli.Command {
	return &cli.Command{
		Name:     "subscribe",
		Aliases:  []string{"sb"},
		Action:   h.MustHaveConfig(h.subscribeStream),
		Usage:    "subscribe from a stream",
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
				Usage:   "Partition to subscribe",
			},
			&cli.Int64Flag{
				Name:    "offset",
				Aliases: []string{"o"},
				Usage:   "Offset to consume from",
			},
			&cli.BoolFlag{
				Name:    "earliest",
				Aliases: []string{"e"},
				Usage:   "Start from earliest",
			},
		},
	}
}

func (h *Hub) subscribeStream(ctx *cli.Context) error {
	stream := ctx.String("stream")
	partition := int32(ctx.Int("partition"))
	earliest := ctx.Bool("earliest")
	offset := ctx.Int64("offset")

	subOpts := []liftbridge.SubscriptionOption{
		liftbridge.Partition(partition),
	}
	if earliest {
		subOpts = append(subOpts, liftbridge.StartAtEarliestReceived())
	} else if offset == 0 {
		subOpts = append(subOpts, liftbridge.StartAtLatestReceived())
	} else {
		subOpts = append(subOpts, liftbridge.StartAtOffset(offset))
	}

	c, err := getClient(h.cfg)
	if err != nil {
		return err
	}
	defer c.Close()

	err = c.Subscribe(ctx.Context, stream, func(msg *liftbridge.Message, err error) {
		if err != nil {
			h.logger.Fatalf("error while subscribing to the stream: %v", err)
		}
		h.logger.Printf("Partition: %d, Offset: %d, Message: %s", msg.Partition(), msg.Offset(), string(msg.Value()))
	}, subOpts...)
	if err != nil {
		return err
	}

	<-ctx.Context.Done()

	return nil
}
