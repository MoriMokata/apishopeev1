package shopee

type V1GetOrderListRs struct {
	RequestId string                `json:"request_id"`
	More      bool                  `json:"more"`
	Order     []*V1GetOrderListItem `json:"orders"`
}

type V1GetOrderListItem struct {
	OrderSn     string `json:"ordersn"`
	OrderStatus string `json:"order_status"`
	UpdateTime  int64  `json:"update_time"`
}

func TransformV2GetOrderListRsToV1(r *V2Response[V2GetOrderListRs]) (*V1GetOrderListRs, error) {
	model := new(V1GetOrderListRs)
	model.RequestId = r.RequestId
	model.More = r.Response.More
	model.Order = make([]*V1GetOrderListItem, 0)

	for _, order := range r.Response.OrderList {
		model.Order = append(model.Order, &V1GetOrderListItem{
			OrderSn:     order.OrderSn,
			OrderStatus: order.OrderStatus,
		})
	}

	return model, nil
}
