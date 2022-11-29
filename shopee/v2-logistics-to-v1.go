package shopee

type V1GetLogisticsListRs struct {
	RequestId string                `json:"request_id"`
	LogisticsChannelId []*V1GetLogisticsChannelIdListItem `json:"logistics_channel_list"`
}

type V1GetLogisticsChannelIdListItem struct {
	LogisticsChannelId string	`json:"logistics_channel_id"`
}


func TransformV2GetLogisticsListRsToV1 (r *V2Response[V2GetLogisticsListRs]) (*V1GetLogisticsListRs, error) {
	model := new(V1GetLogisticsListRs)
	model.RequestId = r.RequestId
	model.LogisticsChannelId = make([]*V1GetLogisticsChannelIdListItem, 0)
	
	return model, nil
}