package shortener

import (
	"context"
	fuzz "github.com/google/gofuzz"
	"perviymoiserver/pkg/model"
	"perviymoiserver/pkg/shortener/mocks"
	"testing"
)

func createTestShortener() *Shortener {
	storageMock := &mocks.Storage{}
	storageMock.On("GetLinkByShortUrl", "aboba", context.TODO()).Return(model.Page{}, nil)
	return New(storageMock)
}

func TestShortener_ShortLinkCreator(t *testing.T) {
	shortener := createTestShortener()
	result := shortener.ShortLinkCreator()
	if result == "" {
		t.Error("Result of ShortLinkCreator() empty")
	}
}

func TestShortener_RedirectToLong(t *testing.T) {
	shortener := createTestShortener()
	page := model.Page{}
	fuzzer := fuzz.New().NilChance(0)
	fuzzer.Fuzz(&page)

	err := shortener.RedirectToLong(&page, context.TODO())
	if err != nil {
		t.Error(err)
	}
}
