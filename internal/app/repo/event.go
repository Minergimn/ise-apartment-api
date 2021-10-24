package repo

import (
	"github.com/ozonmp/omp-demo-api/internal/model"
)

type EventRepo interface {
	Lock(n uint64) ([]apartment.ApartmentEvent, error)
	Unlock(eventIDs []uint64) error

	Add(event []apartment.ApartmentEvent) error
	Remove(eventIDs []uint64) error
}
