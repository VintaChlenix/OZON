package handlers

import (
	"OZON/internal/db"
	"github.com/eknkc/basex"
	"testing"
)

const ALPHABET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"

func TestURLEncode(t *testing.T) {
	//arrange
	testTable := []struct {
		input    string
		expected string
	}{
		{
			input:    "http://vk.com",
			expected: "bGwXqTTUYK",
		},
		{
			input:    "http://google.com",
			expected: "wNcPGimRlA",
		},
		{
			input:    "http://ozon.ru",
			expected: "bWCqEvgLud",
		},
	}

	code, _ := basex.NewEncoding(ALPHABET)
	h := NewAddHandler(db.NewInMemory(), code)
	//act
	for _, testCase := range testTable {
		result := URLEncode(testCase.input, *h)

		if result != testCase.expected {
			t.Errorf("Incorrect result")
		}
	}
}
