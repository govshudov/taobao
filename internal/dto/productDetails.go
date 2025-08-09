package dto

type ProductDetailsResponse struct {
	RequestID string         `json:"requestId"`
	Code      int            `json:"code"`
	DevelopID string         `json:"developId"`
	Msg       string         `json:"msg"`
	Data      ProductDetails `json:"data"`
}

type ProductDetails struct {
	Quantity       int      `json:"quantity"`
	CategoryPath   string   `json:"categoryPath"`
	ShopName       string   `json:"shopName"`
	Description    string   `json:"description"`
	ProductUnit    string   `json:"productUnit"`
	PicUrls        []string `json:"picUrls"`
	Title          string   `json:"title"`
	MpID           string   `json:"mpId"`
	CategoryName   string   `json:"categoryName"`
	ItemID         int64    `json:"itemId"`
	UserNick       string   `json:"userNick"`
	Price          float64  `json:"price"`
	BeginAmount    int      `json:"beginAmount"`
	Status         string   `json:"status"`
	PromotionPrice float64  `json:"promotionPrice"`
	ShopID         int64    `json:"shopId"`
	CategoryID     string   `json:"categoryId"`
	ItemType       string   `json:"itemType"`
	ItemURL        string   `json:"itemUrl"`
	SkuList        []SKU    `json:"skuList"`
}

type SKU struct {
	PicURL         string     `json:"picUrl"`
	Quantity       int        `json:"quantity"`
	Price          float64    `json:"price"`
	SkuID          string     `json:"skuId"`
	Status         string     `json:"status"`
	PromotionPrice string     `json:"promotionPrice"`
	PostFee        float64    `json:"postFee"`
	MpSkuID        int64      `json:"mpSkuId"`
	Properties     []Property `json:"properties"`
}

type Property struct {
	ValueID   int    `json:"valueId"`
	ValueName string `json:"valueName"`
	PropID    int    `json:"propId"`
	PropName  string `json:"propName"`
}
