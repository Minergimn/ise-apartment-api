package api

import (
	"context"

	"github.com/ozonmp/ise-apartment-api/internal/repo"
	pb "github.com/ozonmp/ise-apartment-api/pkg/ise-apartment-api"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
			Object: apartment.Object,
			Owner: apartment.Owner,
		},
	}, nil
}

func (o *apartmentAPI) CreateApartmentV1(
	ctx context.Context,
	req *pb.CreateApartmentV1Request,
) (*pb.CreateApartmentV1Response, error){

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("CreateApartmentV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	//o.repo.CreateApartment calling will be here
	log.Debug().Str("CreateApartmentV1 input", req.String()).Msg("Just a log instead of an implementation")

	return nil, nil
}

func (o *apartmentAPI) ListApartmentsV1(
	ctx context.Context,
	req *pb.ListApartmentsV1Request,
) (*pb.ListApartmentsV1Response, error){

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("ListApartmentsV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	//o.repo.ListApartments calling will be here
	log.Debug().Str("ListApartmentsV1 input", req.String()).Msg("Just a log instead of an implementation")

	return nil, nil
}

func (o *apartmentAPI) RemoveApartmentV1(
	ctx context.Context,
	req *pb.RemoveApartmentV1Request,
) (*pb.RemoveApartmentV1Response, error){

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("RemoveApartmentV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	//o.repo.DeleteApartment calling will be here
	log.Debug().Str("RemoveApartmentV1 input", req.String()).Msg("Just a log instead of an implementation")

	return nil, nil
}