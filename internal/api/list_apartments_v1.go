package api

import (
	"context"
	ise_apartment_api "github.com/ozonmp/ise-apartment-api/pkg/ise-apartment-api"
	"google.golang.org/grpc/metadata"

	"github.com/ozonmp/ise-apartment-api/internal/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *apartmentAPI) ListApartmentsV1(
	ctx context.Context,
	req *ise_apartment_api.ListApartmentsV1Request,
) (*ise_apartment_api.ListApartmentsV1Response, error) {

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

	logger.DebugKV(ctx, "ListApartmentsV1 - started")

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, "ListApartmentsV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	var (
		ids           []uint64
		limit, offset uint64
		owner, object string
	)

	if req.Params != nil {
		offset = req.Params.Offset
		limit = req.Params.Limit
		owner = req.Params.Owner
		object = req.Params.Object
		ids = req.Params.Ids
	}

	apartments, err := a.repo.ListApartments(ctx, offset, limit, owner, object, ids)
	if err != nil {
		logger.ErrorKV(ctx, "ListApartmentsV1 - failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	logger.DebugKV(ctx, "ListApartmentsV1 - success")

	var items []*ise_apartment_api.Apartment

	for _, apartment := range apartments {
		items = append(items, a.mapApartmentFromDbToAPI(&apartment)) //nolint:gosec
	}

	return &ise_apartment_api.ListApartmentsV1Response{
		Items: items,
	}, nil
}
