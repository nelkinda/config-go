package chainconfig

import (
	"github.com/nelkinda/config-go/provider/mapconfig"
	"testing"
)

func TestChainConfig(t *testing.T) {
	link1 := mapconfig.MapConfig(map[string]string{"key1": "value1"})
	link2 := mapconfig.MapConfig(map[string]string{"key2": "value2"})

	chain := Link(link1, link2)

	if val, err := chain("key1"); err != nil {
		t.Error(err)
	} else if val != "value1" {
		t.Errorf("For key %s, expected value %s but got %s", "key1", "value1", val)
	}

	if val, err := chain("key2"); err != nil {
		t.Error(err)
	} else if val != "value2" {
		t.Errorf("For key %s, expected value %s but got %s", "key2", "value2", val)
	}

	if _, err := chain("key3"); err == nil {
		t.Errorf("Expected key %s to be unmapped", "key3")
	}
}
