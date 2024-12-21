package main

import (
	"fmt"

	"distributedCalculator/internal/config"
	s "distributedCalculator/internal/server/http"
)

const (
	serviceName = "distributed-calc"
)

func main() {
	cfg := config.New("local.env")

	e := s.New()

	e.Start(fmt.Sprint(":%s", cfg.RestServerPort))
}
