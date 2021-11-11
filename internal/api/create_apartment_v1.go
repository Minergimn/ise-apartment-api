package api

import (
	"context"
	apartment "github.com/ozonmp/ise-apartment-api/internal/model"

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

	apt := a.mapApartmentFromApiToDb(req.Value)

	id, err := a.repo.CreateApartment(ctx, apt)
	if err != nil {
		log.Error().Err(err).Msg("CreateApartmentV1 - adding to apartment db failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	apt.ID = id
	createEvent := &apartment.ApartmentEvent{
		ApartmentId: id,
		Type: apartment.Created,
		Status: apartment.Deferred,
		Entity: apt,
	}

	var events []apartment.ApartmentEvent
	events = append(events, *createEvent)
	err = a.repoEvent.Add(events)
	if err != nil {
		log.Error().Err(err).Msg("CreateApartmentV1 - adding to apartment_events db failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Debug().Msg("CreateApartmentV1 - success")

	return &ise_apartment_api.CreateApartmentV1Response{
		ApartmentId: id,
	}, nil
}
