package storage

import (
	"fmt"

	mt_domain "github.com/mitribu/mt-residentials/pkg/residential_units/domain"
)

type residentialInMemory struct {
	residentials map[mt_domain.ResidentialId]*mt_domain.Residential
}

func NewInMemoryRepository() mt_domain.ResidentialRepository {
	return &residentialInMemory{}
}

func (m residentialInMemory) Add(r *mt_domain.Residential) error {
	id := fmt.Sprintf("%s_%d", "res_", (len(m.residentials) + 1))
	r.Id = mt_domain.ResidentialId(id)
	m.residentials[r.Id] = r
	return nil
}

func (m residentialInMemory) List() ([]mt_domain.Residential, error) {
	var lr []mt_domain.Residential
	for k := range m.residentials {
		r := m.residentials[k]
		nr := mt_domain.Residential{
			Id:   r.Id,
			Name: r.Name,
		}
		nr.CreatedBy = r.CreatedBy
		nr.CreatedAt = r.CreatedAt
		nr.UpdatedBy = r.UpdatedBy
		nr.UpdatedAt = r.UpdatedAt

		lr = append(lr, nr)
	}
	return lr, nil
}
