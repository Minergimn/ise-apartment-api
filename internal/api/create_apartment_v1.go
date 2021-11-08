package api

import (
	"context"

	"github.com/ozonmp/ise-apartment-api/pkg/ise-apartment-api"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *apartmentAPI) CreateApartmentV1(
	ctx context.Context,
	req *ise_apartment_api.CreateApartmentV1Request,
) (*ise_apartment_api.CreateApartmentV1Response, error) {

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("CreateApartmentV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	id, err := a.repo.CreateApartment(ctx, a.mapApartmentFromApiToDb(req.Value))
	if err != nil {
		log.Error().Err(err).Msg("CreateApartmentV1 - failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Debug().Msg("CreateApartmentV1 - success")

	return &ise_apartment_api.CreateApartmentV1Response{
		ApartmentId: id,
	}, nil
}
