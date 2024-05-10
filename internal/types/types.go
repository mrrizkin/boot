package types

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Debug   string      `json:"debug"`
	Data    interface{} `json:"data"`
}
