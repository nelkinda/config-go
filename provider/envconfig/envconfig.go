package envconfig

import (
	"fmt"
	"os"
)

func Get(key string) (string, error) {
	if value, ok := os.LookupEnv(key); !ok {
		return "", fmt.Errorf("variable %s not defined", key)
	} else {
		return value, nil
	}
}
