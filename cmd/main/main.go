package main

import (
	"context"

	"distributedCalculator/internal/config"
	s "distributedCalculator/internal/server/http"
	"distributedCalculator/pkg/logger"

	"github.com/labstack/gommon/log"
)

const (
	serviceName = "distributed-calc"
)

func main() {
	ctx := context.Background()
	mainLogger := logger.New(serviceName)
	ctx = context.WithValue(ctx, logger.LoggerKey, mainLogger)
	cfg := config.New("local.env")
	_ = cfg

	e := s.New()

	e.EchoServer.Logger.SetLevel(log.INFO)

	e.Start()
}
