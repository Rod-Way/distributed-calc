package main

import (
	"fmt"

	"distributedCalculator/internal/config"
	s "distributedCalculator/internal/server/http"
)

func main() {
	cfg := config.New("local.env")

	e := s.New()

	fmt.Println(cfg)

	e.Start(cfg.RestServerPort)
}
