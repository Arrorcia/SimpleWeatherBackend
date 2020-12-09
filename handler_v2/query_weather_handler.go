package handler_v2

import (
	"github.com/ame-lm/SimpleWeather/clients"
	"github.com/ame-lm/SimpleWeather/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func QueryWeatherHandler(ctx *gin.Context) {
	heClient := clients.NewHeClientImpl("8e3add3c30cd4ab99bd66733fec5bbea")
	lC, err := strconv.Atoi(ctx.Query("location"))
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	heWeather3D, err := heClient.GetHeWeather3d(int64(lC))
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	swWeather3D, err := common.He2Sw(heWeather3D)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	swWeather3D.Weather[0].Basic.ID = ctx.Query("location")
	swWeather3D.Weather[0].Basic.City = "=city="
	ctx.JSON(http.StatusOK, swWeather3D)
}
