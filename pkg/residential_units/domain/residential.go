package domain

import (
	domain "github.com/mitribu/mt-residentials/pkg/shared/domain"
)

type ResidentialId string

type WorkerId string

type Worker struct {
	Fullname string
	Role     string
}

type Residential struct {
	domain.MTModel
	Id      ResidentialId
	Name    string
	workers []*Worker
}

func NewResidential(name string, createdBy string) *Residential {
	residential := &Residential{Name: name}
	residential.CreatedBy = createdBy
	return residential
}

func NewWorker(fn string, r string) *Worker {
	return &Worker{Fullname: fn, Role: r}
}
