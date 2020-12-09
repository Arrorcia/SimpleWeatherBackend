package main

import (
	"github.com/ame-lm/SimpleWeather/handler_v1"
	"github.com/ame-lm/SimpleWeather/handler_v2"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	gv1 := r.Group("/api/v1")
	gv1.GET("/china", handler_v1.ListProvincesHandler)
	gv1.GET("/china/:provinceCode", handler_v1.ListCitiesHandler)
	gv1.GET("/china/:provinceCode/:cityCode", handler_v1.ListCountiesHandler)
	gv1.GET("/weather", handler_v1.QueryWeatherHandler)

	gv2 := r.Group("/api/v2")
	gv2.GET("/china", handler_v2.ListProvincesHandler)
	gv2.GET("/china/:provinceCode", handler_v2.ListCitiesHandler)
	gv2.GET("/china/:provinceCode/:cityCode", handler_v2.ListCountiesHandler)
	gv2.GET("/weather", handler_v2.QueryWeatherHandler)

	_ = r.Run(":12345")
}
