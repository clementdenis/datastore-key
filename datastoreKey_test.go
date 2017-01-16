package main

import (
	"testing"
	"regexp"
)

func TestReplaceInLine(t *testing.T) {
	line := `"key": "aglzfmFvLWRvY3NyKAsSBkZvbGRlciIcMEIwMk9nN1FrTWhJTk9HOXRaV2RLYzNkWmVVawyiAQhhdmV2ZS5iZQ"`
	compiled := regexp.MustCompile(`"key": "([^"]+)"`)
	field := "Name"
	result := replaceKey(line, compiled, field)
	if result != `"key": "0B02Og7QkMhINOG9tZWdKc3dZeUk"` {
		t.Errorf("Not matching: %s", result)
	}
}
