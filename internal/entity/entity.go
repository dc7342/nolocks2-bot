package entity

type Location struct {
	Lon      float64 `json:"longitude"`
	Lat      float64 `json:"latitude"`
	Comment  string  `json:"comment"`
	ImageUrl string  `json:"image"`
}

func NewLocation(lon float64, lat float64) *Location {
	return &Location{Lon: lon, Lat: lat}
}

type NoLocksConfig struct {
	EndpointURL string
	User        string
	Pass        string
}
