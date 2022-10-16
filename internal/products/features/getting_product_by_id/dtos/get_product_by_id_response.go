package dtos

import (
	"github.com/danmaxdanilov/zts.gateway/internal/products/dtos"
)

type GetProductByIdQueryResponse struct {
	Product *dtos.ProductDto `json:"product"`
}
