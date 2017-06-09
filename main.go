package main

import (
	"flag"
	"fmt"

	"github.com/jinzhu/configor"
)

type Config struct {
	Port  string `default:"11300"`
	Debug bool
}

func main() {
	cfg := &Config{}
	flag.StringVar(&cfg.Port, "port", "11333", "port to use")
	flag.BoolVar(&cfg.Debug, "debug", false, "run in debug mode")
	flag.Parse()

	configor.New(&configor.Config{ENVPrefix: "TEST"}).Load(cfg, "config.yaml")

	fmt.Printf("%#v", cfg)
}
