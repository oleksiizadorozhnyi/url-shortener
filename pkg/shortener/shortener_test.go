package shortener

import (
	"perviymoiserver/pkg/shortener/mocks"
	"testing"
)

func createTestShortener() *Shortener {
	return New(&mocks.Storage{})
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
	err := shortener.RedirectToLong()
}
