package repository

import (
	"context"
	"echo_postgre/common"
	dtoReq "echo_postgre/dto/req"
	dtoResp "echo_postgre/dto/resp"
)

type IOrderRepository interface {
	Search(context.Context, string, *common.Paging) (*[]dtoResp.OrderDetailResponseDTO, error)
	GetById(context.Context, uint) (*dtoResp.OrderDetailResponseDTO, error)
	GetByUserId(context.Context, uint) (*dtoResp.OrderCartResponseDTO, error)
	Create(context.Context, dtoReq.CreateOrderDTO) error
	Update(context.Context, map[string]interface{}, dtoReq.UpdateOrderDTO) error
	Delete(context.Context, map[string]interface{}) error
}
