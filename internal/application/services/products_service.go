package services

import (
	"context"
	"errors"
	"time"

	pr "github.com/PedroMartiniano/ecommerce-api-products/internal/application/ports/repositories"
	"github.com/PedroMartiniano/ecommerce-api-products/internal/configs"
	"github.com/PedroMartiniano/ecommerce-api-products/internal/domain/dto"
	"github.com/PedroMartiniano/ecommerce-api-products/internal/domain/entities"
)

var logger = configs.GetLogger()

type ProductsService struct {
	productsRepository pr.IProductsRepository
	stocksRepository   pr.IStocksRepository
	redisRepository    pr.IRedisRepository
}

func NewProductsService(userRepository pr.IProductsRepository, stocksRepository pr.IStocksRepository, redisRepository pr.IRedisRepository) *ProductsService {
	return &ProductsService{
		productsRepository: userRepository,
		stocksRepository:   stocksRepository,
		redisRepository:    redisRepository,
	}
}

func (p *ProductsService) CreateProductExecute(c context.Context, product dto.Product) (dto.Product, error) {
	productEntity, err := entities.NewProduct("", product.Name, product.Description, product.CategoryID, product.Quantity, product.Price, nil, nil)
	if err != nil {
		return dto.Product{}, configs.NewError(configs.ErrBadRequest, err)
	}

	newProduct, err := p.productsRepository.Create(c, productEntity)
	if err != nil {
		return dto.Product{}, configs.NewError(configs.ErrBadRequest, err)
	}

	stock, err := entities.NewStock("", newProduct.ID.GetValue(), product.Quantity, nil)
	if err != nil {
		return dto.Product{}, configs.NewError(configs.ErrBadRequest, err)
	}
	_, err = p.stocksRepository.Create(c, stock)
	if err != nil {
		return dto.Product{}, configs.NewError(configs.ErrBadRequest, err)
	}

	p.redisRepository.Delete(c, "products:all")

	return dto.Product{
		ID:          newProduct.GetID(),
		Name:        newProduct.GetName(),
		Description: newProduct.GetDescription(),
		CategoryID:  newProduct.GetCategoryID(),
		Quantity:    newProduct.GetQuantity(),
		Price:       newProduct.GetPrice(),
		CreatedAt:   newProduct.GetCreatedAt(),
		UpdatedAt:   newProduct.GetUpdatedAt(),
	}, err
}

func (p *ProductsService) ListProductsExecute(c context.Context) ([]dto.Product, error) {
	var productsDTO []dto.Product

	cacheKey := "products:all"
	err := p.redisRepository.Get(c, cacheKey, &productsDTO)
	if err != nil {
		var myErr configs.Error
		if errors.As(err, &myErr) {
			if myErr.TypeError() != configs.ErrNotFound {
				return []dto.Product{}, err
			}
			// cache miss
		}
	}

	if len(productsDTO) > 0 { // cache hit
		return productsDTO, nil
	}

	products, err := p.productsRepository.List(c)

	for _, product := range products {
		productsDTO = append(productsDTO, dto.Product{
			ID:          product.GetID(),
			Name:        product.GetName(),
			Description: product.GetDescription(),
			CategoryID:  product.GetCategoryID(),
			Quantity:    product.GetQuantity(),
			Price:       product.GetPrice(),
			CreatedAt:   product.GetCreatedAt(),
			UpdatedAt:   product.GetUpdatedAt(),
		})
	}

	p.redisRepository.Set(c, cacheKey, productsDTO, 60*time.Minute)

	return productsDTO, err
}

func (p *ProductsService) GetProductByIDHandler(c context.Context, id string) (dto.Product, error) {
	var product dto.Product

	cacheKey := "products:" + id
	err := p.redisRepository.Get(c, cacheKey, &product)
	if err != nil {
		var myErr *configs.Error
		if errors.As(err, &myErr) {
			if myErr.TypeError() != configs.ErrNotFound {
				return dto.Product{}, err
			}
			// cache miss
		}
	}

	if product.ID != "" { // cache hit
		return product, nil
	}

	productEntity, err := p.productsRepository.FindById(c, id)
	if err != nil {
		return dto.Product{}, err
	}

	productDTO := dto.Product{
		ID:          productEntity.GetID(),
		Name:        productEntity.GetName(),
		Description: productEntity.GetDescription(),
		CategoryID:  productEntity.GetCategoryID(),
		Quantity:    productEntity.GetQuantity(),
		Price:       productEntity.GetPrice(),
		CreatedAt:   productEntity.GetCreatedAt(),
		UpdatedAt:   productEntity.GetUpdatedAt(),
	}

	p.redisRepository.Set(c, cacheKey, productDTO, 60*time.Minute)

	return productDTO, err
}

func (p *ProductsService) UpdateProductHandler(c context.Context, product dto.Product) (dto.Product, error) {
	productEntity, err := entities.NewProduct(product.ID, product.Name, product.Description, product.CategoryID, product.Quantity, product.Price, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return dto.Product{}, configs.NewError(configs.ErrBadRequest, err)
	}
	newProduct, err := p.productsRepository.Update(c, productEntity)

	p.redisRepository.Delete(c, "products:all")
	p.redisRepository.Delete(c, "products:"+product.ID)

	return dto.Product{
		ID:          newProduct.GetID(),
		Name:        newProduct.GetName(),
		Description: newProduct.GetDescription(),
		CategoryID:  newProduct.GetCategoryID(),
		Quantity:    newProduct.GetQuantity(),
		Price:       newProduct.GetPrice(),
		CreatedAt:   newProduct.GetCreatedAt(),
		UpdatedAt:   newProduct.GetUpdatedAt(),
	}, err
}

func (p *ProductsService) DeleteProductHandler(c context.Context, id string) error {
	err := p.stocksRepository.DeleteByProductID(c, id)
	if err != nil {
		return err
	}

	err = p.productsRepository.Delete(c, id)
	if err != nil {
		return err
	}

	p.redisRepository.Delete(c, "products:all")
	p.redisRepository.Delete(c, "products:"+id)

	return nil
}

func (p *ProductsService) GetProductStockHandler(c context.Context, id string) (dto.Stock, error) {
	stock, err := p.stocksRepository.GetByProductID(c, id)

	return dto.Stock{
		ID:        stock.GetID(),
		ProductID: stock.GetProductID(),
		Quantity:  stock.GetQuantity(),
		UpdatedAt: stock.GetUpdatedAt(),
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
		err := stock.Add(updateDTO.Quantity)
		if err != nil {
			return dto.Stock{}, configs.NewError(configs.ErrBadRequest, err)
		}
	} else if updateDTO.Operation == "remove" {
		err := stock.Subtract(updateDTO.Quantity)
		if err != nil {
			return dto.Stock{}, configs.NewError(configs.ErrBadRequest, err)
		}
	}

	updatedStock, err := p.stocksRepository.Update(c, stock)
	if err != nil {
		return dto.Stock{}, err
	}

	p.redisRepository.Delete(c, "products:all")
	p.redisRepository.Delete(c, "products:"+updateDTO.ProductID)

	return dto.Stock{
		ID:        updatedStock.GetID(),
		ProductID: updatedStock.GetProductID(),
		Quantity:  updatedStock.GetQuantity(),
		UpdatedAt: updatedStock.GetUpdatedAt(),
	}, err
}
