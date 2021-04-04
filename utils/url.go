package utils

import (
	"fmt"

	"github.com/tdl3/cloudflare-helper/constant"
	"github.com/tdl3/cloudflare-helper/models"
)

type Schema int

const (
	UpdateDns Schema = iota
	GetDnsId  Schema = iota
)

func AssembleUrl(cfg models.Config, s Schema) (url string) {
	switch s {
	case UpdateDns:
		return updateDns(cfg)
	case GetDnsId:
		return getDnsId(cfg)
	}
	return ""
}

func updateDns(cfg models.Config) (url string) {
	path := fmt.Sprintf("%s/dns_records/%s", cfg.Cloudflare.ZoneId, cfg.Cloudflare.DnsId)
	return constant.CloudflareEndPoint + path
}

func getDnsId(cfg models.Config) (url string) {
	path := fmt.Sprintf("%s/dns_records%s", cfg.Cloudflare.ZoneId, "?type=A&name="+cfg.Cloudflare.DomainName)
	return constant.CloudflareEndPoint + path
}
