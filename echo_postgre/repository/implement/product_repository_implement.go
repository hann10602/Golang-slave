package implement

import (
	"context"
	"echo_postgre/common"
	dtoRequest "echo_postgre/dto/req"
	dtoResponse "echo_postgre/dto/resp"
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

// Create implements repository.IProductRepository.
func (p *ProductRepository) Create(ctx context.Context, data dtoRequest.CreateProductDTO) (uint, error) {
	panic("unimplemented")
}

// Delete implements repository.IProductRepository.
func (p *ProductRepository) Delete(ctx context.Context, cond map[string]interface{}) error {
	panic("unimplemented")
}

// GetById implements repository.IProductRepository.
func (p *ProductRepository) GetById(ctx context.Context, cond map[string]interface{}) (*dtoResponse.ProductsResponseDTO, error) {
	panic("unimplemented")
}

// Search implements repository.IProductRepository.
func (p *ProductRepository) Search(ctx context.Context, filter *common.Filter, paging *common.Paging) (*[]model.Products, error) {
	var products []model.Products

	if filter != nil {
		p.db.Table(enum.PRODUCT_TABLE).Where(filter)
	}

	if err := p.db.Table(enum.PRODUCT_TABLE).Find(&products).Error; err != nil {
		return nil, err
	}

	return &products, nil
}

// Update implements repository.IProductRepository.
func (p *ProductRepository) Update(ctx context.Context, cond map[string]interface{}, data dtoRequest.UpdateProductDTO) error {
	panic("unimplemented")
}
