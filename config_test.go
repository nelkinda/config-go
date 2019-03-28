package config

import (
	"fmt"
	"testing"
)
func initProvider() {
	Provider = func(key string) (value string, e error) {
		if key == "key" {
			return "value", nil
		} else {
			return "", fmt.Errorf("can't find key %s", key)
		}
	}
}

func TestGetMapped(t *testing.T) {
	initProvider()
	if val, err := Get("key"); err != nil {
		t.Errorf("Unexpected error %s", err)
	} else if val != "value"{
		t.Errorf("Expected %s but got %s", "value", val)
	}
}

func TestGetUnmapped(t *testing.T) {
	initProvider()

	if val, err := Get("unknown"); err == nil {
		t.Errorf("expected error but got success for key %s with value %s", "unknown", val)
	}
}

func TestMustMapped(t *testing.T) {
	initProvider()

	MustGet("key")
}

func TestGetOrDefaultMapped(t *testing.T) {
	initProvider()

	if val := GetOrDefault("key", "defaultValue"); val != "value" {
		t.Errorf("Expected %s but got %s", "value", val)
	}
}

func TestGetOrDefaultUnmapped(t *testing.T) {
	initProvider()

	if val := GetOrDefault("unknown", "defaultValue"); val != "defaultValue" {
		t.Errorf("Expected %s but got %s", "defaultValue", val)
	}
}

func TestMustGetUnmapped(t *testing.T) {
	initProvider()

	defer func() {
		recover()
	}()

	MustGet("foo")
	t.Errorf("Expected panic")
}
