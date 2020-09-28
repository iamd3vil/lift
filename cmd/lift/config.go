package main

import (
	"path"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/file"
	"github.com/sirupsen/logrus"
)

type cfgServers struct {
	Addresses []string `koanf:"addresses"`
}

type Config struct {
	Server cfgServers
}

func initConfig(logger *logrus.Logger) (Config, error) {
	var (
		cfg = Config{}
		ko  = koanf.New(".")
	)

	cfgPath := path.Join(getHomeDir(), ".config", "lift", "config.toml")

	logger.Printf("reading config: %s", cfgPath)

	if err := ko.Load(file.Provider(cfgPath), toml.Parser()); err != nil {
		logger.Fatalf("error reading config: %v", err)
	}

	// Read the configuration and load it to internal struct.
	err := ko.Unmarshal("servers", &cfg.Server)

	return cfg, err
}
