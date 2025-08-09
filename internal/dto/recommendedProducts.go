package dto

type RecommendedProductsRequest struct {
	PageNo   int    `json:"pageNo"`
	PageSize int    `json:"pageSize"`
	Language string `json:"language"`
}
type HotRecommendedProducts struct {
	MainImageUrl      string      `json:"mainImageUrl"`
	ItemId            int64       `json:"itemId"`
	MultiLanguageInfo interface{} `json:"multiLanguageInfo"`
	Price             string      `json:"price"`
	CouponPrice       string      `json:"couponPrice"`
	Quantity          int         `json:"quantity"`
	ShopName          string      `json:"shopName"`
	Title             string      `json:"title"`
	Tags              []string    `json:"tags"`
	PromotionDisplays interface{} `json:"promotionDisplays"`
}
type HotRecommendedResponse struct {
	RequestID string `json:"requestId"`
	Code      int    `json:"code"`
	DevelopID string `json:"developId"`
	Msg       string `json:"msg"`
	Data      struct {
		PageNo   int                      `json:"pageNo"`
		PageSize int                      `json:"pageSize"`
		Data     []HotRecommendedProducts `json:"data"`
	} `json:"data"`
}
