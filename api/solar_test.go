package api

import (
	"testing"

	"github.com/mheers/solar-calc-go/models"
	"github.com/stretchr/testify/require"
)

func TestGetInsights(t *testing.T) {
	config, err := models.GetConfig()
	require.NoError(t, err)

	solarAgent, err := NewSolarAgent(config)
	require.NoError(t, err)

	lat, long := 51.688034950865166, 8.692962059536352 // paderborn
	insight, err := solarAgent.GetInsights(lat, long)
	require.NoError(t, err)

	address := "Husener Str. 51, 33098 Paderborn"
	err = solarAgent.SaveInsight(address, insight)
	require.NoError(t, err)
}

func TestDownloadGeoTiffs(t *testing.T) {
	config, err := models.GetConfig()
	require.NoError(t, err)

	solarAgent, err := NewSolarAgent(config)
	require.NoError(t, err)

	lat, long := 51.711386, 8.7607221 // paderborn
	insight, err := solarAgent.GetInsights(lat, long)
	require.NoError(t, err)

	address := "Husener Str. 51, 33098 Paderborn"

	dataLayers, err := solarAgent.GetDataLayers(insight)
	require.NoError(t, err)

	images, err := solarAgent.DownloadGeoTiffs(address, dataLayers)
	require.NoError(t, err)
	require.NotEmpty(t, images)
}
