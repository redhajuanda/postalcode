package main

type Province struct {
	ProvinceCode string `json:"province_code"`
	ProvinceName string `json:"province_name"`
}

type PostalCode struct {
	Urban       string   `json:"urban"`
	SubDistrict string   `json:"sub_district"`
	City        string   `json:"city"`
	PostalCode  string   `json:"postal_code"`
	Province    Province `json:"province"`
}

type Response struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	Data    []PostalCode `json:"data"`
}
