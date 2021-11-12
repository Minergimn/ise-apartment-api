package apartment

import "fmt"

type Apartment struct {
	ID     uint64 `db:"id"`
	Object string `db:"object"`
	Owner  string `db:"owner"`
	Status Status `db:"status"`
}

type Status uint8

const (
	Active Status = iota
	Deleted
)

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
	ID          uint64      `db:"id"`
	ApartmentId uint64      `db:"apartment_id"`
	Type        EventType   `db:"type"`
	Status      EventStatus `db:"status"`
	Entity      *Apartment  `db:"payload"`
	IsDeleted   bool        `db:"is_deleted"`
	IsLocked    bool        `db:"is_locked"`
	Updated     int64       `db:"updated"`
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
