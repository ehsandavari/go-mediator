package getting_product_by_id

import (
	uuid "github.com/satori/go.uuid"
)

type GetProductByIdQuery struct {
	ProductID uuid.UUID `validate:"required"`
}

func NewGetProductByIdQuery(productID uuid.UUID) *GetProductByIdQuery {
	return &GetProductByIdQuery{ProductID: productID}
}
