package shopee

import (
	"net/http"
	"net/url"
)

type V2Payment struct {
}

// /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// NewMyIncome
// https://open.shopee.com/documents/v2/v2.payment.get_escrow_detail?module=97&type=1
func (r *V2Order) NewMyIncome() *V2MyIncomeRq {
	rq := &V2MyIncomeRq{}
	rq.method = http.MethodPost
	rq.path = "/api/v2/payment/get_escrow_detail"
	rq.commonKey = defaultCommonKey
	rq.parameter = url.Values{}
	return rq
}

type V2MyIncomeRq struct {
	ShopApiV2[V2MyIncomeRs]

	OrderSn string
}

type V2MyIncomeRs struct {
	OrderSn           string   `json:"order_sn"`
	BuyerUserName     string   `json:"buyer_user_name"`
	ReturnOrderSnList []string `json:"return_order_sn_list"`
	OrderIncome       struct {
		EscrowAmount               float64  `json:"escrow_amount"`
		BuyerTotalAmount           float64  `json:"buyer_total_amount"`
		OriginalPrice              float64  `json:"original_price"`
		SellerDiscount             float64  `json:"seller_discount"`
		ShopeeDiscount             float64  `json:"shopee_discount"`
		VoucherFromSeller          float64  `json:"voucher_from_seller"`
		VoucherFromShopee          float64  `json:"voucher_from_shopee"`
		Coins                      float64  `json:"coins"`
		BuyerPaidShippingFee       float64  `json:"buyer_paid_shipping_fee"`
		BuyerTransactionFee        float64  `json:"buyer_transaction_fee"`
		CrossBorderTax             float64  `json:"cross_border_tax"`
		PaymentPromotion           float64  `json:"payment_promotion"`
		CommissionFee              float64  `json:"commission_fee"`
		ServiceFee                 float64  `json:"service_fee"`
		SellerTransactionFee       float64  `json:"seller_transaction_fee"`
		SellerLostCompensation     float64  `json:"seller_lost_compensation"`
		SellerCoinCashBack         float64  `json:"seller_coin_cash_back"`
		EscrowTax                  float64  `json:"escrow_tax"`
		FinalShippingFee           float64  `json:"final_shipping_fee"`
		ActualShippingFee          float64  `json:"actual_shipping_fee"`
		OrderChargeableWeight      int64    `json:"order_chargeable_weight"`
		ShopeeShoppingRebate       float64  `json:"shopee_shipping_rebate"`
		ShippingFeeDiscountFrom3pl float64  `json:"shipping_fee_discount_from_3pl"`
		SellerShippingDiscount     float64  `json:"seller_shipping_discount"`
		EstimatedShippingFee       float64  `json:"estimated_shipping_fee"`
		SellerVoucherCode          []string `json:"seller_voucher_code"`
		DrcAdjustableRefund        float64  `json:"drc_adjustable_refund"`
		CostOfGoodsSold            float64  `json:"cost_of_goods_sold"`
		OriginalCostOfGoodsSold    float64  `json:"original_cost_of_goods_sold"`
		OriginalShopeeDiscount     float64  `json:"original_shopee_discount"`
		SellerReturnRefund         float64  `json:"seller_return_refund"`
		Items                      []struct {
			ItemId                    int64   `json:"item_id"`
			ItemName                  string  `json:"item_name"`
			ItemSku                   string  `json:"item_sku"`
			ModelId                   int64   `json:"model_id"`
			ModelName                 string  `json:"model_name"`
			ModelSku                  string  `json:"model_sku"`
			OriginalPrice             float64 `json:"original_price"`
			DiscountedPrice           float64 `json:"discounted_price"`
			SellerDiscount            float64 `json:"seller_discount"`
			ShopeeDiscount            float64 `json:"shopee_discount"`
			DiscountFromCoin          float64 `json:"discount_from_coin"`
			DiscountFromVoucherShopee float64 `json:"discount_from_voucher_shopee"`
			DiscountFromVoucherSeller float64 `json:"discount_from_voucher_seller"`
			ActivityType              string  `json:"activity_type"`
			ActivityId                int64   `json:"activity_id"`
			IsMainItem                bool    `json:"is_main_item"`
			QuantityPurchased         int64   `json:"quantity_purchased"`
			IsB2cShopItem             bool    `json:"is_b2c_shop_item"`
		} `json:"items"`
		EscrowAmountPri                     float64 `json:"escrow_amount_pri"`
		BuyerTotalAmountPri                 float64 `json:"buyer_total_amount_pri"`
		OriginalPricePri                    float64 `json:"original_price_pri"`
		SellerReturnRefundPri               float64 `json:"seller_return_refund_pri"`
		CommissionFeePri                    float64 `json:"commission_fee_pri"`
		ServiceFeePri                       float64 `json:"service_fee_pri"`
		DrcAdjustableRefundPri              float64 `json:"drc_adjustable_refund_pri"`
		PriCurrency                         string  `json:"pri_currency"`
		AffCurrency                         string  `json:"aff_currency"`
		ExchangeRate                        float64 `json:"exchange_rate"`
		ReverseShippingFee                  float64 `json:"reverse_shipping_fee"`
		FinalProductProtection              float64 `json:"final_product_protection"`
		CreditCardPromotion                 float64 `json:"credit_card_promotion"`
		CreditCardTransactionFee            float64 `json:"credit_card_transaction_fee"`
		FinalProductVatTex                  float64 `json:"final_product_vat_tax"`
		FinalShippingVatTax                 float64 `json:"final_shipping_vat_tax"`
		CampaignFee                         float64 `json:"campaign_fee"`
		SipSubsidy                          float64 `json:"sip_subsidy"`
		SipSubsidyPri                       float64 `json:"sip_subsidy_pri"`
		RefSellerProtectionFeeClaimAmount   float64 `json:"rsf_seller_protection_fee_claim_amount"`
		RefSellerProtectionFeePremiumAmount float64 `json:"rsf_seller_protection_fee_premium_amount"`
	} `json:"order_income"`
}
