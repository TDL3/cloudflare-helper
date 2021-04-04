package utils

import (
	"fmt"

	"github.com/tdl3/cloudflare-helper/consts"
)

type Schema int

const (
	UpdateDns Schema = iota
	GetDnsId  Schema = iota
)

func AssembleUrl(s Schema) (url string) {
	switch s {
	case UpdateDns:
		return updateDns()
	case GetDnsId:
		return getDnsId()
	}
	return ""
}

func updateDns() (url string) {
	// https://api.cloudflare.com/client/v4/zones/023e105f4ecef8ad9ca31a8372d0c353/dns_records/372e67954025e0ba6aaa6d586b9e0b59
	path := fmt.Sprintf("%s/dns_records/%s", consts.ZoneId, consts.DnsId)
	return consts.CloudflareEndPoint + path
}

func getDnsId() (url string) {
	path := fmt.Sprintf("%s/dns_records%s", consts.ZoneId, "?type=A&name="+consts.DomainName)
	return consts.CloudflareEndPoint + path
}
