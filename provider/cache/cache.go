package cache

import "github.com/nelkinda/config-go"

var cache = make(map[string]string)

func Cache(provider config.ProviderFunc) config.ProviderFunc {
	return func(key string) (string, error) {
		if value, ok := cache[key]; ok {
			return value, nil
		}
		if value, err := provider(key); err != nil {
			return value, err
		} else {
			cache[key] = value
			return value, nil
		}
	}
}
