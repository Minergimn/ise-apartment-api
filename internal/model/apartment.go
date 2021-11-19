package apartment

import "fmt"

//Apartment comment for linter
type Apartment struct {
	ID     uint64 `db:"id"`
	Object string `db:"object"`
	Owner  string `db:"owner"`
	Status Status `db:"status"`
}

//Status comment for linter
type Status uint8

const (
	//Active comment for linter
	Active Status = iota
	//Deleted comment for linter
	Deleted
)

//EventType comment for linter
type EventType uint8

//EventStatus comment for linter
type EventStatus uint8

const (
	//Created comment for linter
	Created EventType = iota
	//Updated comment for linter
	Updated
	//Removed comment for linter
	Removed
)

const (
	//Deferred comment for linter
	Deferred EventStatus = iota
	//Processed comment for linter
	Processed
)

//ApartmentEvent comment for linter
type Event struct {
	ID          uint64      `db:"id"`
	ApartmentID uint64      `db:"apartment_id"`
	Type        EventType   `db:"type"`
	Status      EventStatus `db:"status"`
	Entity      *Apartment  `db:"payload"`
	IsDeleted   bool        `db:"is_deleted"`
	IsLocked    bool        `db:"is_locked"`
	Updated     int64       `db:"updated"`
}

func (e *Event) String() string {
	return fmt.Sprintf("ApartmentEvent id: %d type: %s status: %s apartment id: %d", e.ID, e.Type, e.Status, e.Entity.ID)
}

func (d EventType) String() string {
	return [...]string{"Created", "Updated", "Removed"}[d]
}

func (d EventStatus) String() string {
	return [...]string{"Deferred", "Processed"}[d]
}
