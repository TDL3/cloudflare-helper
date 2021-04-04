package resp

type DNSId struct {
	Result []struct {
		Content string `json:"content"`
		Id      string `json:"id"`
		Locked  bool   `json:"locked"`
		Name    string `json:"name"`
	} `json:"result"`
}
