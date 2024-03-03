package main

import (
	"context"
	"flag"

	_ "github.com/lib/pq"
	"github.com/nevano11/storage-of-goods/internal/storage-api/bootstrap"
	log "github.com/sirupsen/logrus"
)

func main() {
	var (
		configPath string
		logLevel   string
	)

	flag.StringVar(&configPath, "config", "cmd/storage-api/config.yaml", "path to config")
	flag.StringVar(&logLevel, "log-level", "debug", "level of logging")
	flag.Parse()

	level, err := log.ParseLevel(logLevel)
	if err != nil {
		log.Error("Failed to parse log-level")
		return
	}
	log.SetLevel(level)

	ctx := context.Background()

	err = bootstrap.Run(ctx, bootstrap.Config{
		ConfigPath: configPath,
		LogLevel:   logLevel,
	})

	if err != nil {
		log.Error(err)
	}

	log.Info("storage-api stopped")
}
