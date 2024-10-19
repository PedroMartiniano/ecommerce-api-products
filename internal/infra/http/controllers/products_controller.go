package controllers

import (
	"net/http"

	"github.com/PedroMartiniano/ecommerce-api-products/internal/application/services"
	"github.com/PedroMartiniano/ecommerce-api-products/internal/domain/dto"
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

// @BasePath /products
// @Summary Create an product
// @Security BearerAuth
// @Tags Products
// @Accept json
// @Produce json
// @Param request body createProductRequest true "Request Body"
// @Success 201 {object} productResponse1
// @Failure 401 {object} errorResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /products/ [post]
func (p *ProductsController) CreateProductHandler(c *gin.Context) {
	var request createProductRequest

	if err := c.BindJSON(&request); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	product := dto.Product{
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

// @BasePath /products
// @Summary List all products
// @Tags Products
// @Accept json
// @Produce json
// @Success 200 {object} productResponse1
// @Failure 401 {object} errorResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /products/ [get]
func (p *ProductsController) ListProductsHandler(c *gin.Context) {
	products, err := p.productsService.ListProductsExecute(c.Request.Context())
	if err != nil {
		code, message := httpError(err)
		sendError(c, code, message)
		return
	}

	sendSuccess(c, http.StatusOK, products)
}

// @BasePath /products
// @Summary Get a product by ID
// @Tags Products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} productResponse2
// @Failure 401 {object} errorResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /products/{id} [get]
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

// @BasePath /products
// @Summary update a product by ID
// @Security BearerAuth
// @Tags Products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param request body updateProductRequest true "Request Body"
// @Success 200 {object} productResponse1
// @Failure 401 {object} errorResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /products/{id} [put]
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

// @BasePath /products
// @Summary Delete a product by ID
// @Security BearerAuth
// @Tags Products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} productResponse3
// @Failure 401 {object} errorResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /products/{id} [delete]
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

// @BasePath /products
// @Summary Get a product stock by productID
// @Tags Stocks
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} stockResponse1
// @Failure 401 {object} errorResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /products/{id}/stock [get]
func (p *ProductsController) GetProductStockHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		sendError(c, http.StatusBadRequest, "param 'id' is required")
	}

	stock, err := p.productsService.GetProductStockHandler(c.Request.Context(), id)
	if err != nil {
		code, message := httpError(err)
		sendError(c, code, message)
		return
	}

	sendSuccess(c, http.StatusOK, stock)
}

// @BasePath /products
// @Summary Update a product stock
// @Tags Stocks
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param request body updateProductStockRequest true "Product ID"
// @Success 200 {object} stockResponse1
// @Failure 401 {object} errorResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /products/{id}/stock [put]
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
