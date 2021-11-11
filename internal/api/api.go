package api

import (
	event "github.com/ozonmp/ise-apartment-api/internal/app/repo"
	"github.com/ozonmp/ise-apartment-api/internal/repo"
	pb "github.com/ozonmp/ise-apartment-api/pkg/ise-apartment-api"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	totalApartmentNotFound = promauto.NewCounter(prometheus.CounterOpts{
		Name: "ise_apartment_api_apartment_not_found_total",
		Help: "Total number of apartments that were not found",
	})
)

type apartmentAPI struct {
	pb.UnimplementedIseApartmentApiServiceServer
	repo repo.Repo
	repoEvent event.EventRepo
}

// NewApartmentAPI returns api of ise-apartment-api service
func NewApartmentAPI(r repo.Repo, e event.EventRepo) pb.IseApartmentApiServiceServer {
	return &apartmentAPI{repo: r, repoEvent: e}
}
