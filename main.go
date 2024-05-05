package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

type Config struct {
	APIKey string
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: solarcalc <address>")
		os.Exit(1)
	}

	address := os.Args[1]

	if len(address) < 10 {
		log.Fatal("Address must be at least 10 characters long")
		os.Exit(1)
	}

	config, err := getConfig()
	if err != nil {
		panic(err)
	}

	solarAgent, err := NewSolarAgent(config)
	if err != nil {
		panic(err)
	}

	// validate address
	validation, err := solarAgent.validateAddress(address)
	if err != nil {
		panic(err)
	}

	// save validation result
	if err = solarAgent.saveValidationResult(address, validation); err != nil {
		panic(err)
	}

	// get coordinates from validation
	lat, long, err := solarAgent.getCoordinatesFromValidation(validation)
	if err != nil {
		panic(err)
	}

	// get insights
	insight, err := solarAgent.GetInsights(lat, long)
	if err != nil {
		panic(err)
	}

	// save insights
	err = solarAgent.saveInsight(address, insight)
	if err != nil {
		panic(err)
	}

	// // get data layers
	// dataLayers, err := solarAgent.getDataLayers(insight)
	// if err != nil {
	// 	panic(err)
	// }

	// // download GeoTiffs
	// err = solarAgent.downloadGeoTiffs(address, dataLayers)
	// if err != nil {
	// 	panic(err)
	// }

	// print insights
	j, err := insight.MarshalJSON()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(j))
}

func getConfig() (*Config, error) {
	apiKey := os.Getenv("GOOGLE_API_KEY")
	if apiKey == "" {
		return nil, errors.New("GOOGLE_API_KEY environment variable not set")
	}

	return &Config{
		APIKey: apiKey,
	}, nil
}
