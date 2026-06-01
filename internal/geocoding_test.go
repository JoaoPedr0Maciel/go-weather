package internal

import "testing"

func TestValidateCitydata(t *testing.T) {
	err := validateCityData(GeoCodingResponse{}, "")

	if err == nil {
		t.Error("Esperava um erro")
	}
}
