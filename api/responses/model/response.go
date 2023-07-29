package model

type Response struct {
	Status `json:"status"`
	Data   struct{} `json:"data,omitempty"`
}

type Status struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
	Details     `json:"data,omitempty"`
}

type Details struct {
	Source string `json:"source,omitempty"`
	Reason string `json:"reason,omitempty"`
	Error  error  `json:"error,omitempty"`
}
