package config

var Provider ProviderFunc

type ProviderFunc func(string) (string, error)

func Get(key string) (string, error) {
	return Provider(key)
}

func GetOrDefault(key string, defaultValue string) string {
	if value, err := Get(key); err != nil {
		return defaultValue
	} else {
		return value
	}
}

func MustGet(key string) string {
	if value, err := Get(key); err != nil {
		panic(err)
	} else {
		return value
	}
}
