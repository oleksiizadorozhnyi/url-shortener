package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUrlValidator(t *testing.T) {
	testTable := []struct {
		link     string
		expected bool
	}{
		{
			link:     "https://www.youtube.com",
			expected: true,
		}, {
			link:     "https://hotline.ua",
			expected: true,
		}, {
			link:     "http",
			expected: false,
		}, {
			link:     "http:",
			expected: false,
		}, {
			link:     "http:/",
			expected: false,
		}, {
			link:     "jkada",
			expected: false,
		}, {
			link:     "//youtube.com",
			expected: false,
		}, {
			link:     "http://www.com",
			expected: true,
		}, {
			link:     "http://www.",
			expected: true,
		},
	}
	for _, testCase := range testTable {
		result := UrlValidator(testCase.link)
		t.Logf("UrlValidator(%s) returned %t", testCase.link, result)

		assert.Equal(t, testCase.expected, result,
			fmt.Sprintf("\"UrlValidator(%s) returned %t, expected %t",
				testCase.link, result, testCase.expected))
	}
}
