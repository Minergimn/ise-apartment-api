package apartment

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

	Deferred EventStatus = iota
	Processed
)

type ApartmentCreated struct {
	ID     uint64
	Type   EventType
	Status EventStatus
	Entity *Apartment
}
