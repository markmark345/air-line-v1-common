package model

type Lookup struct {
	Key      string `json:"key"`
	HTTPCode int    `json:"http_code"`
	Code     int    `json:"code"`
	DescEN   string `json:"descEN"`
	DescTH   string `json:"descTH"`
}
