package domain

type AvailabilityStore interface {
	GetAll() ([]*Availability, error)
	Insert(availability *Availability) error
	DeleteAll()
	GetByAccommodation(id string) ([]*Availability, error)
	GetAllUnavailable(availability *Availability) ([]*Availability, error)
}
