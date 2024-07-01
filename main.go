package main

import (
	"log"

	"github.com/ddr4869/go_template/config"
	"github.com/ddr4869/go_template/internal"
)

func main() {
	cfg := config.Init()
	router, err := internal.NewRestController(cfg)
	if err != nil {
		log.Fatalf("failed creating server: %v", err)
	}
	router.Start()
}
