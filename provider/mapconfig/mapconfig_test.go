package mapconfig

import (
	"testing"
)

func TestMapConfig(t *testing.T) {
	provider := MapConfig(map[string]string{"key1": "value1", "key2": "value2"})
	if val, err := provider("key1"); err != nil {
		t.Error(err)
	} else if val != "value1" {
		t.Errorf("expected key %s to have value %s but got %s", "key1", "value1", val)
	}

	if val, err := provider("key2"); err != nil {
		t.Error(err)
	} else if val != "value2" {
		t.Errorf("expected key %s to have value %s but got %s", "key2", "value2", val)
	}

	if _, err := provider("key3"); err == nil {
		t.Errorf("Expected key %s to be unmapped", "key3")
	}
}
