module github.com/iamd3vil/liftbridge

go 1.15

require (
	github.com/knadh/koanf v0.13.0
	github.com/liftbridge-io/go-liftbridge/v2 v2.0.2-0.20201119170214-c842e2f19749
	github.com/mattn/go-runewidth v0.0.10 // indirect
	github.com/olekukonko/tablewriter v0.0.4
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/sirupsen/logrus v1.6.0
	github.com/urfave/cli/v2 v2.2.0
)

replace github.com/liftbridge-io/go-liftbridge/v2 => ../go-liftbridge/v2
