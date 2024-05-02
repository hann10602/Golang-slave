package implement

import (
	"context"
	"echo_postgre/common"
	dtoReq "echo_postgre/dto/req"
	dtoResp "echo_postgre/dto/resp"
	"echo_postgre/enum"
	"echo_postgre/model"
	"echo_postgre/repository"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) repository.IOrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (o *OrderRepository) Create(ctx context.Context, data dtoReq.CreateOrderDTO) error {
	if err := o.db.Table(enum.ORDER_TABLE).Create(&data).Error; err != nil {
		return err
	}

	return nil
}

func (o *OrderRepository) Delete(ctx context.Context, cond map[string]interface{}) error {
	if err := o.db.Table(enum.PRODUCT_TABLE).Delete(&model.Orders{}, cond).Error; err != nil {
		return err
	}

	return nil
}

func (o *OrderRepository) Search(ctx context.Context, search string, paging *common.Paging) (*[]dtoResp.OrderDetailResponseDTO, error) {
	panic("unimplemented")
}

func (o *OrderRepository) GetById(ctx context.Context, id uint) (*dtoResp.OrderDetailResponseDTO, error) {
	dto := dtoResp.OrderDetailResponseDTO{}

	if err := o.db.Table(enum.PRODUCT_TABLE).Where("id = ?", id).First(&dto).Error; err != nil {
		return nil, err
	}

	return &dto, nil
}

func (o *OrderRepository) GetByUserId(ctx context.Context, id uint) (*dtoResp.OrderCartResponseDTO, error) {
	panic("unimplemented")
}

func (o *OrderRepository) Update(ctx context.Context, cond map[string]interface{}, data dtoReq.UpdateOrderDTO) error {
	panic("unimplemented")
}
