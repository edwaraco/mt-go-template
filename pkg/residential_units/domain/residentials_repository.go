package domain

type ResidentialRepository interface {
	Add(r *Residential) (*Residential, error)
	List() (*[]Residential, error)
}
