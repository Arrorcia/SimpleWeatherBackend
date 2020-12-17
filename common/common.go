package common

import (
	"fmt"
	"strconv"
	"strings"
)

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
		City: "",
		ID:   "",
		Update: &SwUpdatePoint{
			Loc: "",
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
		Temperature: func() string {
			max, _ := strconv.Atoi(cdd.TempMax)
			min, _ := strconv.Atoi(cdd.TempMin)
			mid := (max + min) / 2
			return strconv.Itoa(mid)
		}(),
		Condition: &SwCondition{
			Txt: cdd.TextDay,
		},
	}
	swd.Suggestions = &SwSuggestions{
		Comfort: &SwSuggestionValue{
			Txt: "无建议",
		},
		Sport: &SwSuggestionValue{
			Txt: "无建议",
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

func Gl2Sw(g *GlWeather7D) (*SwWeather3D, error) {
	if g == nil {
		return nil, fmt.Errorf("h is empty")
	}
	sw := NewSwWeather3D()
	sw.Weather = make([]*SwWeatherData, 1)
	glw := g.Weather[0]
	sw.Weather[0] = &SwWeatherData{
		Status:        glw.Status,
		Msg:           glw.Msg,
		Basic:         nil,
		Detail:        nil,
		Now:           nil,
		Suggestions:   nil,
		DailyForecast: nil,
	}
	swd := sw.Weather[0]
	swd.Basic = &SwBasicInfo{
		City: glw.Basic.City,
		ID:   strings.TrimPrefix(glw.Basic.ID, "CN"),
		Update: &SwUpdatePoint{
			Loc: glw.Basic.Update.Loc,
		},
	}
	swd.Detail = &SwDetailedInfo{
		City: &SwDetailedCityInfo{
			Sunrise:  "", // Gl数据没有该字段
			Sunset:   "", // Gl数据没有该字段
			WindDir:  glw.Now.WindDir,
			WindDeg:  glw.Now.WindDeg,
			Pressure: glw.Now.Pres,
			Humidity: glw.Now.Humidity,
		},
	}
	swd.Now = &SwNow{
		Temperature: glw.Now.Temperature,
		Condition: &SwCondition{
			Txt: glw.Now.Condition.Txt,
		},
	}
	swd.Suggestions = &SwSuggestions{
		Comfort: &SwSuggestionValue{Txt: glw.Suggestions.Comfort.Txt},
		Sport:   &SwSuggestionValue{Txt: glw.Suggestions.Sport.Txt},
	}
	swd.DailyForecast = make([]*SwForecast, 2)
	for i := 0; i < 2; i++ {
		cdd := glw.DailyForecast[i]
		swd.DailyForecast[i] = &SwForecast{
			Temperature: &SwTemperature{
				Max: cdd.Temperature.Max,
				Min: cdd.Temperature.Min,
			},
			Date:      cdd.Date,
			Condition: &SwCondition{Txt: cdd.Condition.Txt},
		}
	}
	return sw, nil
}

func GetSuggestions(tmp int) (string, string, error) {
	var comfSug, sportSug string
	switch {
	case tmp < 10:
		comfSug += "气温较低，注意保暖，不宜外出。"
		sportSug += "气温过低，不宜进行户外运动。"
	case tmp > 30:
		comfSug += "气温较高，注意散热，不宜外出。"
		sportSug += "气温过高，不宜进行户外运动。"
	default:
		comfSug += "气温适中，天气较好，适合外出。"
		sportSug += "气温适宜，建议外出活动。"
	}
	return comfSug, sportSug, nil
}
