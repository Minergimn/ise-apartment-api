package api

import (
	"context"

	"github.com/ozonmp/ise-apartment-api/pkg/ise-apartment-api"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *apartmentAPI) ListApartmentsV1(
	ctx context.Context,
	req *ise_apartment_api.ListApartmentsV1Request,
) (*ise_apartment_api.ListApartmentsV1Response, error) {

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("ListApartmentsV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	offset := uint64(0)
	limit := uint64(0)
	owner := ""
	object := ""
	var ids []uint64
	if req.Params != nil{
		offset = req.Params.Offset
		limit = req.Params.Limit
		owner = req.Params.Owner
		object = req.Params.Object
		ids = req.Params.Ids
	}

	apartments, err := a.repo.ListApartments(ctx, offset, limit, owner, object, ids)
	if err != nil {
		log.Error().Err(err).Msg("ListApartmentsV1 - failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	if apartments == nil {
		log.Debug().Str("params", req.Params.String()).Msg("apartments not found")
		totalApartmentNotFound.Inc()

		return nil, status.Error(codes.NotFound, "apartments not found")
	}

	log.Debug().Msg("ListApartmentsV1 - success")

	var items []*ise_apartment_api.Apartment

	for _, apartment := range apartments {
		items = append(items, a.mapApartmentFromDbToApi(&apartment))
	}

	return &ise_apartment_api.ListApartmentsV1Response{
		Items: items,
	}, nil
}
