package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/SShlykov/lbalancer/lbcer/internal/bootstrap/app"
	"github.com/SShlykov/lbalancer/lbcer/internal/config"
)

const (
	StoppedOK = iota
	ErrConfigLoadCode
	StoppedNotCreated
	StoppedCrashed
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "./config", "path to configuration file")

	ctx := context.Background()
	cfg, err := config.Load(configPath)
	if err != nil {
		log.Println("error while loading configuration", err)
		os.Exit(ErrConfigLoadCode)
	}

	appl, err := app.New(ctx, cfg)
	if err != nil {
		log.Println("Stopped during creating app: ", err)
		os.Exit(StoppedNotCreated)
	}

	if err = appl.Run(); err != nil {
		log.Println("Stopped while running app: ", err)
		os.Exit(StoppedCrashed)
	}

	os.Exit(StoppedOK)
}
