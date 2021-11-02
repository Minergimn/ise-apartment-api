package api

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ozonmp/ise-apartment-api/internal/repo"

	pb "github.com/ozonmp/ise-apartment-api/pkg/ise-apartment-api"
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
}

// NewApartmentAPI returns api of ise-apartment-api service
func NewApartmentAPI(r repo.Repo) pb.IseApartmentApiServiceServer {
	return &apartmentAPI{repo: r}
}

func (o *apartmentAPI) DescribeApartmentV1(
	ctx context.Context,
	req *pb.DescribeApartmentV1Request,
) (*pb.DescribeApartmentV1Response, error) {

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("DescribeApartmentV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	apartment, err := o.repo.DescribeApartment(ctx, req.ApartmentId)
	if err != nil {
		log.Error().Err(err).Msg("DescribeApartmentV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	if apartment == nil {
		log.Debug().Uint64("apartmentId", req.ApartmentId).Msg("apartment not found")
		totalApartmentNotFound.Inc()

		return nil, status.Error(codes.NotFound, "apartment not found")
	}

	log.Debug().Msg("DescribeApartmentV1 - success")

	return &pb.DescribeApartmentV1Response{
		Value: &pb.Apartment{
			Id:  apartment.ID,
			Foo: apartment.Foo,
		},
	}, nil
}
