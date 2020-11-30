package clients

import "github.com/ame-lm/SimpleWeather/common"

type HeClient interface {
	GetHeWeather3d(location int64) (*common.HeWeather3D, error)
}

type GlClient interface {
	GetProvinces() ([]*common.Province, error)
	GetCities(provinceCode int64) ([]*common.City, error)
	GetCounties(provinceCode int64, cityCode int64) ([]*common.County, error)
}
