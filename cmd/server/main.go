package main

import (
	"context"
	"perviymoiserver/pkg/adapter/rest"
	"perviymoiserver/pkg/adapter/storage/mongo"
	"perviymoiserver/pkg/shortener"
)

func main() {
	ctx := context.Background()
	storage := mongo.New(ctx)
	s := shortener.New(storage)
	server := rest.New(s)
	err := server.Start()
	if err != nil {
		panic(err)
	}
}
