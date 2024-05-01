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

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) repository.IProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (p *ProductRepository) Create(ctx context.Context, data dtoReq.CreateProductDTO) error {
	if err := p.db.Table(enum.PRODUCT_TABLE).Create(&data).Error; err != nil {
		return err
	}

	return nil
}

func (p *ProductRepository) Delete(ctx context.Context, cond map[string]interface{}) error {
	if err := p.db.Table(enum.PRODUCT_TABLE).Delete(&model.Users{}, cond).Error; err != nil {
		return err
	}

	return nil
}

func (p *ProductRepository) GetById(ctx context.Context, id uint) (*dtoResp.ProductResponseDTO, error) {
	dto := dtoResp.ProductResponseDTO{}

	if err := p.db.Table(enum.PRODUCT_TABLE).Where("id = ?", id).First(&dto).Error; err != nil {
		return nil, err
	}

	return &dto, nil
}

func (p *ProductRepository) Search(ctx context.Context, filter *common.Filter, paging *common.Paging) (*[]dtoResp.ProductResponseDTO, error) {
	var products []dtoResp.ProductResponseDTO

	if filter != nil {
		p.db.Table(enum.PRODUCT_TABLE).Where(filter)
	}

	if err := p.db.Table(enum.PRODUCT_TABLE).Find(&products).Error; err != nil {
		return nil, err
	}

	return &products, nil
}

func (p *ProductRepository) Update(ctx context.Context, cond map[string]interface{}, data dtoReq.UpdateProductDTO) error {
	panic("unimplemented")
}
