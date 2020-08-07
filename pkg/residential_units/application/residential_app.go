package application

import (
	domain "github.com/mitribu/mt-residentials/pkg/residential_units/domain"
)

type ResidentialApp interface {
	Add(Name string) (*domain.Residential, error)
	AddWorker(id domain.ResidentialId, wn string, wr string) (*domain.Residential, error)
	List() (*[]domain.Residential, error)
}

type residentialAppImpl struct {
	ResidentialApp
	rr domain.ResidentialRepository
}

func newResidentialApp(r domain.ResidentialRepository) ResidentialApp {
	return &residentialAppImpl{rr: r}
}

func (r *residentialAppImpl) Add(Name string) (*domain.Residential, error) {
	return nil, nil
}

func (r *residentialAppImpl) AddWorker(id domain.ResidentialId, wn string, wr string) (*domain.Residential, error) {
	return nil, nil
}

func (r *residentialAppImpl) List() (*[]domain.Residential, error) {
	return nil, nil
}
