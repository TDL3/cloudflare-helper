package config

import (
	"log"
	"os"

	"github.com/tdl3/cloudflare-helper/models"
	"gopkg.in/yaml.v3"
)

func loadFromFile(path string, cfg *models.Config) {
	if path == "" {
		path = "./config.yml"
	}
	f, err := os.Open(path)
	if err != nil {
		log.Fatal("Can not find config file", err)
	}
	defer f.Close()
	var data []byte
	f.Read(data)
	decoder := yaml.NewDecoder(f)
	decoder.KnownFields(true)
	err = decoder.Decode(cfg)
	if err != nil {
		log.Fatal("Invalid options", err)
	}
}

func LoadConfig(path string) (cfg models.Config) {
	loadFromFile(path, &cfg)
	return
}
