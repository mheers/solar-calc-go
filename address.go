package main

import (
	"context"
	"encoding/base64"
	"os"

	"google.golang.org/api/addressvalidation/v1"
	"google.golang.org/api/option"
)

func GetAddressService(config *Config) (*addressvalidation.Service, error) {
	return addressvalidation.NewService(context.Background(), option.WithAPIKey(config.APIKey))
}

func (sa *SolarAgent) validateAddress(address string) (*addressvalidation.GoogleMapsAddressvalidationV1ValidationResult, error) {
	response, err := sa.addressService.V1.ValidateAddress(&addressvalidation.GoogleMapsAddressvalidationV1ValidateAddressRequest{
		Address: &addressvalidation.GoogleTypePostalAddress{
			AddressLines: []string{address},
		},
	}).Do()
	if err != nil {
		return nil, err
	}
	return response.Result, nil
}

func (sa *SolarAgent) getCoordinatesFromValidation(result *addressvalidation.GoogleMapsAddressvalidationV1ValidationResult) (float64, float64, error) {
	return result.Geocode.Location.Latitude, result.Geocode.Location.Longitude, nil
}

// Save the validation result to a file
func (sa *SolarAgent) saveValidationResult(address string, result *addressvalidation.GoogleMapsAddressvalidationV1ValidationResult) error {
	addressB64 := base64.StdEncoding.EncodeToString([]byte(address))
	fileName := "results/validation_result_" + addressB64 + ".json"
	j, err := result.MarshalJSON()
	if err != nil {
		return err
	}
	if err = os.MkdirAll("results", 0755); err != nil {
		return err
	}
	return os.WriteFile(fileName, j, 0644)
}
