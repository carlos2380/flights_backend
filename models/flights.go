package models

type Flight struct {
	Airline     string `json:"airline"`
	Flight      string `json:"flight"`
	Aircraft    string `json:"aircraft"`
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
	HourDep     string `json:"hour_dep"`
	HourArr     string `json:"hour_arr"`
	Status      string `json:"status"`
}
