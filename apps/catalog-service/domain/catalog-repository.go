package domain

type CatalogRepository interface {
	Find(id string) (*Mower, error)
	Add(input CreateMowerDTO) (*Mower, error)
	Patch(id string, input UpdateMowerDTO) (*Mower, error)
	FindAvailableMowers() ([]*Mower, error)
}
