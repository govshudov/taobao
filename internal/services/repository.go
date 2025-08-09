package services

import (
	"context"
	"taobao/internal/dto"
)

type ProductRepository interface {
	RecommendedProducts(ctx context.Context, pageNo, pageSize int, language string) (*dto.HotRecommendedResponse, error)
	ProductDetails(ctx context.Context, itemId int64) (*dto.ProductDetailsResponse, error)
}
