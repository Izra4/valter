package service

import (
	"Valter/db/sqlc"
	"Valter/repository"
	"Valter/utility"
	"database/sql"
	"github.com/gin-gonic/gin"
)

type ProductService interface {
	Dummy() (sql.Result, error)
	Dummy2() (sql.Result, error)
	GetAllProducts(c *gin.Context) ([]sqlc.Product, error)
	GetProductsById(c *gin.Context, id uint32) (*sqlc.Product, error)
	GetFeatures(productid uint32) ([]sqlc.Feature, error)
}

type productService struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) ProductService {
	return &productService{productRepository}
}

func (p *productService) Dummy() (sql.Result, error) {
	return p.productRepository.Dummy()
}

func (p *productService) Dummy2() (sql.Result, error) {
	return p.productRepository.Dummy2()
}

func (p *productService) GetAllProducts(c *gin.Context) ([]sqlc.Product, error) {
	data, err := p.productRepository.GetAllProducts()
	if err != nil {
		utility.HttpInternalErrorResponse(c, "Failed to get data", err)
		return nil, err
	}
	return data, nil
}

func (p *productService) GetProductsById(c *gin.Context, id uint32) (*sqlc.Product, error) {
	data, err := p.productRepository.GetProductsById(id)
	if err != nil {
		utility.HttpInternalErrorResponse(c, "Failed to get data", err)
		return nil, err
	}
	return data, nil
}

func (p *productService) GetFeatures(productid uint32) ([]sqlc.Feature, error) {
	return p.productRepository.GetFeatures(productid)
}
