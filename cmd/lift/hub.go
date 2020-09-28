package main

import (
	"github.com/sirupsen/logrus"
)

type Hub struct {
	logger *logrus.Logger
	cfg    Config
}

func NewHub(logger *logrus.Logger) (*Hub, error) {
	hub := &Hub{
		logger: logger,
	}

	return hub, nil
}
