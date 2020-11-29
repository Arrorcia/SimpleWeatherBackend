package clients

import "github.com/ame-lm/SimpleWeather/common"

type HeClient interface {
	GetHeWeather3d(location int64) (*common.HeWeather3D, error)
}
