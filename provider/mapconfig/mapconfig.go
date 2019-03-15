package mapconfig

import (
	"fmt"
	"github.com/nelkinda/config-go"
)

// MapConfig creates a config.ProviderFunc based on a map.
// This is good for testing.
func MapConfig(cfg map[string]string) config.ProviderFunc {
	return func(key string) (string, error) {
		if value, ok := cfg[key]; ok {
			return value, nil
		} else {
			return "", fmt.Errorf("key %s not mapped", key)
		}
	}
}
