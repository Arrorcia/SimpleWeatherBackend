package handler_v1

import (
	"github.com/ame-lm/SimpleWeather/clients"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ListCitiesHandler(ctx *gin.Context) {
	glClient := clients.NewGlClientImpl()
	pC, err := strconv.Atoi(ctx.Param("provinceCode"))
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	cities, err := glClient.GetCities(int64(pC))
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, cities)
}
