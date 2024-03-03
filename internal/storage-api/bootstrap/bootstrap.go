package bootstrap

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/nevano11/storage-of-goods/internal/storage-api/config"
	"github.com/nevano11/storage-of-goods/internal/storage-api/handler"
	"github.com/nevano11/storage-of-goods/internal/storage-api/repository/postgres"
	"github.com/nevano11/storage-of-goods/internal/storage-api/repository/postgres/repository"
	"github.com/nevano11/storage-of-goods/internal/storage-api/server"
	"github.com/nevano11/storage-of-goods/internal/storage-api/service"
)

func Run(ctx context.Context, cfg Config) error {
	appConfig, err := config.Load(cfg.ConfigPath)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	err = validator.New(validator.WithRequiredStructEnabled()).Struct(appConfig)
	if err != nil {
		return fmt.Errorf("error validating config: %w", err)
	}

	db, err := postgres.NewPostgresDb(appConfig.Database)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	storage := repository.NewRepository(db)
	svc := service.NewService(storage)
	httpHandler := handler.NewHttpHandler(svc)

	routes := httpHandler.InitRoutes()

	apiserver, err := server.NewHttpServer(
		cfg.LogLevel,
		appConfig.Server.Port,
		routes)

	if err != nil {
		return fmt.Errorf("failed to start http server: %w", err)
	}

	return apiserver.Start()
}
