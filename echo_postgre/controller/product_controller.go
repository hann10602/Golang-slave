package controller

import (
	"echo_postgre/common"
	dtoReq "echo_postgre/dto/req"
	"echo_postgre/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	productService service.IProductService
}

func NewProductController(productService service.IProductService) ProductController {
	return ProductController{
		productService: productService,
	}
}

func (u ProductController) DeleteProduct(ctx echo.Context) error {
	productId := ctx.Param("id")

	if productId == "" {
		return ctx.JSON(http.StatusBadRequest, &common.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Cannot received product id",
			Data:       false,
		})
	}

	if err := u.productService.HandleDeleteProduct(ctx, map[string]interface{}{
		"id": productId,
	}); err != nil {
		return ctx.JSON(http.StatusBadRequest, &common.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       false,
		})
	}

	return ctx.JSON(http.StatusOK, &common.Response{
		StatusCode: http.StatusOK,
		Message:    "Deleted product successfully",
		Data:       true,
	})
}

func (u ProductController) GetProductById(ctx echo.Context) error {
	productId, err := strconv.ParseUint(ctx.Param("id"), 10, 32)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &common.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Cannot received product id",
			Data:       false,
		})
	}

	product, err := u.productService.HandleGetProductById(ctx, uint(productId))

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &common.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       false,
		})
	}

	return ctx.JSON(http.StatusOK, &common.Response{
		StatusCode: http.StatusOK,
		Message:    "Get product successfully",
		Data:       product,
	})
}

func (u ProductController) SearchProduct(ctx echo.Context) error {
	var paging common.Paging
	var filter common.Filter

	err := ctx.Bind(&paging)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &common.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       false,
		})
	}

	err = ctx.Bind(&filter)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &common.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       false,
		})
	}

	paging.Process()

	data, err := u.productService.HandleSearchProducts(ctx, &filter, &paging)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &common.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       false,
		})
	}

	return ctx.JSON(http.StatusOK, &common.Response{
		StatusCode: http.StatusOK,
		Message:    "Search products successfully",
		Data:       common.HandleResponseWithPagination(data, paging),
	})
}

func (u ProductController) UpdateProduct(ctx echo.Context) error {
	var product dtoReq.UpdateProductDTO

	if err := ctx.Bind(&product); err != nil {
		return ctx.JSON(http.StatusBadRequest, &common.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       false,
		})
	}

	productId := ctx.Param("id")

	if productId == "" {
		return ctx.JSON(http.StatusBadRequest, &common.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Cannot received product id",
			Data:       false,
		})
	}

	err := u.productService.HandleUpdateProduct(ctx, map[string]interface{}{
		"id": productId,
	}, product)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &common.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       false,
		})
	}

	return ctx.JSON(http.StatusOK, &common.Response{
		StatusCode: http.StatusOK,
		Message:    "Updated product successfully",
		Data:       true,
	})
}
