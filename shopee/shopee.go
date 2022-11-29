package shopee

type ApiV2 struct {
}

func NewApiV2() *ApiV2 {
	return &ApiV2{}
}

func (r *ApiV2) Order() *V2Order {
	return &V2Order{}
}

func (r *ApiV2) Payment() *V2Payment {
	return &V2Payment{}
}

func (r *ApiV2) Logistics() *V2Logistics {
	return &V2Logistics{}
}