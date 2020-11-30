package main

import (
	"github.com/ame-lm/SimpleWeather/handler_v1"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	gv1 := r.Group("/api/v1")
	gv1.GET("/china", handler_v1.ListProvincesHandler)
	gv1.GET("/china/:provinceCode", handler_v1.ListCitiesHandler)
	gv1.GET("/china/:provinceCode/:cityCode", handler_v1.ListCountiesHandler)

	_ = r.Run(":12345")
}
