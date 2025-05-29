package services

import (
	"context"

	"github.com/PedroMartiniano/ecommerce-api-products/internal/domain/entities"
	"github.com/stretchr/testify/mock"
)

// MockCategoriesRepository é o mock do repository de categorias
type MockCategoriesRepository struct {
	mock.Mock
}

func (m *MockCategoriesRepository) Create(ctx context.Context, category entities.Category) (entities.Category, error) {
	args := m.Called(ctx, category)
	return args.Get(0).(entities.Category), args.Error(1)
}

func (m *MockCategoriesRepository) FindById(ctx context.Context, id string) (entities.Category, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(entities.Category), args.Error(1)
}

func (m *MockCategoriesRepository) List(ctx context.Context) ([]entities.Category, error) {
	args := m.Called(ctx)
	return args.Get(0).([]entities.Category), args.Error(1)
}

func (m *MockCategoriesRepository) Update(ctx context.Context, product entities.Product) (entities.Product, error) {
	args := m.Called(ctx, product)
	return args.Get(0).(entities.Product), args.Error(1)
}

func (m *MockCategoriesRepository) Delete(ctx context.Context, product entities.Product) (entities.Product, error) {
	args := m.Called(ctx, product)
	return args.Get(0).(entities.Product), args.Error(1)
}
