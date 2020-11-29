package common

type SwWeather3D struct {
	Weather []*SwWeatherData `json:"HeWeather"`
}

type SwWeatherData struct {
	Status        string          `json:"status"`
	Msg           string          `json:"msg"`
	Basic         *SwBasicInfo    `json:"basic"`
	Detail        *SwDetailedInfo `json:"detail"`
	Now           *SwNow          `json:"now"`
	Suggestions   *SwSuggestions  `json:"suggestion"`
	DailyForecast []*SwForecast   `json:"daily_forecast"`
}

type SwBasicInfo struct {
	City   string         `json:"city"`
	ID     string         `json:"id"`
	Update *SwUpdatePoint `json:"update"`
}

type SwDetailedInfo struct {
	City *SwDetailedCityInfo `json:"city"`
}

type SwDetailedCityInfo struct {
	Sunrise  string `json:"sunrise"`
	Sunset   string `json:"sunset"`
	WindDir  string `json:"win_dir"`
	WindDeg  string `json:"win_deg"`
	Pressure string `json:"pre"`
	Humidity string `json:"hum"`
}

type SwNow struct {
	Temperature string       `json:"tmp"`
	Condition   *SwCondition `json:"cond"`
}

type SwCondition struct {
	Txt string `json:"txt"`
}

type SwSuggestions struct {
	Comfort *SwSuggestionValue `json:"comf"`
	Sport   *SwSuggestionValue `json:"sport"`
}

type SwForecast struct {
	Temperature *SwTemperature `json:"tmp"`
	Date        string         `json:"date"`
	Condition   *SwCondition   `json:"cond"`
}

type SwUpdatePoint struct {
	Loc string `json:"loc"`
}

type SwSuggestionValue struct {
	Txt string `json:"txt"`
}

type SwTemperature struct {
	Max string `json:"max"`
	Min string `json:"min"`
}

func NewSwWeather3D() *SwWeather3D {
	return &SwWeather3D{}
}
