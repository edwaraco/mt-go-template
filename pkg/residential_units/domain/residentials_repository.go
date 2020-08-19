package domain

type ResidentialRepository interface {
	Add(r *Residential) error
	List() ([]Residential, error)
}
