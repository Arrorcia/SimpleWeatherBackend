package handler_v2

import (
	"github.com/ame-lm/SimpleWeather/handler_v1"
	"github.com/gin-gonic/gin"
)

func ListProvincesHandler(ctx *gin.Context) {
	handler_v1.ListProvincesHandler(ctx)
}
