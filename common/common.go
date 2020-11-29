package common

import "fmt"

// He2Sw returns a SwWeather3D built from HeWeather3D
func He2Sw(h *HeWeather3D) (*SwWeather3D, error) {
	if h == nil {
		return nil, fmt.Errorf("h is empty")
	}
	sw := new(SwWeather3D)
	sw.Weather = make([]*SwWeatherData, 1)
	sw.Weather[0] = &SwWeatherData{
		Status:        "",
		Msg:           "",
		Basic:         nil,
		Detail:        nil,
		Now:           nil,
		Suggestions:   nil,
		DailyForecast: nil,
	}
	swd := sw.Weather[0]
	if h.Code != "200" {
		swd.Status = "Error"
		swd.Msg = fmt.Sprintf("code:%s", h.Code)
		return nil, fmt.Errorf("h's code is not 200")
	}
	swd.Status = "Ok"
	swd.Msg = ""
	swd.Basic = &SwBasicInfo{
		// TODO: fill city and id field
		City: "",
		ID:   "",
		Update: &SwUpdatePoint{
			// TODO: convert time format
			// current: 2020-11-28T19:35+08:00
			// target: 2020-11-28 19:35
			Loc: h.UpdateTime,
		},
	}
	cdd := h.Daily[0]
	swd.Detail = &SwDetailedInfo{
		City: &SwDetailedCityInfo{
			Sunrise:  cdd.Sunrise,
			Sunset:   cdd.Sunset,
			WindDir:  cdd.WindDirDay,
			WindDeg:  cdd.WindScaleDay,
			Pressure: cdd.Pressure,
			Humidity: cdd.Humidity,
		},
	}
	swd.Now = &SwNow{
		// TODO:
		Temperature: "-1",
		Condition: &SwCondition{
			Txt: cdd.TextDay,
		},
	}
	swd.Suggestions = &SwSuggestions{
		Comfort: &SwSuggestionValue{
			Txt: "comfort",
		},
		Sport: &SwSuggestionValue{
			Txt: "sport",
		},
	}
	swd.DailyForecast = make([]*SwForecast, 2)
	for i := 0; i < 2; i++ {
		cdd = h.Daily[i+1]
		swd.DailyForecast[i] = &SwForecast{
			Temperature: &SwTemperature{
				Max: cdd.TempMax,
				Min: cdd.TempMin,
			},
			Date: cdd.FxDate,
			Condition: &SwCondition{
				Txt: cdd.TextDay,
			},
		}
	}
	return sw, nil
}
