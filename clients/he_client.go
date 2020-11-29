package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ame-lm/SimpleWeather/common"
	"net/http"
)

type HeClientImpl struct {
	key string
}

func NewHeClientImpl(key string) *HeClientImpl {
	return &HeClientImpl{key: key}
}

func (h *HeClientImpl) GetHeWeather3d(location int64) (*common.HeWeather3D, error) {
	if h == nil {
		return nil, fmt.Errorf("client is empty")
	}
	if h.key == "" {
		return nil, fmt.Errorf("client's key is empty")
	}
	resp, err := http.Get(fmt.Sprintf("https://devapi.qweather.com/v7/weather/3d?location=%d&key=%s", location, h.key))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf(resp.Status)
	}
	buf := bytes.NewBuffer(nil)
	if _, err := buf.ReadFrom(resp.Body); err != nil {
		return nil, err
	}
	hw := common.NewHeWeather3D()
	if err = json.Unmarshal(([]byte)(buf.String()), hw); err != nil {
		return nil, err
	}
	return hw, nil
}
