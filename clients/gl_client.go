package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ame-lm/SimpleWeather/common"
	"net/http"
	"strings"
)

const PREFIX = "http://guolin.tech/api/china"

type GlClientImpl struct{}

func NewGlClientImpl() *GlClientImpl {
	return &GlClientImpl{}
}

func (g *GlClientImpl) GetProvinces() ([]*common.Province, error) {
	resp, err := http.Get(PREFIX)
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

func (g *GlClientImpl) GetCities(provinceCode int64) ([]*common.City, error) {
	resp, err := http.Get(PREFIX + fmt.Sprintf("/%d", provinceCode))
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

func (g *GlClientImpl) GetCounties(provinceCode int64, cityCode int64) ([]*common.County, error) {
	resp, err := http.Get(PREFIX + fmt.Sprintf("/%d/%d", provinceCode, cityCode))
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
