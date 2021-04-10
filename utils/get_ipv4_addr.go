package utils

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/tdl3/cloudflare-helper/models"
	"go.uber.org/zap"
)

func GetMyPublicIpV4(cfg models.Config) (ipv4 string) {
	providers := cfg.Provider.Ip
	for _, v := range providers {
		code := getIpv4(v, &ipv4)
		if code == http.StatusOK {
			return
		}
		zap.L().Warn("Can not get ipv4 address tring next provider", zap.String("provider", v))
	}
	return
}

func getIpv4(provider string, ipv4 *string) (statusCode int) {
	var client http.Client
	resp, err := client.Get(provider)
	if err != nil {
		zap.L().Fatal("Can not resolve host", zap.Error(err))
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			zap.L().Fatal("Can not parse response", zap.Error(err))
		}
		*ipv4 = strings.Trim(string(bodyBytes), "\n")
		return resp.StatusCode
	}
	return
}
