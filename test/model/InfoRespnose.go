package model

type InfoResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
	Token  string `json:"token"`
}
