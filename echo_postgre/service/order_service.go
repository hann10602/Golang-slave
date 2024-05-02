package service

import (
	"echo_postgre/common"
	dtoReq "echo_postgre/dto/req"
	dtoResp "echo_postgre/dto/resp"
	"echo_postgre/repository"

	"github.com/labstack/echo/v4"
)

type OrderService struct {
	orderRepository repository.IOrderRepository
}

type IOrderService interface {
	HandleCreateOrder(ctx echo.Context, data dtoReq.CreateOrderDTO) error
	HandleSearchOrders(echo.Context, *common.Filter, *common.Paging) (*[]dtoResp.OrderDetailResponseDTO, error)
	HandleGetOrderByUserId(echo.Context, uint) (*dtoResp.OrderDetailResponseDTO, error)
	HandleGetOrderById(echo.Context, uint) (*dtoResp.OrderCartResponseDTO, error)
	HandleUpdateOrder(echo.Context, map[string]interface{}, dtoReq.UpdateOrderDTO) error
	HandleDeleteOrder(echo.Context, map[string]interface{}) error
}

func NewOrderService(orderRepository repository.IOrderRepository) IOrderService {
	return &OrderService{
		orderRepository: orderRepository,
	}
}

func (o *OrderService) HandleCreateOrder(ctx echo.Context, data dtoReq.CreateOrderDTO) error {
	panic("unimplemented")
}

func (o *OrderService) HandleDeleteOrder(echo.Context, map[string]interface{}) error {
	panic("unimplemented")
}

func (o *OrderService) HandleGetOrderById(echo.Context, uint) (*dtoResp.OrderCartResponseDTO, error) {
	panic("unimplemented")
}

func (o *OrderService) HandleGetOrderByUserId(echo.Context, uint) (*dtoResp.OrderDetailResponseDTO, error) {
	panic("unimplemented")
}

func (o *OrderService) HandleSearchOrders(echo.Context, *common.Filter, *common.Paging) (*[]dtoResp.OrderDetailResponseDTO, error) {
	panic("unimplemented")
}

func (o *OrderService) HandleUpdateOrder(echo.Context, map[string]interface{}, dtoReq.UpdateOrderDTO) error {
	panic("unimplemented")
}
