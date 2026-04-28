package main

import (
	"log"

	"learningTemp/internal/config"
	"learningTemp/internal/temporal"
)

func main() {
	cfg := config.Load()
	c, err := temporal.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	temporal.StartWorker(c)
}
