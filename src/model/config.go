package model

type Config struct {
	Id         int    `json:"id"`
	Upper_code string `json:"upper_code"`
	Upper_name string `json:"upper_name"`
	Code       string `json:"code"`
	Name       string `json:"name"`
}
