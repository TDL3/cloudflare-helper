package models

type Config struct {
	Cloudflare struct {
		BearerToken string `yaml:"bearer_token"`
		ZoneId      string `yaml:"zone_id"`
		DnsId       string `yaml:"dns_id"`
		DomainName  string `yaml:"domain_name"`
	} `yaml:"cloudflare"`
}
