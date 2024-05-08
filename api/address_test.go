package api

import (
	"testing"

	"github.com/mheers/solar-calc-go/models"
	"github.com/stretchr/testify/require"
)

func TestGetCoordinatesFromAddress(t *testing.T) {
	config, err := models.GetConfig()
	require.NoError(t, err)

	solarAgent, err := NewSolarAgent(config)
	require.NoError(t, err)

	address := "Husener Str. 51, 33098 Paderborn"
	validation, err := solarAgent.ValidateAddress(address)
	require.NoError(t, err)

	err = solarAgent.SaveValidationResult(address, validation)
	require.NoError(t, err)

	lat, long, err := solarAgent.GetCoordinatesFromValidation(validation)
	require.NoError(t, err)

	require.Equal(t, 51.711386, lat)
	require.Equal(t, 8.7607221, long)
}
