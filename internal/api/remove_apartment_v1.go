package api

import (
	"context"
	"github.com/ozonmp/ise-apartment-api/internal/logger"
	"google.golang.org/grpc/metadata"

	"github.com/ozonmp/ise-apartment-api/pkg/ise-apartment-api"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *apartmentAPI) RemoveApartmentV1(
	ctx context.Context,
	req *ise_apartment_api.RemoveApartmentV1Request,
) (*ise_apartment_api.RemoveApartmentV1Response, error) {

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

	logger.DebugKV(ctx, "RemoveApartmentV1 - started")

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, "RemoveApartmentV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	found, err := a.repo.DeleteApartment(ctx, req.ApartmentId)
	if err != nil {
		logger.ErrorKV(ctx, "RemoveApartmentV1 - failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	if found == false {
		logger.DebugKV(ctx, "apartment not found", "apartmentId", req.ApartmentId)
		totalApartmentNotFound.Inc()

		return nil, status.Error(codes.NotFound, "apartment not found")
	}

	logger.DebugKV(ctx, "RemoveApartmentV1 - success")

	return &ise_apartment_api.RemoveApartmentV1Response{
		Found: true,
	}, nil
}
