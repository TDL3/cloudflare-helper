package main

import (
	"github.com/tdl3/cloudflare-helper/cloudflare"
	"github.com/tdl3/cloudflare-helper/option"
	"github.com/tdl3/cloudflare-helper/utils"
	"go.uber.org/zap"
)

func main() {
	cfg := option.HandleFlags()
	myIp := utils.GetMyPublicIpV4()
	defer zap.L().Sync()
	zap.L().Sugar().Info("Public IPv4 Address is ", myIp)
	code, err := cloudflare.UpdateDNSRecord(&cfg, myIp)
	if code != 200 && err != nil {
		zap.L().Fatal("Something Went Wrong, DNS Record Not Updated")
	}
	zap.L().Sugar().Info(cfg.Cloudflare.DomainName, " DNS Record Updated to ", myIp)
}
