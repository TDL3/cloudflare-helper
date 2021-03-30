package option

import (
	"flag"
	"log"

	"github.com/tdl3/cloudflare-helper/config"
)

func HandleFlags() {
	var configPath string
	flag.StringVar(&configPath, "config", "./config.yml", "Config file path")
	flag.Parse()
	config.LoadConfig(configPath)
	log.Print("Using config ", configPath)
}
