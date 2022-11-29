package shopee

import (
	"net/http"
	"net/url"


)

type V2Logistics struct {
}

func (r *V2Logistics) NewGetLogisticsList() *V2GetLogisticsListRq {
	rq := &V2GetLogisticsListRq{}
	rq.method = http.MethodGet
	rq.path = "/api/v2/logistics/get_channel_list"
	rq.commonKey = defaultCommonKey
	rq.parameter = url.Values{}
	return rq
}

type V2GetLogisticsListRq struct {
	ShopApiV2[V2GetLogisticsListRs]
	PartnerId int64 `json:"partner_id"`
	TimeStamp int64	`json:"time_stamp"` 
	ShopId int64	`json:"shop_id"`
}

type V2GetLogisticsListRs struct {
	LogisticsChannelId []struct {
		LogisticsChannelId int	`json:"logistics_channel_id"`
		// Preferred bool 	`json:"preferred"`
		// LogisticsChannelName string 	`json:"logistics_channel_name"`
		// CodEnabled bool	`json:"cod_enabled"`
		// Enabled bool	`json:"enabled"`
		// FeeType string	`json:"fee_type"`
		// SizeList []struct {
		// 	SizeId string	`json:"size_id"`
		// 	Name string		`json:"name"`
		// 	DefaultPrice float64	`json:"default_price"`
		// }`json:"size_list"`
		// WeightLimit []struct {
		// 	ItemMaxWeight float64	`json:"item_max_weight"`
		// 	ItemMinWeight float64	`json:"item_min_weight"`
		// }`json:"weight_limit"`
		// ItemMaxDimension []struct {
		// 	Height float64	`json:"height"`
		// 	Width float64	`json:"width"`
		// 	Length float64 	`json:"length"`
		// 	Uint string		`json:"unit"`
		// 	DimensionSum float64	`json:"dimension_sum"`
		// }`json:"item_max_dimension"`
		// VolumeLimit []struct {
		// 	ItemMaxVolume float64	`json:"item_max_volume"`
		// 	ItemMinVolume float64	`json:"item_min_volume"`
		// }`json:"volume_limit"`
		// LogisticsDescription string	`json:"logistics_description"`
		// ForceEnable bool	`json:"force_enable"`
		// MaskChannelId int	`json:"mask_channel_id"`
	}`json:"logistics_channel_list"`
}


func (r *V2GetLogisticsListRq) BindParameter() error {
	if r.PartnerId > 0 {
		r.SetParameter("partner_id", r.PartnerId)
	}
	if r.TimeStamp > 0 {
		r.SetParameter("time_stamp", r.TimeStamp)
	}
	if r.ShopId > 0 {
		r.SetParameter("shop_id", r.ShopId)
	}
	return nil
}