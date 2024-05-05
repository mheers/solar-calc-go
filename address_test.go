package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetCoordinatesFromAddress(t *testing.T) {
	config, err := getConfig()
	require.NoError(t, err)

	solarAgent, err := NewSolarAgent(config)
	require.NoError(t, err)

	address := "Husener Str. 51, 33098 Paderborn"
	validation, err := solarAgent.validateAddress(address)
	require.NoError(t, err)

	err = solarAgent.saveValidationResult(address, validation)
	require.NoError(t, err)

	lat, long, err := solarAgent.getCoordinatesFromValidation(validation)
	require.NoError(t, err)

	require.Equal(t, 51.711386, lat)
	require.Equal(t, 8.7607221, long)
}
