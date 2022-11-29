package shopee

import (
	"net/http"
	"net/url"
	"strings"
)

type V2Order struct {
}

// /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// NewGetOrderList
// Use this api to search orders.
func (r *V2Order) NewGetOrderList() *V2GetOrderListRq {
	rq := &V2GetOrderListRq{}
	rq.method = http.MethodGet
	rq.path = "/api/v2/order/get_order_list"
	rq.commonKey = defaultCommonKey
	rq.parameter = url.Values{}
	return rq
}

type V2GetOrderListRq struct {
	ShopApiV2[V2GetOrderListRs]

	// The kind of time_from and time_to. Available value: create_time, update_time.
	TimeRangeField string
	// The time_from and time_to fields specify a date range for retrieving orders (based on the time_range_field).
	// The time_from field is the starting date range.
	// The maximum date range that may be specified with the time_from and time_to fields is 15 days.
	TimeFrom int64
	TimeTo   int64
	// Each result set is returned as a page of entries.
	// Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call).
	// This integer value is used to specify the maximum number of entries to return in a single "page" of data.
	// The limit of page_size if between 1 and 100.
	PageSize int64
	// Specifies the starting entry of data to return in the current call. Default is "".
	// If data is more than one page, the offset can be some entry to start next call.
	Cursor string
	// The order_status filter for retrieving orders and each one only every request.
	// Available value: UNPAID/READY_TO_SHIP/PROCESSED/SHIPPED/COMPLETED/IN_CANCEL/CANCELLED/INVOICE_PENDING
	OrderStatus string
	// Optional fields in response. Available value: order_status.
	ResponseOptionalFields string
}

type V2GetOrderListRs struct {
	More       bool   `json:"more"`
	NextCursor string `json:"next_cursor"`
	OrderList  []struct {
		OrderSn     string `json:"order_sn"`
		OrderStatus string `json:"order_status"`
	} `json:"order_list"`
}

func (r *V2GetOrderListRq) BindParameter() error {
	if r.TimeRangeField != "" {
		r.SetParameter("time_range_field", r.TimeRangeField)
	}
	if r.TimeFrom > 0 {
		r.SetParameter("time_from", r.TimeFrom)
	}
	if r.TimeTo > 0 {
		r.SetParameter("time_to", r.TimeTo)
	}
	if r.PageSize > 0 {
		r.SetParameter("page_size", r.PageSize)
	} else {
		r.SetParameter("page_size", "100")
	}
	if r.Cursor != "" {
		r.SetParameter("cursor", r.Cursor)
	}
	if r.OrderStatus != "" {
		r.SetParameter("order_status", r.OrderStatus)
	}
	if r.ResponseOptionalFields != "" {
		r.SetParameter("response_optional_fields", r.ResponseOptionalFields)
	}

	return nil
}

// /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// NewGetOrderDetail
// Use this api to search orders.
func (r *V2Order) NewGetOrderDetail() *V2GetOrderDetailRq {
	rq := &V2GetOrderDetailRq{}
	rq.method = http.MethodGet
	rq.path = "/api/v2/order/get_order_detail"
	rq.commonKey = defaultCommonKey
	rq.parameter = url.Values{}
	rq.OrderSnList = make([]string, 0)
	rq.ResponseOptionalFields = make([]string, 0)
	return rq
}

type V2GetOrderDetailRq struct {
	ShopApiV2[V2GetOrderDetailRs]

	// order_sn_list (Required)
	// The set of order_sn. If there are multiple order_sn, you need to use English comma to connect them. limit [1,50]
	OrderSnList []string
	// response_optional_fields
	ResponseOptionalFields []string
}

type V2GetOrderDetailRs struct {
	OrderList []struct {
		OrderSn              string  `json:"order_sn"`
		Region               string  `json:"region"`
		Currency             string  `json:"currency"`
		Cod                  bool    `json:"cod"`
		TotalAmount          float64 `json:"total_amount"`
		OrderStatus          string  `json:"order_status"`
		ShippingCarrier      string  `json:"shipping_carrier"`
		PaymentMethod        string  `json:"payment_method"`
		EstimatedShippingFee float64 `json:"estimated_shipping_fee"`
		MessageToSeller      string  `json:"message_to_seller"`
		CreateTime           int64   `json:"create_time"`
		UpdateTime           int64   `json:"update_time"`
		DaysToShip           int64   `json:"days_to_ship"`
		ShipByDate           int64   `json:"ship_by_date"`
		BuyerUserId          int64   `json:"buyer_user_id"`
		BuyerUsername        string  `json:"buyer_username"`
		RecipientAddress     struct {
			Name        string `json:"name"`
			Phone       string `json:"phone"`
			Town        string `json:"town"`
			District    string `json:"district"`
			City        string `json:"city"`
			State       string `json:"state"`
			Region      string `json:"region"`
			ZipCode     string `json:"zipcode"`
			FullAddress string `json:"full_address"`
		} `json:"recipient_address"`
		ActualShippingFee float64 `json:"actual_shipping_fee"`
		GoodsToDeclare    bool    `json:"goods_to_declare"`
		Note              string  `json:"note"`
		NoteUpdateTime    int64   `json:"note_update_time"`
		ItemList          []struct {
			ItemId                 int64   `json:"item_id"`
			ItemName               string  `json:"item_name"`
			ItemSku                string  `json:"item_sku"`
			ModelId                int64   `json:"model_id"`
			ModelName              string  `json:"model_name"`
			ModelSku               string  `json:"model_sku"`
			ModelQuantityPurchased int64   `json:"model_quantity_purchased"`
			ModelOriginalPrice     float64 `json:"model_original_price"`
			ModelDiscountedPrice   float64 `json:"model_discounted_price"`
			Wholesale              bool    `json:"wholesale"`
			Weight                 float64 `json:"weight"`
			AddOnDeal              bool    `json:"add_on_deal"`
			MainItem               bool    `json:"main_item"`
			AddOnDealId            int64   `json:"add_on_deal_id"`
			PromotionType          string  `json:"promotion_type"`
			PromotionId            int64   `json:"promotion_id"`
			OrderItemId            int64   `json:"order_item_id"`
			PromotionGroupId       int64   `json:"promotion_group_id"`
			ImageInfo              struct {
				ImageUrl string `json:"image_url"`
			} `json:"image_info"`
			ProductLocationId []string `json:"product_location_id"`
		} `json:"item_list"`
		PayTime                    int64  `json:"pay_time"`
		DropShipper                string `json:"dropshipper"`
		DropShipperPhone           string `json:"dropshipper_phone"`
		SplitUp                    bool   `json:"split_up"`
		BuyerCancelReason          string `json:"buyer_cancel_reason"`
		CancelBy                   string `json:"cancel_by"`
		CancelReason               string `json:"cancel_reason"`
		ActualShippingFeeConfirmed bool   `json:"actual_shipping_fee_confirmed"`
		BuyerCpfId                 string `json:"buyer_cpf_id"`
		FulfillmentFlag            string `json:"fulfillment_flag"`
		PickingDoneTime            int64  `json:"picking_done_time"`
		PackageList                []struct {
			PackageNumber   string `json:"package_number"`
			LogisticsStatus string `json:"logistics_status"`
			ShippingCarrier string `json:"shipping_carrier"`
			ItemList        []struct {
				ItemId        int64 `json:"item_id"`
				ModelId       int64 `json:"model_id"`
				ModelQuantity int64 `json:"model_quantity"`
			} `json:"item_list"`
			ParcelChargeableWeightGram int64 `json:"parcel_chargeable_weight_gram"`
		} `json:"package_list"`
		InvoiceData struct {
			Number             string  `json:"number"`
			SeriesNumber       string  `json:"series_number"`
			AccessKey          string  `json:"access_key"`
			IssueDate          int64   `json:"issue_date"`
			TotalValue         float64 `json:"total_value"`
			ProductsTotalValue float64 `json:"products_total_value"`
			TaxCode            string  `json:"tax_code"`
		} `json:"invoice_data"`
		CheckoutShippingCarrier   string   `json:"checkout_shipping_carrier"`
		ReverseShippingFee        float64  `json:"reverse_shipping_fee"`
		OrderChargeableWeightGram int64    `json:"order_chargeable_weight_gram"`
		EdtFrom                   int64    `json:"edt_from"`
		EdtTo                     int64    `json:"edt_to"`
		PrescriptionImage         []string `json:"prescription_images"`
		PrescriptionCheckStatus   int64    `json:"prescription_check_status"`
	} `json:"order_list"`
	Warning []string `json:"warning"`
}

func (r *V2GetOrderDetailRq) BindParameter() error {
	if len(r.OrderSnList) > 0 {
		r.parameter.Add("order_sn_list", strings.Join(r.OrderSnList, ","))
	}

	return nil
}

// /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
