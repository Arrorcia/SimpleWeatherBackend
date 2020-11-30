package clients

import "github.com/ame-lm/SimpleWeather/common"

type HeClient interface {
	GetHeWeather3d(location int64) (*common.HeWeather3D, error)
}

type GlClient interface {
	ListProvinces() ([]*common.Province, error)
	ListCities(provinceCode int64) ([]*common.City, error)
	ListCounties(provinceCode int64, cityCode int64) ([]*common.County, error)
	QueryWeather(locationCode int64) (*common.SwWeather3D, error)
}
