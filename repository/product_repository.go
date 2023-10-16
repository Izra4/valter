package repository

import (
	"Valter/db/sqlc"
	"context"
	"database/sql"
)

type ProductRepository interface {
	Dummy() (sql.Result, error)
	Dummy2() (sql.Result, error)
	GetAllProducts() ([]sqlc.Product, error)
	GetProductsById(id uint32) (*sqlc.Product, error)
	GetFeatures(productid uint32) ([]sqlc.Feature, error)
}

type productReposiotry struct {
	db *sqlc.Queries
}

func NewProductRepository(db *sqlc.Queries) ProductRepository {
	return &productReposiotry{db}
}

func (p *productReposiotry) Dummy() (sql.Result, error) {
	return p.db.Dummy(context.Background())
}

func (p *productReposiotry) Dummy2() (sql.Result, error) {
	return p.db.Dummy2(context.Background())
}

func (p *productReposiotry) GetAllProducts() ([]sqlc.Product, error) {
	data, err := p.db.GetAllProducts(context.Background())
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (p *productReposiotry) GetProductsById(id uint32) (*sqlc.Product, error) {
	data, err := p.db.GetProductsById(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (p *productReposiotry) GetFeatures(productid uint32) ([]sqlc.Feature, error) {
	data, err := p.db.GetFeatures(context.Background(), productid)
	if err != nil {
		return nil, err
	}
	return data, nil
}
