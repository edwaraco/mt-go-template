package application

import (
	mt_domain "github.com/mitribu/mt-residentials/pkg/residential_units/domain"
)

type ResidentialApp interface {
	Add(name string, createdBy string) (*mt_domain.Residential, error)
	AddWorker(id mt_domain.ResidentialId, wn string, wr string, updatedby string) (*mt_domain.Residential, error)
	List() (*[]mt_domain.Residential, error)
}

type residentialAppImpl struct {
	ResidentialApp
	rr mt_domain.ResidentialRepository
}

func NewResidentialApp(r mt_domain.ResidentialRepository) ResidentialApp {
	return &residentialAppImpl{rr: r}
}

func (r *residentialAppImpl) Add(name string, createdBy string) (*mt_domain.Residential, error) {
	residential := mt_domain.NewResidential(name, createdBy)
	return residential, nil
}

func (r *residentialAppImpl) AddWorker(id mt_domain.ResidentialId, wn string, wr string, updatedby string) (*mt_domain.Residential, error) {
	return nil, nil
}

func (r *residentialAppImpl) List() (*[]mt_domain.Residential, error) {
	return nil, nil
}
