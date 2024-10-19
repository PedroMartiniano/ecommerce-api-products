package controllers

import (
	"net/http"

	"github.com/PedroMartiniano/ecommerce-api-products/internal/application/services"
	"github.com/PedroMartiniano/ecommerce-api-products/internal/domain/dto"
	"github.com/gin-gonic/gin"
)

type CategoriesController struct {
	categoriesService *services.CategoriesService
}

func NewCategoriesController(categoriesService *services.CategoriesService) *CategoriesController {
	return &CategoriesController{
		categoriesService: categoriesService,
	}
}

// @BasePath /categories
// @Summary Create an product category
// @Security BearerAuth
// @Tags Categories
// @Accept json
// @Produce json
// @Param request body createCategoryRequest true "Request Body"
// @Success 201 {object} categoryResponse1
// @Failure 401 {object} errorResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /categories/ [post]
func (p *CategoriesController) CreateHandler(c *gin.Context) {
	var request createCategoryRequest

	if err := c.BindJSON(&request); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	category := dto.Category{
		Name:        request.Name,
		Description: request.Description,
	}

	newCategory, err := p.categoriesService.CreateCategoryExecute(c.Request.Context(), category)
	if err != nil {
		code, message := httpError(err)
		sendError(c, code, message)
		return
	}

	sendSuccess(c, http.StatusCreated, newCategory)
}

// @BasePath /categories
// @Summary List all Products Categories
// @Tags Categories
// @Accept json
// @Produce json
// @Success 200 {object} categoryResponse2
// @Failure 401 {object} errorResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /categories/ [get]
func (p *CategoriesController) ListHandler(c *gin.Context) {
	categories, err := p.categoriesService.ListCategoriesExecute(c.Request.Context())
	if err != nil {
		code, message := httpError(err)
		sendError(c, code, message)
		return
	}

	sendSuccess(c, http.StatusOK, categories)
}
