package model

type Page struct {
	LongUrl  string `bson:"long_url"`
	ShortUrl string `bson:"short_url"`
}
