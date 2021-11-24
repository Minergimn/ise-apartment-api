package api

import (
	"context"
	"database/sql"
	"errors"
	"github.com/ozonmp/ise-apartment-api/internal/metrics"
	ise_apartment_api "github.com/ozonmp/ise-apartment-api/pkg/ise-apartment-api"
	"google.golang.org/grpc/metadata"

	"github.com/ozonmp/ise-apartment-api/internal/logger"
	apartment "github.com/ozonmp/ise-apartment-api/internal/model"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *apartmentAPI) UpdateApartmentV1(
	ctx context.Context,
	req *ise_apartment_api.UpdateApartmentV1Request,
) (*ise_apartment_api.UpdateApartmentV1Response, error) {

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

	logger.DebugKV(ctx, "UpdateApartmentV1 - started")

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, "UpdateApartmentV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	apt := a.mapApartmentFromAPIToDb(req.Value)

	apt.ID = req.ApartmentId

	updated, err := a.repo.UpdateApartment(ctx, apt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logger.DebugKV(ctx, "apartment not found", "apartmentId", req.ApartmentId)
			metrics.IncTotalApartmentNotFound()

			return nil, status.Error(codes.NotFound, "apartment not found")
		}

		logger.ErrorKV(ctx, "UpdateApartmentV1 - apartment update in db failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	createEvent := &apartment.Event{
		ApartmentID: req.ApartmentId,
		Type:        apartment.Updated,
		Status:      apartment.Deferred,
		Entity:      apt.MapToApartmentPaypoad(),
	}

	var events []apartment.Event
	events = append(events, *createEvent)
	err = a.repoEvent.Add(ctx, events)
	if err != nil {
		logger.ErrorKV(ctx, "UpdateApartmentV1 - adding to apartment_events db failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	logger.DebugKV(ctx, "UpdateApartmentV1 - success")

	metrics.IncTotalApartmentCUDEvents(metrics.Created)

	return &ise_apartment_api.UpdateApartmentV1Response{
		Updated: updated,
	}, nil
}
