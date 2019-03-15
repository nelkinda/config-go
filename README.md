# Configuration for Go

A configuration library for Golang.
Supports configuration caching and reading configuration from MongoDB or CosmosDB.
It's modular, you can write your own configuration provider to get configuration from other sources, like DynamoDB, files, or existing configuration servers.

## Status
- MVP: 100% done

## List of providers
- Cache, decorates another provider to provide caching of configuration values.
  Typically used in conjunction with a database provider.
- Chain, links multiple providers.
  Typically used to link environment before a database to allow environment configuration to override database configuration.
- Env, reads configuration values from the environment
- Map, a simple go-map based provider, useful for testing
- MongoDB, reads configuration values from a MongoDB or CosmosDB document (see below)

## Usage

```go
package main

import (
	"github.com/nelkinda/config-go"
	"github.com/nelkinda/config-go/provider/chainconfig"
	"github.com/nelkinda/config-go/provider/envconfig"
	"github.com/nelkinda/config-go/provider/mongoconfig"
)

func main() {
	// Setup the provider to be a cached mongodb
	config.Provider = config.Cache(chainconfig.Link(envconfig.Get, mongoconfig.CreateMongoConfigProvider(&mongoconfig.Config{})))

	key := "my-configuration-key"

	// Use optional configuration values
	value, err := config.Get(key)

	// Use mandatory configuration values
	value := config.MustGet(key)
}
```

## MongoDB / CosmosDB
You can use MongoDB or CosmosDB as configuration provider.
Actually, as of now, MongoDB / CosmosDB is the only provider provided by this library.
The database schema is as follows:
```json
[
  {
    "environment": "nameOfTheEnvironment",
    "configuration": {
      "key1": "value1",
      "key2": "value2"
    }
  }
]
```

## Handling multiple environments
One of the main purposes of having a configuration module like this is to support running the same software or container in multiple environments.
A typical use case would thus be to determine the actual environment from an environment variable.
