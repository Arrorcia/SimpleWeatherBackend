package handler_v2

import (
	"github.com/ame-lm/SimpleWeather/handler_v1"
	"github.com/gin-gonic/gin"
)

func ListCountiesHandler(ctx *gin.Context) {
	handler_v1.ListCountiesHandler(ctx)
}
