package utils

import (
	"encoding/json"

	"go.uber.org/zap"
)

func PrettyJson(v interface{}) string {
	jsonData, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		zap.L().Warn("JSON Parsing error", zap.Error(err))
		return ""
	}
	return string(jsonData)
}
