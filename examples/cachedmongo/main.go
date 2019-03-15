package main

import (
	"fmt"
	"github.com/nelkinda/config-go"
	"github.com/nelkinda/config-go/provider/chainconfig"
	"github.com/nelkinda/config-go/provider/envconfig"
	"github.com/nelkinda/config-go/provider/mongoconfig"
	"os"
)

func main() {
	config.Provider = config.Cache(
		chainconfig.Link(
			envconfig.Get,
			mongoconfig.CreateMongoConfigProvider(&mongoconfig.Config{}),
		),
	)
	fmt.Fprintf(os.Stdout, "Read config value %s for key %s", config.MustGet("mustExist"), "mustExist")
}
