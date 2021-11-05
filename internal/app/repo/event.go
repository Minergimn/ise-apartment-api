package repo

import (
	"github.com/ozonmp/ise-apartment-api/internal/model"
)

type EventRepo interface {
	Lock(n uint64) ([]apartment.ApartmentEvent, error)
	Unlock(eventIDs []uint64) error

	Add(event []apartment.ApartmentEvent) error
	Remove(eventIDs []uint64) error
}
