package common

type GlWeather7D struct {
	Weather []*GlWeatherData `json:"HeWeather"`
}

type GlWeatherData struct {
	Status        string         `json:"status"`
	Msg           string         `json:"msg"`
	Basic         *GlBasicInfo   `json:"basic"`
	Detail        *GlAQIInfo     `json:"aqi"`
	Now           *GlNow         `json:"now"`
	Suggestions   *GlSuggestions `json:"suggestion"`
	DailyForecast []*GlForecast  `json:"daily_forecast"`
}

type GlBasicInfo struct {
	CID        string         `json:"cid"`
	Location   string         `json:"location"`
	ParentCity string         `json:"parent_city"`
	AdminArea  string         `json:"admin_area"`
	Cnty       string         `json:"cnty"`
	Lat        string         `json:"lat"`
	Lon        string         `json:"lon"`
	TZ         string         `json:"tz"`
	City       string         `json:"city"`
	ID         string         `json:"id"`
	Update     *GlUpdatePoint `json:"update"`
}

type GlAQIInfo struct {
	City *GlAQICityInfo `json:"city"`
}

type GlAQICityInfo struct {
	AQI     string `json:"aqi"`
	PM25    string `json:"pm25"`
	Quality string `json:"qlty"`
}

type GlNow struct {
	Cloud       string          `json:"cloud"`
	CondCode    string          `json:"cond_code"`
	CondTxt     string          `json:"cond_txt"`
	Fl          string          `json:"fl"`
	Humidity    string          `json:"hum"`
	Pcpn        string          `json:"pcpn"`
	Pres        string          `json:"pres"`
	Vis         string          `json:"vis"`
	WindDeg     string          `json:"wind_deg"`
	WindDir     string          `json:"wind_dir"`
	WindSc      string          `json:"wind_sc"`
	WindSpd     string          `json:"wind_spd"`
	Temperature string          `json:"tmp"`
	Condition   *GlNowCondition `json:"cond"`
}

type GlNowCondition struct {
	Txt  string `json:"txt"`
	Code string `json:"code"`
}

type GlSuggestions struct {
	Comfort *GlSuggestionValue `json:"comf"`
	Sport   *GlSuggestionValue `json:"sport"`
	Cw      *GlSuggestionValue `json:"cw"`
}

type GlForecast struct {
	Temperature *GlTemperature    `json:"tmp"`
	Date        string            `json:"date"`
	Condition   *GlDailyCondition `json:"cond"`
}

type GlDailyCondition struct {
	Txt string `json:"txt_d"`
}

type GlUpdatePoint struct {
	Loc string `json:"loc"`
	Utc string `json:"utc"`
}

type GlSuggestionValue struct {
	Txt   string `json:"txt"`
	Type  string `json:"type"`
	Brief string `json:"brf"`
}

type GlTemperature struct {
	Max string `json:"max"`
	Min string `json:"min"`
}

func NewGlWeather7D() *GlWeather7D {
	return &GlWeather7D{}
}
