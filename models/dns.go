package models

type DNS struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	Content string `json:"content"`
	TTL     int64  `json:"ttl"`
	Proxied bool   `json:"proxied"`
}
