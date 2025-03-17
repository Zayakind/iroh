package models

type Aircraft struct {
	Code  string `json:"aircraft_code"`
	Model string `json:"model"`
	Range string `json:"range"`
}
