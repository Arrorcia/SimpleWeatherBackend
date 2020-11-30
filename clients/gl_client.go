package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ame-lm/SimpleWeather/common"
	"net/http"
	"strings"
)

const ChinaPrefix = "http://guolin.tech/api/china"
const WeatherPrefix = "http://guolin.tech/api/weather"

type GlClientImpl struct {
	key string
}

func NewGlClientImpl() *GlClientImpl {
	return &GlClientImpl{"bc0418b57b2d4918819d3974ac1285d9"}
}

func (g *GlClientImpl) ListProvinces() ([]*common.Province, error) {
	resp, err := http.Get(ChinaPrefix)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf(resp.Status)
	}
	buf := bytes.NewBuffer(nil)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return nil, err
	}
	var provinces []*common.Province
	err = json.Unmarshal(buf.Bytes(), &provinces)
	if err != nil {
		return nil, err
	}
	return provinces, nil
}

func (g *GlClientImpl) ListCities(provinceCode int64) ([]*common.City, error) {
	resp, err := http.Get(ChinaPrefix + fmt.Sprintf("/%d", provinceCode))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf(resp.Status)
	}
	buf := bytes.NewBuffer(nil)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return nil, err
	}
	var cities []*common.City
	err = json.Unmarshal(buf.Bytes(), &cities)
	if err != nil {
		return nil, err
	}
	return cities, nil
}

func (g *GlClientImpl) ListCounties(provinceCode int64, cityCode int64) ([]*common.County, error) {
	resp, err := http.Get(ChinaPrefix + fmt.Sprintf("/%d/%d", provinceCode, cityCode))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf(resp.Status)
	}
	buf := bytes.NewBuffer(nil)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return nil, err
	}
	var counties []*common.County
	err = json.Unmarshal(buf.Bytes(), &counties)
	if err != nil {
		return nil, err
	}
	for _, c := range counties {
		c.WeatherID = strings.TrimPrefix(c.WeatherID, "CN")
	}
	return counties, nil
}

func (g *GlClientImpl) QueryWeather(locationCode int64) (*common.SwWeather3D, error) {
	resp, err := http.Get(WeatherPrefix + fmt.Sprintf("?cityid=CN%d&key=%s", locationCode, g.key))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf(resp.Status)
	}
	buf := bytes.NewBuffer(nil)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return nil, err
	}
	var swWeather3D common.SwWeather3D
	err = json.Unmarshal(buf.Bytes(), &swWeather3D)
	if err != nil {
		return nil, err
	}
	return &swWeather3D, nil
}
