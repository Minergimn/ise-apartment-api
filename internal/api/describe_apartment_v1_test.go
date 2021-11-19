package api

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/ozonmp/ise-apartment-api/internal/mocks"
	model "github.com/ozonmp/ise-apartment-api/internal/model"
	ise_apartment_api "github.com/ozonmp/ise-apartment-api/pkg/ise-apartment-api"

	"github.com/golang/mock/gomock"
)

func Test_apartmentAPI_DescribeApartmentV1(t *testing.T) {
	ctrl := gomock.NewController(t)
	repo := mocks.NewMockRepo(ctrl)
	repoEvent := mocks.NewMockEventRepo(ctrl)
	api := NewApartmentAPI(repo, repoEvent)
	ctx := context.Background()

	repo.EXPECT().GetApartment(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx2 context.Context, id uint64) (*model.Apartment, error) {
		a := new(model.Apartment)
		a.ID = id
		a.Owner = fmt.Sprintf("%d", id)
		a.Object = fmt.Sprintf("%d", id)
		return a, nil
	}).AnyTimes()

	tests := []struct {
		name    string
		req     ise_apartment_api.DescribeApartmentV1Request
		want    *ise_apartment_api.DescribeApartmentV1Response
		wantErr bool
	}{
		{
			name:    "Id more than 0. Validation should pass",
			req:     ise_apartment_api.DescribeApartmentV1Request{ApartmentId: 1},
			want:    &ise_apartment_api.DescribeApartmentV1Response{Value: getApartment(1, "1", "1")},
			wantErr: false,
		},
		{
			name:    "Id is 0. Validation should fail",
			req:     ise_apartment_api.DescribeApartmentV1Request{ApartmentId: 0},
			wantErr: true,
		},
	}

	for _, tt := range tests { //nolint:govet
		t.Run(tt.name, func(t *testing.T) {
			got, err := api.DescribeApartmentV1(ctx, &tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("DescribeApartmentV1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DescribeApartmentV1() got = %v, want %v", got, tt.want)
			}
		})
	}
}
