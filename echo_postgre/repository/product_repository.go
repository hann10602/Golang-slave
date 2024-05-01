package repository

import (
	"context"
	"echo_postgre/common"
	dtoRequest "echo_postgre/dto/req"
	dtoResponse "echo_postgre/dto/resp"
	"echo_postgre/model"
)

type IProductRepository interface {
	Search(context.Context, *common.Filter, *common.Paging) (*[]model.Products, error)
	GetById(context.Context, map[string]interface{}) (*dtoResponse.ProductsResponseDTO, error)
	Create(context.Context, dtoRequest.CreateProductDTO) (uint, error)
	Update(context.Context, map[string]interface{}, dtoRequest.UpdateProductDTO) error
	Delete(context.Context, map[string]interface{}) error
}
