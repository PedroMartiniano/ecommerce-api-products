package controllers

import (
	"net/http"

	"github.com/PedroMartiniano/ecommerce-api-products/internal/application/services"
	"github.com/PedroMartiniano/ecommerce-api-products/internal/domain/entities"
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

func (p *CategoriesController) CreateHandler(c *gin.Context) {
	var request createCategoryRequest

	if err := c.BindJSON(&request); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	category := entities.Category{
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

func (p *CategoriesController) ListHandler(c *gin.Context) {
	categories, err := p.categoriesService.ListCategoriesExecute(c.Request.Context())
	if err != nil {
		code, message := httpError(err)
		sendError(c, code, message)
		return
	}

	sendSuccess(c, http.StatusOK, categories)
}
