package api

import (
	"github.com/ozonmp/ise-apartment-api/internal/model"
	"github.com/ozonmp/ise-apartment-api/pkg/ise-apartment-api"
)

func (a *apartmentAPI) mapApartmentFromDbToApi(from *apartment.Apartment) *ise_apartment_api.Apartment {
	return &ise_apartment_api.Apartment{
		Id:     from.ID,
		Object: from.Object,
		Owner:  from.Owner,
	}
}

func (a *apartmentAPI) mapApartmentFromApiToDb(from *ise_apartment_api.Apartment) *apartment.Apartment {
	return &apartment.Apartment{
		Object: from.Object,
		Owner:  from.Owner,
	}
}
