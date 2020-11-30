package handler_v1

import (
	"github.com/ame-lm/SimpleWeather/clients"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ListCountiesHandler(ctx *gin.Context) {
	glClient := clients.NewGlClientImpl()
	pC, err := strconv.Atoi(ctx.Param("provinceCode"))
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	cC, err := strconv.Atoi(ctx.Param("cityCode"))
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	counties, err := glClient.ListCounties(int64(pC), int64(cC))
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, counties)
}
