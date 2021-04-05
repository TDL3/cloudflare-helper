package utils

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/tdl3/cloudflare-helper/constant"
	"go.uber.org/zap"
)

func GetMyPublicIpV4() (ipv4 string) {
	var client http.Client
	resp, err := client.Get(constant.GetIpv4Addr)
	if err != nil {
		zap.L().Fatal("Can not connect to internet", zap.Error(err))
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			zap.L().Fatal("Can not parse response", zap.Error(err))
		}
		ipv4 := strings.Trim(string(bodyBytes), "\n")
		return ipv4
	}
	return ""
}
