package services

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/PedroMartiniano/ecommerce-api-products/internal/configs"
	"github.com/PedroMartiniano/ecommerce-api-products/internal/domain/dto"
	"github.com/PedroMartiniano/ecommerce-api-products/internal/domain/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CategoriesServiceTestSuite struct {
	suite.Suite
	mockRepository *MockCategoriesRepository
	service        *CategoriesService
}

func (suite *CategoriesServiceTestSuite) SetupTest() {
	suite.mockRepository = new(MockCategoriesRepository)
	suite.service = NewCategoriesService(suite.mockRepository)
}

func (suite *CategoriesServiceTestSuite) TearDownTest() {
	suite.mockRepository.AssertExpectations(suite.T())
}

func (suite *CategoriesServiceTestSuite) TestCreateCategoryExecute_Success() {
	ctx := context.Background()
	inputDTO := dto.Category{
		Name:        "Electronics",
		Description: "Electronic products",
	}

	now := time.Now()
	expectedEntity, _ := entities.NewCategory("test-id", "Electronics", "Electronic products", &now, &now)

	suite.mockRepository.On("Create", ctx, mock.AnythingOfType("entities.Category")).Return(expectedEntity, nil)

	result, err := suite.service.CreateCategoryExecute(ctx, inputDTO)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "test-id", result.ID)
	assert.Equal(suite.T(), "Electronics", result.Name)
	assert.Equal(suite.T(), "Electronic products", result.Description)
	assert.Equal(suite.T(), now, result.CreatedAt)
	assert.Equal(suite.T(), now, result.UpdatedAt)
}

func (suite *CategoriesServiceTestSuite) TestCreateCategoryExecute_RepositoryError() {
	ctx := context.Background()
	inputDTO := dto.Category{
		Name:        "Electronics",
		Description: "Electronic products",
	}

	expectedError := configs.NewError(configs.ErrInternalServer, errors.New("RepositoryError"))
	suite.mockRepository.On("Create", ctx, mock.AnythingOfType("entities.Category")).Return(entities.Category{}, expectedError)

	result, err := suite.service.CreateCategoryExecute(ctx, inputDTO)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), expectedError, err)
	assert.Equal(suite.T(), dto.Category{}, result)
}

func (suite *CategoriesServiceTestSuite) TestListCategoriesExecute_Success() {
	ctx := context.Background()
	now := time.Now()

	category1, _ := entities.NewCategory("id1", "Electronics", "Electronic products", &now, &now)
	category2, _ := entities.NewCategory("id2", "Books", "Books and literature", &now, &now)

	expectedCategories := []entities.Category{category1, category2}
	suite.mockRepository.On("List", ctx).Return(expectedCategories, nil)

	result, err := suite.service.ListCategoriesExecute(ctx)

	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), result, 2)

	assert.Equal(suite.T(), "id1", result[0].ID)
	assert.Equal(suite.T(), "Electronics", result[0].Name)
	assert.Equal(suite.T(), "Electronic products", result[0].Description)

	assert.Equal(suite.T(), "id2", result[1].ID)
	assert.Equal(suite.T(), "Books", result[1].Name)
	assert.Equal(suite.T(), "Books and literature", result[1].Description)
}

func (suite *CategoriesServiceTestSuite) TestListCategoriesExecute_EmptyList() {
	ctx := context.Background()
	expectedCategories := []entities.Category{}
	suite.mockRepository.On("List", ctx).Return(expectedCategories, nil)

	result, err := suite.service.ListCategoriesExecute(ctx)

	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), result, 0)
}

func (suite *CategoriesServiceTestSuite) TestListCategoriesExecute_RepositoryError() {
	ctx := context.Background()
	expectedError := configs.NewError(configs.ErrInternalServer, errors.New("repository error"))
	suite.mockRepository.On("List", ctx).Return([]entities.Category{}, expectedError)

	result, err := suite.service.ListCategoriesExecute(ctx)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), expectedError, err)
	assert.Len(suite.T(), result, 0)
}

func (suite *CategoriesServiceTestSuite) TestNewCategoriesService_Success() {
	service := NewCategoriesService(suite.mockRepository)

	assert.NotNil(suite.T(), service)
	assert.Equal(suite.T(), suite.mockRepository, service.categoriesRepository)
}

func TestCategoriesServiceTestSuite(t *testing.T) {
	suite.Run(t, new(CategoriesServiceTestSuite))
}
