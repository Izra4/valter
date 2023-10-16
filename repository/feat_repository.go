package repository

import (
	"Valter/db/sqlc"
	"context"
	"database/sql"
)

type FeatureRepository interface {
	FeatDummy() (sql.Result, error)
	FeatDummy2() (sql.Result, error)
	FeatDummy3() (sql.Result, error)
	FeatDummy4() (sql.Result, error)
	FeatDummy5() (sql.Result, error)
	FeatDummy6() (sql.Result, error)
	FeatDummy7() (sql.Result, error)
	FeatDummy8() (sql.Result, error)
	GetFeatures(productid uint32) ([]sqlc.Feature, error)
}

type featureRepository struct {
	db *sqlc.Queries
}

func NewFeatureRepository(db *sqlc.Queries) FeatureRepository {
	return &featureRepository{db}
}

func (f *featureRepository) FeatDummy() (sql.Result, error) {
	return f.db.FeatDummy(context.Background())
}

func (f *featureRepository) FeatDummy2() (sql.Result, error) {
	return f.db.FeatDummy2(context.Background())
}

func (f *featureRepository) FeatDummy3() (sql.Result, error) {
	return f.db.FeatDummy3(context.Background())
}

func (f *featureRepository) FeatDummy4() (sql.Result, error) {
	return f.db.FeatDummy4(context.Background())
}

func (f *featureRepository) FeatDummy5() (sql.Result, error) {
	return f.db.FeatDummy5(context.Background())
}

func (f *featureRepository) FeatDummy6() (sql.Result, error) {
	return f.db.FeatDummy6(context.Background())
}

func (f *featureRepository) FeatDummy7() (sql.Result, error) {
	return f.db.FeatDummy7(context.Background())
}

func (f *featureRepository) FeatDummy8() (sql.Result, error) {
	return f.db.FeatDummy8(context.Background())
}

func (f *featureRepository) GetFeatures(productid uint32) ([]sqlc.Feature, error) {
	data, err := f.db.GetFeatures(context.Background(), productid)
	if err != nil {
		return nil, err
	}
	return data, nil
}
