package main

import (
	"flag"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

//Config is an exported type
type Config struct {
	Server   string
	Name     string
	Channels []string
}

func loadDefaults() Config {
	var defaultConfig Config
	defaultConfig.Server = "irc.freenode.net:6667"
	defaultConfig.Name = "gobot"
	defaultConfig.Channels[0] = "##gobot"
	return defaultConfig
}

func readConfig() Config {
	var configfile string
	var config Config

	flag.StringVar(&configfile, "config", "./config.toml", "Config file toml file for bot")
	flag.Parse()

	_, err := os.Stat(configfile)
	if err != nil {
		config = loadDefaults()
	} else {

		if _, errs := toml.DecodeFile(configfile, &config); errs != nil {
			log.Fatal(errs)
		}
	}
	return config
}
