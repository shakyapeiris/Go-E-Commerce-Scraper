package api

type Response struct {
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Success    bool        `json:"success"`
	StatusCode int         `json:"status_code"`
}
