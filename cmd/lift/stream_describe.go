package main

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli/v2"
)

func (h *Hub) describeCommand() *cli.Command {
	return &cli.Command{
		Name:     "describe",
		Aliases:  []string{"desc"},
		Action:   h.MustHaveConfig(h.describeStream),
		Usage:    "Describe a stream",
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

func (h *Hub) describeStream(ctx *cli.Context) error {
	stream := ctx.String("stream")

	c, err := getClient(h.cfg)
	if err != nil {
		return err
	}
	defer c.Close()

	metadata, err := c.FetchMetadata(ctx.Context)
	streamInfo := metadata.GetStream(stream)
	fmt.Printf("Stream: %s\n", stream)
	fmt.Printf("Subject: %s\n", streamInfo.Subject())
	fmt.Printf("Number of partitions: %d\n\n", len(streamInfo.Partitions()))

	table := tablewriter.NewWriter(os.Stdout)
	table.SetBorder(false)
	table.SetHeader([]string{"Partition", "HighWaterMark", "Newest Offset", "Replicas", "In Sync Replicas"})
	for _, info := range streamInfo.Partitions() {
		// TODO: Change this to metadata from `info` when the bug in liftbridge client is fixed.
		// No need to fetch partition metadata again
		pMeta, err := c.FetchPartitionMetadata(ctx.Context, stream, info.ID())
		if err != nil {
			return err
		}
		partitionID := fmt.Sprintf("%d", pMeta.ID())
		hw := fmt.Sprintf("%d", pMeta.HighWatermark())
		offset := fmt.Sprintf("%d", pMeta.NewestOffset())
		replicas := fmt.Sprintf("[%d]", len(pMeta.Replicas()))
		isr := fmt.Sprintf("[%d]", len(pMeta.ISR()))
		// partitionID := fmt.Sprintf("%d", info.ID())
		// hw := fmt.Sprintf("%d", info.HighWatermark())
		// offset := fmt.Sprintf("%d", info.NewestOffset())
		// replicas := fmt.Sprintf("[%d]", len(info.Replicas()))
		// isr := fmt.Sprintf("[%d]", len(info.ISR()))

		table.Append([]string{partitionID, hw, offset, replicas, isr})
	}

	table.Render()
	return nil
}
