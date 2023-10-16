package service

import (
	"Valter/db/sqlc"
	"Valter/repository"
	"database/sql"
)

type FeatService interface {
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

func NewFeatService(featureRepository repository.FeatureRepository) FeatService {
	return &featServ{featureRepository}
}

type featServ struct {
	featRepo repository.FeatureRepository
}

func (f *featServ) FeatDummy() (sql.Result, error) {
	return f.featRepo.FeatDummy()
}

func (f *featServ) FeatDummy2() (sql.Result, error) {
	return f.featRepo.FeatDummy2()
}

func (f *featServ) FeatDummy3() (sql.Result, error) {
	return f.featRepo.FeatDummy3()
}

func (f *featServ) FeatDummy4() (sql.Result, error) {
	return f.featRepo.FeatDummy4()
}

func (f *featServ) FeatDummy5() (sql.Result, error) {
	return f.featRepo.FeatDummy5()
}

func (f *featServ) FeatDummy6() (sql.Result, error) {
	return f.featRepo.FeatDummy6()
}

func (f *featServ) FeatDummy7() (sql.Result, error) {
	return f.featRepo.FeatDummy7()
}

func (f *featServ) FeatDummy8() (sql.Result, error) {
	return f.featRepo.FeatDummy8()
}

func (f *featServ) GetFeatures(productid uint32) ([]sqlc.Feature, error) {
	return f.featRepo.GetFeatures(productid)
}
