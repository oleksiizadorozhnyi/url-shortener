package utils

import (
	"net/url"
)

func UrlValidator(longUrl string) bool {
	_, err := url.ParseRequestURI(longUrl)
	if err != nil {
		return false
	}
	u, err := url.Parse(longUrl)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}
	return true
}
