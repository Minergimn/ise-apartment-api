package apartment

import "fmt"

type Apartment struct {
	ID     uint64
	Object string
	Owner  string
}

type EventType uint8

type EventStatus uint8

const (
	Created EventType = iota
	Updated
	Removed
)

const (
	Deferred EventStatus = iota
	Processed
)

type ApartmentEvent struct {
	ID     uint64
	Type   EventType
	Status EventStatus
	Entity *Apartment
}
func (e *ApartmentEvent) String() string {
	return fmt.Sprintf("ApartmentEvent id: %d type: %s status: %s apartment id: %d", e.ID, e.Type, e.Status, e.Entity.ID)
}

func (d EventType) String() string {
	return [...]string{"Created", "Updated", "Removed"}[d]
}

func (d EventStatus) String() string {
	return [...]string{"Deferred", "Processed"}[d]
}
