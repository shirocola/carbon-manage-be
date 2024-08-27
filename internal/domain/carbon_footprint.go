package domain

type CarbonFootprint struct {
	ID       string  `json:"id"`
	Emission float64 `json:"emission"`
	Sources  string  `json:"sources"`
	Date     string  `json:"date"`
}
