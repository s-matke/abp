package domain

type AvailabilityStore interface {
	GetAll() ([]*Availability, error)
	Insert(availability *Availability) error
	DeleteAll()
}
