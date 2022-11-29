package controllers

import (
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
	"net/http"
	"shopeeadapterapi/models"
	"shopeeadapterapi/shopee"
	"strconv"
)

// /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type OrderListController struct {
	beego.Controller
}

type OrderListParam struct {
	CreateTimeForm           int64 `json:"create_time_from"`
	CreateTimeTo             int64 `json:"create_time_to"`
	UpdateTimeFrom           int64 `json:"update_time_from"`
	UpdateTimeTo             int64 `json:"update_time_to"`
	PaginationEntriesPerPage int64 `json:"pagination_entries_per_page"`
	PageOffset               int64 `json:"pagination_offset"`
	PartnerId                int64 `json:"partner_id"`
	ShopId                   int64 `json:"shopid"`
	Timestamp                int64 `json:"timestamp"`
}

// Post /api/v1/orders/basics
func (c *OrderListController) Post() {
	param := new(OrderListParam)
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &param); err != nil {
		c.CustomAbort(http.StatusBadRequest, models.NewErrorResponseString(models.ErrorParameter, err.Error()))
	}
	cParam, err := createCommonParameter(c.Ctx.Request.Header)
	if err != nil {
		c.CustomAbort(http.StatusBadRequest, models.NewErrorResponseString(models.ErrorParameter, err.Error()))
	}
	cParam.ShopId = param.ShopId
	cParam.PartnerId = param.PartnerId
	cParam.TimeStamp = param.Timestamp

	order := shopee.NewApiV2().Order()
	orderList := order.NewGetOrderList()
	// mapping parameter
	if param.CreateTimeForm > 0 && param.CreateTimeTo > 0 {
		orderList.TimeRangeField = "create_time"
		orderList.TimeFrom = param.CreateTimeForm
		orderList.TimeTo = param.CreateTimeTo
	}
	if param.UpdateTimeFrom > 0 && param.UpdateTimeTo > 0 {
		orderList.TimeRangeField = "update_time"
		orderList.TimeFrom = param.UpdateTimeFrom
		orderList.TimeTo = param.UpdateTimeTo
	}
	if param.PaginationEntriesPerPage > 0 {
		orderList.PageSize = param.PaginationEntriesPerPage
	}
	if param.PageOffset > 0 {
		orderList.Cursor = strconv.FormatInt(param.PageOffset, 10)
	}

	// request shopee order list v2
	err = orderList.BindParameter()
	if err != nil {
		c.CustomAbort(http.StatusBadRequest, models.NewErrorResponseString(models.ErrorParameter, err.Error()))
	}
	orderRs, err := orderList.Do(cParam)
	if err != nil {
		c.CustomAbort(http.StatusInternalServerError, models.NewErrorResponseString(models.ErrorInternal, err.Error()))
	}
	// transform to v1
	orderV1, err := shopee.TransformV2GetOrderListRsToV1(orderRs)
	if err != nil {
		c.CustomAbort(http.StatusInternalServerError, models.NewErrorResponseString(models.ErrorInternal, err.Error()))
	}
	_ = c.Ctx.Output.JSON(&orderV1, false, false)
}

// /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type OrderDetailController struct {
	beego.Controller
}

type OrderDetailParam struct {
	OrderSnList []string `json:"order_sn_list"`
	PartnerId   int      `json:"partner_id"`
	ShopId      int64    `json:"shopid"`
	Timestamp   int64    `json:"timestamp"`
}

// Post /api/v1/orders/detail
func (c *OrderDetailController) Post() {
	param := new(OrderDetailParam)
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &param); err != nil {
		c.CustomAbort(http.StatusBadRequest, models.NewErrorResponseString(models.ErrorParameter, err.Error()))
	}
	cParam, err := createCommonParameter(c.Ctx.Request.Header)
	if err != nil {
		c.CustomAbort(http.StatusBadRequest, models.NewErrorResponseString(models.ErrorParameter, err.Error()))
	}
	cParam.PartnerId = param.ShopId
	cParam.ShopId = param.ShopId
	cParam.TimeStamp = param.Timestamp

	orderDetail := shopee.NewApiV2().Order().NewGetOrderDetail()

	if len(param.OrderSnList) > 0 {
		orderDetail.OrderSnList = param.OrderSnList
	}

	err = orderDetail.BindParameter()
	if err != nil {
		c.CustomAbort(http.StatusBadRequest, models.NewErrorResponseString(models.ErrorParameter, err.Error()))
	}
	orderDetailRs, err := orderDetail.Do(cParam)
	if err != nil {
		c.CustomAbort(http.StatusInternalServerError, models.NewErrorResponseString(models.ErrorInternal, err.Error()))
	}

	_ = c.Ctx.Output.JSON(&orderDetailRs, false, false)
}

// /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
