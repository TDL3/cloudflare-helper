package config

import (
	"log"
	"os"

	"github.com/tdl3/cloudflare-helper/consts"
	"github.com/tdl3/cloudflare-helper/models"
	"gopkg.in/yaml.v3"
)

func loadFromFile(path string, cfg *models.Config) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal("Can not find config file")
		panic(err)
	}
	defer f.Close()
	var data []byte
	f.Read(data)
	decoder := yaml.NewDecoder(f)
	decoder.KnownFields(true)
	err = decoder.Decode(cfg)
	if err != nil {
		log.Fatal("invalid options")
		panic(err)
	}
}

func LoadConfig(path string) {
	var cfg models.Config
	loadFromFile(path, &cfg)
	consts.BearerToken = cfg.Cloudflare.BearerToken
	consts.ZoneId = cfg.Cloudflare.ZoneId
	consts.DnsId = cfg.Cloudflare.DnsId
	consts.DomainName = cfg.Cloudflare.DomainName
}
