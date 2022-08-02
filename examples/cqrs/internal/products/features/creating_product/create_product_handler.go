package creatingProduct

import (
	"context"
	creatingProductDtos "github.com/mehdihadeli/mediatr/examples/cqrs/internal/products/features/creating_product/dtos"
	"github.com/mehdihadeli/mediatr/examples/cqrs/internal/products/models"
	"github.com/mehdihadeli/mediatr/examples/cqrs/internal/products/repository"
)

type CreateProductCommandHandler struct {
	productRepository *repository.InMemoryProductRepository
}

func NewCreateProductCommandHandler(productRepository *repository.InMemoryProductRepository) *CreateProductCommandHandler {
	return &CreateProductCommandHandler{productRepository: productRepository}
}

func (c *CreateProductCommandHandler) Handle(ctx context.Context, command *CreateProductCommand) (*creatingProductDtos.CreateProductCommandResponse, error) {

	product := &models.Product{
		ProductID:   command.ProductID,
		Name:        command.Name,
		Description: command.Description,
		Price:       command.Price,
		CreatedAt:   command.CreatedAt,
	}

	createdProduct, err := c.productRepository.CreateProduct(ctx, product)
	if err != nil {
		return nil, err
	}

	response := &creatingProductDtos.CreateProductCommandResponse{ProductID: createdProduct.ProductID}

	return response, nil
}
