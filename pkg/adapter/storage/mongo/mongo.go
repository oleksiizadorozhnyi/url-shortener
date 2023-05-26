package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"perviymoiserver/pkg/model"
)

type Storage struct {
	client          *mongo.Client
	database        *mongo.Database
	collectionPages *mongo.Collection
}

func New() *Storage {
	client, err := mongo.Connect(context.Background(),
		options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	database := client.Database("db")
	fmt.Println("Successfully connected to db")
	return &Storage{
		client:          client,
		database:        database,
		collectionPages: database.Collection("pages"),
	}
}

func (s *Storage) SaveLink(link model.Page) error {
	_, err := s.collectionPages.InsertOne(context.Background(), link)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) GetLinkByShortUrl(shortUrl string) (model.Page, error) {
	var link model.Page
	err := s.collectionPages.FindOne(context.Background(), bson.M{"short_url": shortUrl}).Decode(&link)
	return link, err
}

func (s *Storage) GetLinkByLongUrl(longUrl string) (model.Page, error) {
	var link model.Page
	err := s.collectionPages.FindOne(context.Background(), bson.M{"long_url": longUrl}).Decode(&link)
	return link, err
}

func (s *Storage) IsShortUrlAlreadyExists(shortUrl string) bool {
	count, _ := s.collectionPages.CountDocuments(context.Background(),
		bson.M{"short_url": shortUrl})
	if count > 0 {
		return true
	}
	return false
}

func (s *Storage) IsLongUrlAlreadyExists(longUrl string) bool {
	count, _ := s.collectionPages.CountDocuments(context.Background(),
		bson.M{"long_url": longUrl})
	if count > 0 {
		return true
	}
	return false
}
