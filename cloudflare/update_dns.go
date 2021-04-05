package cloudflare

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/tdl3/cloudflare-helper/models"
	"github.com/tdl3/cloudflare-helper/utils"
	"go.uber.org/zap"
)

func UpdateDNSRecord(cfg *models.Config, newIp string) (StatusCode int) {
	payload := models.DNS{
		Type:    "A",
		Name:    "ts.tdl3.com",
		Content: newIp,
		TTL:     1,
		Proxied: false,
	}

	client := &http.Client{}
	data, err := json.Marshal(payload)
	if err != nil {
		zap.L().Fatal("Can not parse json", zap.Error(err))
	}
	url := utils.AssembleUrl(*cfg, utils.UpdateDns)
	request, err := http.NewRequest("PATCH", url, bytes.NewBufferString(string(data)))
	if err != nil {
		zap.L().Fatal("Something went wrong", zap.Error(err))
	}
	request.Header.Add("Authorization", cfg.Cloudflare.BearerToken)
	request.Header.Add("Content-Type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		zap.L().Fatal("Con not resolve host", zap.Error(err))
	}
	defer response.Body.Close()
	var content interface{}
	err = json.NewDecoder(response.Body).Decode(&content)
	if err != nil {
		zap.L().Info("Can not parse response", zap.Error(err))
	}
	zap.L().Debug(utils.PrettyJson(content))

	return response.StatusCode
}
