package handler

import "net/http"

/*
	type Handler interface {
		ServeHTTP(ResponseWriter, *Request)
	}
*/
type DemoHandler struct{}

func (h *DemoHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	_, _ = resp.Write(([]byte)(`{"HeWeather":[{"status":"Ok","msg":"","basic":{"city":"朝阳","id":"101010300","update":{"loc":"2020-11-28T19:35+08:00"}},"detail":{"city":{"sunrise":"07:15","sunset":"16:50","win_dir":"北风","win_deg":"1-2","pre":"1025","hum":"36"}},"now":{"tmp":"-1","cond":{"txt":"晴"}},"suggestion":{"comf":{"txt":"comfort"},"sport":{"txt":"sport"}},"daily_forecast":[{"tmp":{"max":"6","min":"-4"},"date":"2020-11-29","cond":{"txt":"多云"}},{"tmp":{"max":"4","min":"-4"},"date":"2020-11-30","cond":{"txt":"晴"}}]}]}`))
}
