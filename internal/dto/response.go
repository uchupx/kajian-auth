package dto

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Meta    interface{} `json:"meta"`
	Data    interface{} `json:"data"`
}

type TokenResponse struct {
	Token   string `json:"token"`
	Expired int64  `json:"expired"`
}
