package controllers

import (
	"net/http"

	"github.com/PedroMartiniano/ecommerce-api-products/internal/application/services"
	"github.com/PedroMartiniano/ecommerce-api-products/internal/domain/dto"
	"github.com/PedroMartiniano/ecommerce-api-products/internal/domain/entities"
	"github.com/PedroMartiniano/ecommerce-api-products/internal/utils"
	"github.com/gin-gonic/gin"
)

type ProductsController struct {
	productsService *services.ProductsService
}

func NewProductsController(productsService *services.ProductsService) *ProductsController {
	return &ProductsController{
		productsService: productsService,
	}
}

func (p *ProductsController) CreateProductHandler(c *gin.Context) {
	var request createProductRequest

	if err := c.BindJSON(&request); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	product := entities.Product{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		CategoryID:  request.CategoryID,
		Quantity:    request.Quantity,
	}

	newProduct, err := p.productsService.CreateProductExecute(c.Request.Context(), product)
	if err != nil {
		code, message := httpError(err)
		sendError(c, code, message)
		return
	}

	sendSuccess(c, http.StatusCreated, newProduct)
}

func (p *ProductsController) ListProductsHandler(c *gin.Context) {
	products, err := p.productsService.ListProductsExecute(c.Request.Context())
	if err != nil {
		code, message := httpError(err)
		sendError(c, code, message)
		return
	}

	sendSuccess(c, http.StatusOK, products)
}

func (p *ProductsController) GetProductByIDHandler(c *gin.Context) {
	id := c.Param("id")

	product, err := p.productsService.GetProductByIDHandler(c.Request.Context(), id)
	if err != nil {
		code, message := httpError(err)
		sendError(c, code, message)
		return
	}

	sendSuccess(c, http.StatusOK, product)
}

func (p *ProductsController) UpdateProductHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		sendError(c, http.StatusBadRequest, "param 'id' is required")
		return
	}

	var request updateProductRequest

	if err := c.BindJSON(&request); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	product, err := p.productsService.GetProductByIDHandler(c.Request.Context(), id)
	if err != nil {
		code, message := httpError(err)
		sendError(c, code, message)
		return
	}

	utils.UpdateStrValues(&product.Name, request.Name)
	utils.UpdateStrValues(&product.Description, request.Description)
	utils.UpdateStrValues(&product.CategoryID, request.CategoryID)
	utils.UpdateFloatValues(&product.Price, request.Price)

	newProduct, err := p.productsService.UpdateProductHandler(c.Request.Context(), product)
	if err != nil {
		code, message := httpError(err)
		sendError(c, code, message)
		return
	}

	sendSuccess(c, http.StatusOK, newProduct)
}

func (p *ProductsController) DeleteProductHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		sendError(c, http.StatusBadRequest, "param 'id' is required")
		return
	}

	product, err := p.productsService.GetProductByIDHandler(c.Request.Context(), id)
	if err != nil {
		code, message := httpError(err)
		sendError(c, code, message)
		return
	}

	err = p.productsService.DeleteProductHandler(c.Request.Context(), product.ID)
	if err != nil {
		code, message := httpError(err)
		sendError(c, code, message)
		return
	}

	sendSuccess(c, http.StatusOK, "Product deleted successfully")
}

func (p *ProductsController) GetProductStockHandler(c *gin.Context) {
	id := c.Param("id")

	stock, err := p.productsService.GetProductStockHandler(c.Request.Context(), id)
	if err != nil {
		code, message := httpError(err)
		sendError(c, code, message)
		return
	}

	sendSuccess(c, http.StatusOK, stock)
}

func (p *ProductsController) UpdateProductStockHandler(c *gin.Context) {
	var request updateProductStockRequest

	if err := c.BindJSON(&request); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	productID := c.Param("id")
	if productID == "" {
		sendError(c, http.StatusBadRequest, "param 'id' is required")
		return
	}

	updateProductStock := dto.UpdateProductStock{
		ProductID: productID,
		Quantity:  request.Quantity,
		Operation: request.Operation,
	}

	stock, err := p.productsService.UpdateProductStockHandler(c.Request.Context(), updateProductStock)
	if err != nil {
		code, message := httpError(err)
		sendError(c, code, message)
		return
	}

	sendSuccess(c, http.StatusOK, stock)
}
