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
	// 借助郭霖的数据来补全和风的数据（ City 、 ID ）
	glClient := clients.NewGlClientImpl()
	glWeather7D, err := glClient.QueryWeather(int64(lC))
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	tmpSw, err := common.Gl2Sw(glWeather7D)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	swWeather3D.Weather[0].Basic.ID = strconv.Itoa(lC)
	swWeather3D.Weather[0].Basic.City = tmpSw.Weather[0].Basic.City
	swWeather3D.Weather[0].Basic.Update.Loc = tmpSw.Weather[0].Basic.Update.Loc
	// 计算填写 Suggestion
	tmp, _ := strconv.Atoi(swWeather3D.Weather[0].Now.Temperature)
	comfSug, sportSug, err := common.GetSuggestions(tmp)
	swWeather3D.Weather[0].Suggestions.Comfort.Txt = comfSug
	swWeather3D.Weather[0].Suggestions.Sport.Txt = sportSug
	ctx.JSON(http.StatusOK, swWeather3D)
}
