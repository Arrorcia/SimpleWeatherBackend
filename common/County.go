package common

type County struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	WeatherID string `json:"weather_id"`
}
