package shortener

import (
	"errors"
	"math/rand"
	"perviymoiserver/pkg/adapter/storage/mongo"
	"perviymoiserver/pkg/model"
	"time"
)

type Shortener struct {
	storage *mongo.Storage
}

func New(storage *mongo.Storage) *Shortener {
	return &Shortener{
		storage: storage,
	}
}

func (s *Shortener) NewLink(newPage *model.Page) error {
	for i := 0; i < 100; i++ {
		newPage.ShortUrl = s.ShortLinkCreator()
		if s.storage.IsAlreadyExists(newPage.ShortUrl) {
			if i == 99 {
				return errors.New("can't make short link")
			}
			continue
		}
		break
	}
	err := s.storage.SaveLink(*newPage)
	if err != nil {
		return err
	}
	return nil
}

func (s *Shortener) RedirectToLong(link *model.Page) error {
	var err error
	*link, err = s.storage.GetLinkByShortUrl(link.ShortUrl)
	if err != nil {
		return err
	}
	return nil
}

func (s *Shortener) ShortLinkCreator() string {
	rand.Seed(time.Now().UnixNano())
	length := 5
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
