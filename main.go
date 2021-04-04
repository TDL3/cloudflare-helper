package main

import (
	"log"

	"github.com/tdl3/cloudflare-helper/cloudflare"
	"github.com/tdl3/cloudflare-helper/option"
	"github.com/tdl3/cloudflare-helper/utils"
)

func main() {
	cfg := option.HandleFlags()
	myIp := utils.GetMyPublicIpV4()
	log.Print("Public IPv4 Address ", myIp)
	code, err := cloudflare.UpdateDNSRecord(&cfg, myIp)
	if code != 200 && err != nil {
		log.Fatal("Something Went Wrong, DNS Record Not Updated")
		return
	}
	log.Print(cfg.Cloudflare.DomainName, " DNS Record Updated to ", myIp)
}
