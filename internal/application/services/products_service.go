package services

import (
	"context"
	"errors"
	"time"

	pr "github.com/PedroMartiniano/ecommerce-api-products/internal/application/ports"
	"github.com/PedroMartiniano/ecommerce-api-products/internal/configs"
	"github.com/PedroMartiniano/ecommerce-api-products/internal/domain/dto"
	"github.com/PedroMartiniano/ecommerce-api-products/internal/domain/entities"
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

func (p *ProductsService) CreateProductExecute(c context.Context, product entities.Product) (entities.Product, error) {
	newProduct, err := p.productsRepository.Create(c, product)
	if err != nil {
		return entities.Product{}, configs.NewError(configs.ErrBadRequest, err)
	}

	stock, err := entities.NewStock("", newProduct.ID, product.Quantity, time.Now())
	if err != nil {
		return entities.Product{}, configs.NewError(configs.ErrBadRequest, err)
	}
	p.stocksRepository.Create(c, stock)

	return newProduct, err
}

func (p *ProductsService) ListProductsExecute(c context.Context) ([]entities.Product, error) {
	products, err := p.productsRepository.List(c)

	return products, err
}

func (p *ProductsService) GetProductByIDHandler(c context.Context, id string) (entities.Product, error) {
	product, err := p.productsRepository.FindById(c, id)

	return product, err
}

func (p *ProductsService) UpdateProductHandler(c context.Context, product entities.Product) (entities.Product, error) {
	newProduct, err := p.productsRepository.Update(c, product)

	return newProduct, err
}

func (p *ProductsService) DeleteProductHandler(c context.Context, id string) error {
	err := p.stocksRepository.DeleteByProductID(c, id)
	if err != nil {
		return err
	}

	err = p.productsRepository.Delete(c, id)

	return err
}

func (p *ProductsService) GetProductStockHandler(c context.Context, id string) (dto.Stock, error) {
	stock, err := p.stocksRepository.GetByProductID(c, id)

	return dto.Stock{
		ID:        stock.ID.GetValue(),
		ProductID: stock.ProductID.GetValue(),
		Quantity:  stock.Quantity.GetValue(),
		UpdatedAt: stock.UpdatedAt,
	}, err
}

func (p *ProductsService) UpdateProductStockHandler(c context.Context, updateDTO dto.UpdateProductStock) (dto.Stock, error) {
	stock, err := p.stocksRepository.GetByProductID(c, updateDTO.ProductID)
	if err != nil {
		return dto.Stock{}, err
	}

	if updateDTO.Operation != "add" && updateDTO.Operation != "remove" {
		return dto.Stock{}, configs.NewError(configs.ErrBadRequest, errors.New("operation must be 'add' or 'remove'"))
	}

	if updateDTO.Operation == "add" {
		stock.Add(updateDTO.Quantity)
	} else if updateDTO.Operation == "remove" {
		stock.Subtract(updateDTO.Quantity)
	}

	updatedStock, err := p.stocksRepository.Update(c, stock)
	if err != nil {
		return dto.Stock{}, err
	}

	return dto.Stock{
		ID:        updatedStock.ID.GetValue(),
		ProductID: updatedStock.ProductID.GetValue(),
		Quantity:  updatedStock.Quantity.GetValue(),
		UpdatedAt: updatedStock.UpdatedAt,
	}, err
}
