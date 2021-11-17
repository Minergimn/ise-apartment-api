package api

import (
	"context"
	"github.com/ozonmp/ise-apartment-api/internal/metrics"
	ise_apartment_api "github.com/ozonmp/ise-apartment-api/pkg/ise-apartment-api"
	"google.golang.org/grpc/metadata"

	"github.com/ozonmp/ise-apartment-api/internal/logger"
	apartment "github.com/ozonmp/ise-apartment-api/internal/model"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *apartmentAPI) CreateApartmentV1(
	ctx context.Context,
	req *ise_apartment_api.CreateApartmentV1Request,
) (*ise_apartment_api.CreateApartmentV1Response, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		levels := md.Get("log-level")
		logger.InfoKV(ctx, "got log level", "levels", levels)
		if len(levels) > 0 {
			if parsedLevel, ok := parseLevel(levels[0]); ok {
				newLogger := logger.CloneWithLevel(ctx, parsedLevel)
				ctx = logger.AttachLogger(ctx, newLogger)
			}
		}
	}

	logger.DebugKV(ctx, "CreateApartmentV1 - started")

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, "CreateApartmentV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	apt := a.mapApartmentFromApiToDb(req.Value)

	id, err := a.repo.CreateApartment(ctx, apt)
	if err != nil {
		logger.ErrorKV(ctx, "CreateApartmentV1 - adding to apartment db failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	apt.ID = id
	createEvent := &apartment.ApartmentEvent{
		ApartmentId: id,
		Type:        apartment.Created,
		Status:      apartment.Deferred,
		Entity:      apt,
	}

	var events []apartment.ApartmentEvent
	events = append(events, *createEvent)
	err = a.repoEvent.Add(ctx, events)
	if err != nil {
		logger.ErrorKV(ctx, "CreateApartmentV1 - adding to apartment_events db failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	logger.DebugKV(ctx, "CreateApartmentV1 - success")

	metrics.IncTotalApartmentCUDEvents(metrics.Created)

	return &ise_apartment_api.CreateApartmentV1Response{
		ApartmentId: id,
	}, nil
}
