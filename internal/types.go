package internal

type Outage struct {
	Info      string `json:"info"`
	ID        string `json:"id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Zones     []Zone `json:"zones"`
}

type Zone struct {
	District      string   `json:"district"`
	Neighborhoods []string `json:"neighborhoods"`
}

type ApiResponse struct {
	Features []Feature `json:"features"`
}

type Feature struct {
	Properties Properties `json:"properties"`
}

type Properties struct {
	IDInfo       string `json:"ARIZA_NO"`
	Description  string `json:"ARIZA_NEVI_ACIKLAMASI"`
	District     string `json:"ILCE_ADI"`
	Neighborhood string `json:"MAHALLE_ADI"`
	StartDate    string `json:"BASLAMA_TARIHI"`
	EndDate      string `json:"TAHMINI_BITIS_TARIHI"`
}
