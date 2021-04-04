package utils

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tdl3/cloudflare-helper/constant"
)

func GetMyPublicIpV4() (ipv4 string) {
	var client http.Client
	resp, err := client.Get(constant.GetIpv4Addr)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		ipv4 := string(bodyBytes)
		return ipv4
	}
	return ""
}
