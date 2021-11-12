package api

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ozonmp/ise-apartment-api/internal/mocks"
	ise_apartment_api "github.com/ozonmp/ise-apartment-api/pkg/ise-apartment-api"
)

func Test_apartmentAPI_RemoveApartmentV1(t *testing.T) {
	ctrl := gomock.NewController(t)
	repo := mocks.NewMockRepo(ctrl)
	repoEvent := mocks.NewMockEventRepo(ctrl)
	api := NewApartmentAPI(repo, repoEvent)
	ctx := context.Background()

	tests := []struct {
		name    string
		req     ise_apartment_api.RemoveApartmentV1Request
		want    *ise_apartment_api.RemoveApartmentV1Response
		wantErr bool
	}{
		{
			name:    "Id more than 0. Validation should pass",
			req:     ise_apartment_api.RemoveApartmentV1Request{ApartmentId: 1},
			want:    &ise_apartment_api.RemoveApartmentV1Response{},
			wantErr: false,
		},
		{
			name:    "Id is 0. Validation should fail",
			req:     ise_apartment_api.RemoveApartmentV1Request{ApartmentId: 0},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := api.RemoveApartmentV1(ctx, &tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveApartmentV1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveApartmentV1() got = %v, want %v", got, tt.want)
			}
		})
	}
}
