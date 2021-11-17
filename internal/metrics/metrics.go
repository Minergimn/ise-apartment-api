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
	Created = eventType(iota)
	Updated
	Deleted
)

func IncTotalApartmentNotFound() {
	totalApartmentNotFound.Inc()
}

func IncTotalApartmentCUDEvents(e eventType) {
	totalApartmentCUDEvents.WithLabelValues(e.String()).Inc()
}

func AddCurrentRetranslatorEventsCount(n float64) {
	currentRetranslatorEventsCount.Add(n)
}

func SubCurrentRetranslatorEventsCount(n float64) {
	currentRetranslatorEventsCount.Sub(n)
}

func (e eventType) String() string {
	return [...]string{"Created", "Updated", "Deleted"}[e]
}
