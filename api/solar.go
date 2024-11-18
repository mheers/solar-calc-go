package api

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"

	"github.com/mheers/solar-calc-go/models"
	"google.golang.org/api/addressvalidation/v1"
	"google.golang.org/api/option"
	"google.golang.org/api/solar/v1"
)

type SolarAgent struct {
	config         *models.Config
	solarService   *solar.Service
	addressService *addressvalidation.Service
}

func NewSolarAgent(config *models.Config) (*SolarAgent, error) {
	solarService, err := GetSolarService(config)
	if err != nil {
		return nil, err
	}

	addressService, err := GetAddressService(config)
	if err != nil {
		return nil, err
	}

	return &SolarAgent{
		config:         config,
		solarService:   solarService,
		addressService: addressService,
	}, nil
}

func GetSolarService(config *models.Config) (*solar.Service, error) {
	return solar.NewService(context.Background(), option.WithAPIKey(config.APIKey))
}

func (sa *SolarAgent) GetInsights(lat, long float64) (*solar.BuildingInsights, error) {
	bis := solar.NewBuildingInsightsService(sa.solarService)
	insight, err := bis.FindClosest().LocationLatitude(lat).LocationLongitude(long).RequiredQuality("MEDIUM").Do()
	if err != nil {
		return nil, err
	}

	return insight, nil
}

func (sa *SolarAgent) SaveInsight(address string, insight *solar.BuildingInsights) error {
	addressB64 := base64.StdEncoding.EncodeToString([]byte(address))
	fileName := "results/insights_" + addressB64 + ".json"
	j, err := insight.MarshalJSON()
	if err != nil {
		return err
	}
	if err = os.MkdirAll("results", 0755); err != nil {
		return err
	}
	return os.WriteFile(fileName, j, 0644)
}

func (sa *SolarAgent) GetDataLayers(insight *solar.BuildingInsights) (*solar.DataLayers, error) {
	dls := solar.NewDataLayersService(sa.solarService)
	lat, long := insight.Center.Latitude, insight.Center.Longitude
	dataLayers, err := dls.Get().LocationLatitude(lat).LocationLongitude(long).RadiusMeters(50).RequiredQuality("LOW").Do()
	if err != nil {
		return nil, err
	}

	return dataLayers, nil
}

func (sa *SolarAgent) DownloadGeoTiffs(address string, dataLayers *solar.DataLayers) (map[string]string, error) {
	addressB64 := base64.StdEncoding.EncodeToString([]byte(address))
	// gts := solar.NewGeoTiffService(sa.solarService) // TODO: this service is broken

	urls := map[string]string{
		"annualFlux":  dataLayers.AnnualFluxUrl,
		"dsm":         dataLayers.DsmUrl,
		"mask":        dataLayers.MaskUrl,
		"rgb":         dataLayers.RgbUrl,
		"monthlyFlux": dataLayers.MonthlyFluxUrl,
	}

	for i, url := range dataLayers.HourlyShadeUrls {
		urls[fmt.Sprintf("shade_%d", i)] = url
	}

	filePaths := make(map[string]string)

	wg := sync.WaitGroup{}
	wg.Add(len(urls))

	errs := []error{}

	errors := make(chan error)

	for name, url := range urls {
		fileName := "results/" + addressB64 + "_" + name + ".tiff"
		filePaths[name] = fileName

		go func(name, url, fileName string) {
			// download the file
			url = url + "&key=" + sa.config.APIKey
			resp, err := http.DefaultClient.Get(url)
			if err != nil {
				errors <- err
				return
			}

			// save the file
			if err := os.MkdirAll("results/", 0755); err != nil {
				errors <- err
				return
			}

			data, err := io.ReadAll(resp.Body)
			if err != nil {
				errors <- err
				return
			}

			if err := os.WriteFile(fileName, data, 0644); err != nil {
				errors <- err
				return
			}
			wg.Done()
		}(name, url, fileName)

	}

	wg.Wait()

	select {
	case err := <-errors:
		errs = append(errs, err)
	default:
	}

	if len(errs) > 0 {
		return nil, errs[0]
	}

	return filePaths, nil
}
