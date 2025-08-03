package model

type ProvinceResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CityResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	ProvinceID string `json:"province_id"`
}
