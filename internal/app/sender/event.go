package sender

import (
	"github.com/ozonmp/ise-apartment-api/internal/model"
)

//EventSender comment for linter
type EventSender interface {
	Send(subdomain *apartment.ApartmentEvent) error
}
