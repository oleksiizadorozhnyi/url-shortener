package main

import (
	"context"
	"github.com/caarlos0/env/v8"
	"perviymoiserver/internal/config"
	"perviymoiserver/pkg/adapter/rest"
	"perviymoiserver/pkg/adapter/storage/mongo"
	"perviymoiserver/pkg/shortener"
)

func main() {
	cfg := config.Config{}
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}
	ctx := context.Background()
	storage := mongo.New(ctx, cfg.MongoURI)
	s := shortener.New(storage)
	server := rest.New(s)
	err := server.Start(cfg.Port)
	if err != nil {
		panic(err)
	}
}
