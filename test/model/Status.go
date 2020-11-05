package model

type Status struct {
	Status      string `json:"status"`
	Code        int    `json:"code"`
	Description string `json:"description"`
}
