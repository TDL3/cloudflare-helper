package cloudflare

import (
	"encoding/json"
	"net/http"

	"github.com/tdl3/cloudflare-helper/models"
	"github.com/tdl3/cloudflare-helper/models/resp"
	"github.com/tdl3/cloudflare-helper/utils"
	"go.uber.org/zap"
)

// This will get the dns_id from cloudflare, and store it in the cfg
func GetDNSId(cfg *models.Config) (StatusCode int, err error) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", utils.AssembleUrl(*cfg, utils.GetDnsId), nil)
	if err != nil {
		panic(err)
	}
	request.Header.Add("Authorization", cfg.Cloudflare.BearerToken)
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
	zap.L().Debug(utils.PrettyJson(content.Result[0]))
	cfg.Cloudflare.DnsId = id
	return response.StatusCode, nil
}
