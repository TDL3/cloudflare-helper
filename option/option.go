package option

import (
	"flag"

	"github.com/tdl3/cloudflare-helper/config"
	"github.com/tdl3/cloudflare-helper/models"
)

func HandleFlags() (cfg models.Config) {
	var configPath string
	flag.StringVar(&configPath, "config", "./config.yml", "Config file path")
	flag.Parse()
	cfg = config.LoadConfig(configPath)
	return
}
