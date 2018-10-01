package main

import (
	"flag"
	"fmt"
	"github.com/tanan/bq-import/config"
	"github.com/tanan/bq-import/database"
	"log"
)

var configFile = flag.String("c", "/Users/toshifumi.anan/uzabase/go/src/github.com/tanan/bq-import/config.toml", "config file path")

func main() {
	flag.Parse()

	// load config
	con, err := config.LoadFile(*configFile)
	if err != nil {
		log.Fatalf("Error while loading configuration file --- %v", err)
	}

	// connect input database
	handler, err := database.Connect(con.Database)
	if err != nil {
		fmt.Errorf("Error: %v", err)
	}

	// read data from input database
	handler.GetRow()

	// auth bigquery

	// write data for bigquery
}
