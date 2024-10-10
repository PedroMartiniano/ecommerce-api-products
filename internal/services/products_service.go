package services

import (
	"context"

	"github.com/PedroMartiniano/ecommerce-api-products/internal/configs"
	"github.com/PedroMartiniano/ecommerce-api-products/internal/models"
	pr "github.com/PedroMartiniano/ecommerce-api-products/internal/ports/irepositories"
)

var logger = configs.GetLogger()

type ProductsService struct {
	productsRepository pr.IProductsRepository
	stocksRepository   pr.IStocksRepository
}

func NewProductsService(userRepository pr.IProductsRepository, stocksRepository pr.IStocksRepository) *ProductsService {
	return &ProductsService{
		productsRepository: userRepository,
		stocksRepository:   stocksRepository,
	}
}

func (p *ProductsService) CreateProductExecute(c context.Context, product models.Products) (models.Products, error) {
	newProduct, err := p.productsRepository.Create(c, product)

	p.stocksRepository.Create(c, models.Stocks{
		ProductID: newProduct.ID,
		Quantity:  product.Quantity,
	})

	return newProduct, err
}

func (p *ProductsService) ListProductsExecute(c context.Context) ([]models.Products, error) {
	products, err := p.productsRepository.List(c)

	return products, err
}

func (p *ProductsService) GetProductByIDHandler(c context.Context, id string) (models.Products, error) {
	product, err := p.productsRepository.FindById(c, id)

	return product, err
}

func (p *ProductsService) UpdateProductHandler(c context.Context, product models.Products) (models.Products, error) {
	newProduct, err := p.productsRepository.Update(c, product)

	return newProduct, err
}

func (p *ProductsService) DeleteProductHandler(c context.Context, id string) error {
	err := p.productsRepository.Delete(c, id)

	return err
}
