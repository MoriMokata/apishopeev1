package controllers

import (
	"encoding/json"
	"net/http"

	beego "github.com/beego/beego/v2/server/web"
	"shopeeadapterapi/models"
	"shopeeadapterapi/shopee"
)

type LogisticsController struct{
	beego.Controller
}

type LogisticsParam struct {
	PartnerId	int64 `json:"partner_id"`
	ShopId		int64 `json:"shopid"`
	TimeStamp	int64 `json:"timestamp"`
}

// Post

func (c *LogisticsController) Post() {
	param := new(LogisticsParam)
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &param); err != nil {
		c.CustomAbort(http.StatusBadRequest, models.NewErrorResponseString(models.ErrorParameter, err.Error()))
	}

	cParam, err := createCommonParameter(c.Ctx.Request.Header)
	if err != nil {
		c.CustomAbort(http.StatusBadRequest, models.NewErrorResponseString(models.ErrorParameter, err.Error()))
	}
	cParam.PartnerId = param.PartnerId
	cParam.ShopId = param.ShopId
	cParam.TimeStamp = param.TimeStamp

	logistics := shopee.NewApiV2().Logistics()
	logisticsList := logistics.NewGetLogisticsList()
	//mapping parameter

	err = logisticsList.BindParameter()
	if err != nil {
		c.CustomAbort(http.StatusBadRequest, models.NewErrorResponseString(models.ErrorParameter, err.Error()))
	}
	
	logisticsRs, err := logisticsList.Do(cParam)
	if err != nil {
		c.CustomAbort(http.StatusBadRequest, models.NewErrorResponseString(models.ErrorParameter, err.Error()))
	}

	logisticsV1, err := shopee.TransformV2GetLogisticsListRsToV1(logisticsRs)
	if err != nil {
		c.CustomAbort(http.StatusInternalServerError, models.NewErrorResponseString(models.ErrorParameter, err.Error()))
	}
	_= c.Ctx.Output.JSON(&logisticsV1, false, false)
}