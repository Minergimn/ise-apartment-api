package api

import (
	"context"

	"github.com/ozonmp/ise-apartment-api/pkg/ise-apartment-api"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (o *apartmentAPI) ListApartmentsV1(
	ctx context.Context,
	req *ise_apartment_api.ListApartmentsV1Request,
) (*ise_apartment_api.ListApartmentsV1Response, error) {

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("ListApartmentsV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	//o.repo.ListApartments calling will be here
	log.Debug().Str("ListApartmentsV1 input", req.String()).Msg("Just a log instead of an implementation")

	var items []*ise_apartment_api.Apartment

	return &ise_apartment_api.ListApartmentsV1Response{
		Items: items,
	}, nil
}
