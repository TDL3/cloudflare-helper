package utils

import (
	"encoding/json"
	"log"
)

func PrettyJson(v interface{}) {
	jsonData, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		log.Print("JSON Parsing error", err)
		return
	}
	log.Print(string(jsonData))
}
