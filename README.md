# Configuration for Go

A configuration library for Golang.
Supports configuration caching and reading configuration from MongoDB or CosmosDB.
It's modular, you can write your own configuration provider to get configuration from other sources, like DynamoDB, files, or existing configuration servers.

## Usage

```go
package main

import (
	"github.com/nelkinda/config-go"
	"github.com/nelkinda/config-go/provider/cache"
	"github.com/nelkinda/config-go/provider/mongoconfig"
)

func main() {
	// Setup the provider to be a cached mongodb
	config.Provider = cache.Cache(mongoconfig.CreateMongoConfigProvider("mongodb://confighost/", "config", "config"))

	key := "my-configuration-key"

	// Use optional configuration values
	value, err := config.Get(key)

	// Use mandatory configuration values
	value := config.MustGet(key)
}
```