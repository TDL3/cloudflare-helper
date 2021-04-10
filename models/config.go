package models

type (
	Config struct {
		Cloudflare struct {
			BearerToken string `yaml:"bearer_token"`
			ZoneId      string `yaml:"zone_id"`
			DnsId       string `yaml:"dns_id"`
			DomainName  string `yaml:"domain_name"`
		} `yaml:"cloudflare"`

		Log struct {
			Level    string `yaml:"level"`
			Encoding string `yaml:"encoding"`
		} `yaml:"log"`
		Provider struct {
			Ip []string `yaml:"ip"`
		} `yaml:"provider"`
	}
)
