package entity

type Location struct {
	Lon      string `json:"longitude"`
	Lat      string `json:"latitude"`
	Comment  string `json:"comment"`
	ImageUrl string `json:"image"`
}

func NewLocation(lon string, lat string) *Location {
	return &Location{Lon: lon, Lat: lat}
}

type NoLocksConfig struct {
	EndpointURL string
	User        string
	Pass        string
}
