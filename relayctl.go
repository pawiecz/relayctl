package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Pins PinMap
}

var (
	conf     Config
	confFile string
	address  string
)

func setFlags() {
	flag.StringVar(&confFile, "config", "pins.toml", "TOML file with relay to pin mapping")
	flag.StringVar(&address, "address", ":8080", "API server listen address")
}

func main() {
	setFlags()
	flag.Parse()
	if _, err := toml.DecodeFile(confFile, &conf); err != nil {
		log.Fatal(err)
	}

	SetupModule(conf.Pins)
	defer CloseModule()

	router := NewRouter(AllRoutes())
	log.Fatal(http.ListenAndServe(address, router))
}
