package cloudflare

import (
	"encoding/json"
	"net/http"

	"github.com/tdl3/cloudflare-helper/consts"
	"github.com/tdl3/cloudflare-helper/models/resp"
	"github.com/tdl3/cloudflare-helper/utils"
)

func GetDNSId() (StatusCode int, err error) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", utils.AssembleUrl(utils.GetDnsId), nil)
	if err != nil {
		panic(err)
	}
	request.Header.Add("Authorization", consts.BearerToken)
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	var content resp.DNSId
	err = json.NewDecoder(response.Body).Decode(&content)
	if err != nil {
		return response.StatusCode, err
	}
	id := content.Result[0].Id
	// utils.PrettyJson(id)
	consts.DnsId = id
	return response.StatusCode, nil
}
