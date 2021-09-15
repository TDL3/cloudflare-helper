package models

type (
	Config struct {
		Cloudflare `yaml:"cloudflare"`
		Log        `yaml:"log"`
		Provider   `yaml:"provider"`
	}
	Cloudflare struct {
		BearerToken string `yaml:"bearer_token"`
		ZoneId      string `yaml:"zone_id"`
		DnsId       string `yaml:"dns_id"`
		DomainName  string `yaml:"domain_name"`
	}

	Log struct {
		Level    string `yaml:"level"`
		Encoding string `yaml:"encoding"`
	}
	Provider struct {
		Ipv4 []string `yaml:"ipv4"`
	}
)
