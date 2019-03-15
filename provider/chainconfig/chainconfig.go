package chainconfig

import (
	"fmt"
	"github.com/nelkinda/config-go"
)

func Link(configProviders ...config.ProviderFunc) config.ProviderFunc {
	return func(key string) (string, error) {
		var e error = nil
		for _, provider := range configProviders {
			if value, err := provider(key); err == nil {
				return value, nil
			} else if e == nil {
				e = err
			} else {
				e = fmt.Errorf("%v\n%v", e, err)
			}
		}
		return "", e
	}
}
