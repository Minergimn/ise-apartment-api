package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	totalApartmentNotFound = promauto.NewCounter(prometheus.CounterOpts{
		Name: "ise_apartment_api_apartment_not_found_total",
		Help: "Total number of apartments that were not found",
	})

	totalApartmentCUDEvents = promauto.NewCounterVec(prometheus.CounterOpts{
		Subsystem: "ise_apartment",
		Name:      "ise_apartment_api_apartment_cud_events_total",
		Help:      "Total number of apartments CUD events",
	}, []string{"event_type"})

	currentRetranslatorEventsCount = promauto.NewGauge(prometheus.GaugeOpts{
		Subsystem: "ise_apartment",
		Name:      "retranslator_events_count_total",
		Help:      "Number of processed events in the retranslator",
	})
)

type eventType uint8

const (
	//Created comment for linter
	Created eventType = eventType(iota)
	//Updated comment for linter
	Updated
	//Deleted comment for linter
	Deleted
)

//IncTotalApartmentNotFound comment for linter
func IncTotalApartmentNotFound() {
	totalApartmentNotFound.Inc()
}

//IncTotalApartmentCUDEvents comment for linter
func IncTotalApartmentCUDEvents(e eventType) {
	totalApartmentCUDEvents.WithLabelValues(e.String()).Inc()
}

//AddCurrentRetranslatorEventsCount comment for linter
func AddCurrentRetranslatorEventsCount(n float64) {
	currentRetranslatorEventsCount.Add(n)
}

//SubCurrentRetranslatorEventsCount comment for linter
func SubCurrentRetranslatorEventsCount(n float64) {
	currentRetranslatorEventsCount.Sub(n)
}

func (e eventType) String() string {
	return [...]string{"Created", "Updated", "Deleted"}[e]
}
