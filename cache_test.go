package config

import (
	"fmt"
	"testing"
)

func TestCache(t *testing.T) {
	calls := 0
	provider := func(key string) (string, error) {
		calls++
		if key == "key" {
			return "value", nil
		} else {
			return "", fmt.Errorf("key not found: %s", key)
		}
	}
	Provider = Cache(provider)
	value := MustGet("key")
	if value != "value" {
		t.Errorf("expected value for key %s to be %s but got %s", "key", "value", value)
	}
	value = MustGet("key")
	if value != "value" {
		t.Errorf("expected value for key %s to be %s but got %s", "key", "value", value)
	}
	if calls != 1 {
		t.Errorf("expected %d calls but got %d calls", 1, calls)
	}
}
