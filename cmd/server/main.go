package main

import (
	"perviymoiserver/pkg/adapter/rest"
	"perviymoiserver/pkg/adapter/storage/mongo"
	"perviymoiserver/pkg/shortener"
)

func main() {
	storage := mongo.New()
	s := shortener.New(storage)
	server := rest.New(s)
	err := server.Start()
	if err != nil {
		panic(err)
	}
}
