package bq_import

import (
	"flag"
	"fmt"
	"github.com/tanan/bq-import/config"
)

var configFile = flag.String("c", "config.toml", "config file path")

func main() {
	flag.Parse()
	con, err := config.LoadFile(*configFile)
	if err != nil {
		fmt.Errorf("Error while loading configuration file --- %v", err)
	}
	fmt.Println("config:", con)
	fmt.Println("hello, world.")
}
