package option

import (
	"flag"

	"github.com/tdl3/cloudflare-helper/cloudflare"
	"github.com/tdl3/cloudflare-helper/config"
	"github.com/tdl3/cloudflare-helper/log"
	"github.com/tdl3/cloudflare-helper/models"
	"go.uber.org/zap"
)

func HandleFlags() (cfg models.Config) {
	var configPath string
	flag.StringVar(&configPath, "config", "./config.yml", "Config file path")
	flag.Parse()

	cfg = config.LoadConfig(configPath)
	log.Init(cfg)
	zap.L().Sugar().Info("Using config ", configPath)
	if cfg.Cloudflare.DnsId == "" {
		zap.L().Info("DNS ID not specified, querying from Cloudflare")
		cloudflare.GetDnsId(&cfg)
	}
	return
}
