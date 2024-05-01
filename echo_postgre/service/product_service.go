package service

import (
	"echo_postgre/common"
	dtoReq "echo_postgre/dto/req"
	dtoResp "echo_postgre/dto/resp"
	"echo_postgre/repository"

	"github.com/labstack/echo/v4"
)

type ProductService struct {
	productRepository repository.IProductRepository
}

type IProductService interface {
	HandleSearchProducts(echo.Context, *common.Filter, *common.Paging) (*[]dtoResp.ProductResponseDTO, error)
	HandleGetProductById(echo.Context, uint) (*dtoResp.ProductResponseDTO, error)
	HandleUpdateProduct(echo.Context, map[string]interface{}, dtoReq.UpdateProductDTO) error
	HandleDeleteProduct(echo.Context, map[string]interface{}) error
}

func NewProductService(productRepository repository.IProductRepository) IProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

func (p *ProductService) HandleDeleteProduct(ctx echo.Context, cond map[string]interface{}) error {
	if err := p.productRepository.Delete(ctx.Request().Context(), cond); err != nil {
		return err
	}

	return nil
}

func (p *ProductService) HandleGetProductById(ctx echo.Context, id uint) (*dtoResp.ProductResponseDTO, error) {
	productData, err := p.productRepository.GetById(ctx.Request().Context(), id)

	if err != nil {
		return nil, err
	}

	return productData, nil
}

func (p *ProductService) HandleSearchProducts(ctx echo.Context, filter *common.Filter, paging *common.Paging) (*[]dtoResp.ProductResponseDTO, error) {
	data, err := p.productRepository.Search(ctx.Request().Context(), filter, paging)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (p *ProductService) HandleUpdateProduct(ctx echo.Context, cond map[string]interface{}, data dtoReq.UpdateProductDTO) error {
	if err := p.productRepository.Update(ctx.Request().Context(), cond, data); err != nil {
		return err
	}

	return nil
}
