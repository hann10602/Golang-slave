package repository

import (
	"context"
	"echo_postgre/common"
	dtoReq "echo_postgre/dto/req"
	dtoResp "echo_postgre/dto/resp"
)

type IProductRepository interface {
	Search(context.Context, *common.Filter, *common.Paging) (*[]dtoResp.ProductResponseDTO, error)
	GetById(context.Context, uint) (*dtoResp.ProductResponseDTO, error)
	Create(context.Context, dtoReq.CreateProductDTO) error
	Update(context.Context, map[string]interface{}, dtoReq.UpdateProductDTO) error
	Delete(context.Context, map[string]interface{}) error
}
