package api

import (
	"context"
	"math/rand"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/ozonmp/ise-apartment-api/internal/mocks"
	ise_apartment_api "github.com/ozonmp/ise-apartment-api/pkg/ise-apartment-api"

	"github.com/golang/mock/gomock"
)

func Test_apartmentAPI_CreateApartmentV1(t *testing.T) {
	ctrl := gomock.NewController(t)
	repo := mocks.NewMockRepo(ctrl)
	api := NewApartmentAPI(repo)
	ctx := context.Background()
	wantResponse := &ise_apartment_api.CreateApartmentV1Response{}

	tests := []struct {
		name    string
		req     ise_apartment_api.CreateApartmentV1Request
		want    *ise_apartment_api.CreateApartmentV1Response
		wantErr bool
	}{
		{
			name:    "All field specified. Validation should pass",
			req:     ise_apartment_api.CreateApartmentV1Request{Value: getApartment(1, "Owner", "Object")},
			want:    wantResponse,
			wantErr: false,
		},
		{
			name:    "Owner and object specified. Validation should pass",
			req:     ise_apartment_api.CreateApartmentV1Request{Value: getApartment(0, "Owner", "Object")},
			want:    wantResponse,
			wantErr: false,
		},
		{
			name:    "Owner not specified. Validation should fail",
			req:     ise_apartment_api.CreateApartmentV1Request{Value: getApartment(1, "", "Object")},
			wantErr: true,
		},
		{
			name:    "Object not specified. Validation should fail",
			req:     ise_apartment_api.CreateApartmentV1Request{Value: getApartment(1, "Owner", "")},
			wantErr: true,
		},
		{
			name:    "Owner more than 1000 runes. Validation should fail",
			req:     ise_apartment_api.CreateApartmentV1Request{Value: getApartment(1, RandStringBytesMaskImprSrcSB(1001), "Object")},
			wantErr: true,
		},
		{
			name:    "Object more than 1000 runes. Validation should fail",
			req:     ise_apartment_api.CreateApartmentV1Request{Value: getApartment(1, "Owner", RandStringBytesMaskImprSrcSB(1001))},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := api.CreateApartmentV1(ctx, &tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateApartmentV1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateApartmentV1() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func getApartment(id uint64, owner string, object string) *ise_apartment_api.Apartment {
	a := new(ise_apartment_api.Apartment)
	a.Id = id
	a.Owner = owner
	a.Object = object

	return a
}

//Rand string with length from https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func RandStringBytesMaskImprSrcSB(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}
