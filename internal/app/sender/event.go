package sender

import (
	"github.com/ozonmp/ise-apartment-api/internal/model"
)

type EventSender interface {
	Send(subdomain *apartment.ApartmentEvent) error
}
