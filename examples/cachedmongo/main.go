package main

import (
	"fmt"
	"github.com/nelkinda/config-go"
	"github.com/nelkinda/config-go/provider/cache"
	"github.com/nelkinda/config-go/provider/mongoconfig"
	"os"
)

func main() {
	config.Provider = cache.Cache(mongoconfig.CreateMongoConfigProvider("", "config", "config"))
	fmt.Fprintf(os.Stdout, "Read config value %s for key %s", config.MustGet("mustExist"), "mustExist")
}
