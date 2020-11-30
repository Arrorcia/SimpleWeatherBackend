package handler_v1

import (
	"github.com/ame-lm/SimpleWeather/clients"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListProvincesHandler(ctx *gin.Context) {
	glClient := clients.NewGlClientImpl()
	provinces, err := glClient.GetProvinces()
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, provinces)
}
