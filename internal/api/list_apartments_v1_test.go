package api

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ozonmp/ise-apartment-api/internal/mocks"
	ise_apartment_api "github.com/ozonmp/ise-apartment-api/pkg/ise-apartment-api"
)

func Test_apartmentAPI_ListApartmentsV1(t *testing.T) {
	ctrl := gomock.NewController(t)
	repo := mocks.NewMockRepo(ctrl)
	repoEvent := mocks.NewMockEventRepo(ctrl)
	api := NewApartmentAPI(repo, repoEvent)
	ctx := context.Background()

	var items []*ise_apartment_api.Apartment

	tests := []struct {
		name    string
		req     ise_apartment_api.ListApartmentsV1Request
		want    *ise_apartment_api.ListApartmentsV1Response
		wantErr bool
	}{
		{
			name:    "Any params specified. Validation should pass",
			req:     ise_apartment_api.ListApartmentsV1Request{Params: getParams()},
			want:    &ise_apartment_api.ListApartmentsV1Response{Items: items},
			wantErr: false,
		},
		{
			name:    "Params not specified. Validation should pass",
			req:     ise_apartment_api.ListApartmentsV1Request{},
			want:    &ise_apartment_api.ListApartmentsV1Response{Items: items},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := api.ListApartmentsV1(ctx, &tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListApartmentsV1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListApartmentsV1() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func getParams() *ise_apartment_api.ListApartmentsV1Request_ListApartmentsParams {
	p := new(ise_apartment_api.ListApartmentsV1Request_ListApartmentsParams)
	p.Owner = "Owner"
	p.Object = "Object"
	p.Offset = 1
	p.Limit = 10
	p.Ids = []uint64{1, 2, 3}
	return p
}
