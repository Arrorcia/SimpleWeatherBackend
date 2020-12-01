package handler_v1

import (
	"github.com/ame-lm/SimpleWeather/clients"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func QueryWeatherHandler(ctx *gin.Context) {
	glClient := clients.NewGlClientImpl()
	lC, err := strconv.Atoi(ctx.Query("location"))
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	glWeather7D, err := glClient.QueryWeather(int64(lC))
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, glWeather7D)
}
