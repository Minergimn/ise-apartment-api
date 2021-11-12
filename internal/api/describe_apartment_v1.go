package api

import (
	"context"
	"github.com/ozonmp/ise-apartment-api/pkg/ise-apartment-api"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *apartmentAPI) DescribeApartmentV1(
	ctx context.Context,
	req *ise_apartment_api.DescribeApartmentV1Request,
) (*ise_apartment_api.DescribeApartmentV1Response, error) {

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("DescribeApartmentV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	apartment, err := a.repo.GetApartment(ctx, req.ApartmentId)
	if err != nil {
		log.Error().Err(err).Msg("DescribeApartmentV1 - failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	if apartment == nil {
		log.Debug().Uint64("apartmentId", req.ApartmentId).Msg("apartment not found")
		totalApartmentNotFound.Inc()

		return nil, status.Error(codes.NotFound, "apartment not found")
	}

	log.Debug().Msg("DescribeApartmentV1 - success")

	return &ise_apartment_api.DescribeApartmentV1Response{
		Value: a.mapApartmentFromDbToApi(apartment),
	}, nil
}
