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
	configor.New(&configor.Config{ENVPrefix: "TEST"}).Load(cfg, "config.yaml")

	flag.StringVar(&cfg.Port, "port", cfg.Port, "port to use")
	flag.BoolVar(&cfg.Debug, "debug", cfg.Debug, "run in debug mode")
	flag.Parse()

	fmt.Printf("%#v", cfg)
}
