package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mheers/solar-calc-go/api"
	"github.com/mheers/solar-calc-go/models"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <address>", os.Args[0])
		os.Exit(1)
	}

	address := os.Args[1]

	if len(address) < 10 {
		log.Fatal("Address must be at least 10 characters long")
		os.Exit(1)
	}

	config, err := models.GetConfig()
	if err != nil {
		panic(err)
	}

	solarAgent, err := api.NewSolarAgent(config)
	if err != nil {
		panic(err)
	}

	// validate address
	validation, err := solarAgent.ValidateAddress(address)
	if err != nil {
		panic(err)
	}

	// save validation result
	if err = solarAgent.SaveValidationResult(address, validation); err != nil {
		panic(err)
	}

	// get coordinates from validation
	lat, long, err := solarAgent.GetCoordinatesFromValidation(validation)
	if err != nil {
		panic(err)
	}

	// get insights
	insight, err := solarAgent.GetInsights(lat, long)
	if err != nil {
		panic(err)
	}

	// save insights
	err = solarAgent.SaveInsight(address, insight)
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
