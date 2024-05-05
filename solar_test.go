package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetInsights(t *testing.T) {
	config, err := getConfig()
	require.NoError(t, err)

	solarAgent, err := NewSolarAgent(config)
	require.NoError(t, err)

	lat, long := 51.711386, 8.7607221 // paderborn
	insight, err := solarAgent.GetInsights(lat, long)
	require.NoError(t, err)

	address := "Husener Str. 51, 33098 Paderborn"
	err = solarAgent.saveInsight(address, insight)
	require.NoError(t, err)
}

func TestDownloadGeoTiffs(t *testing.T) {
	config, err := getConfig()
	require.NoError(t, err)

	solarAgent, err := NewSolarAgent(config)
	require.NoError(t, err)

	lat, long := 51.711386, 8.7607221 // paderborn
	insight, err := solarAgent.GetInsights(lat, long)
	require.NoError(t, err)

	address := "Husener Str. 51, 33098 Paderborn"

	dataLayers, err := solarAgent.getDataLayers(insight)
	require.NoError(t, err)

	err = solarAgent.downloadGeoTiffs(address, dataLayers)
	require.NoError(t, err)
}
