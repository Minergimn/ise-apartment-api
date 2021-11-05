package api

import (
	"context"

	"github.com/ozonmp/ise-apartment-api/pkg/ise-apartment-api"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (o *apartmentAPI) RemoveApartmentV1(
	ctx context.Context,
	req *ise_apartment_api.RemoveApartmentV1Request,
) (*ise_apartment_api.RemoveApartmentV1Response, error) {

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("RemoveApartmentV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Debug().Str("RemoveApartmentV1 input", req.String()).Msg("Just a log instead of an implementation")

	return &ise_apartment_api.RemoveApartmentV1Response{}, nil
}
