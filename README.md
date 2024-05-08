# SolarCalc

> Calculate solar rooftop potential using a post address

From the resulting json we can remove the following fields:
- roofSegmentStats
- solarPanelConfigs
- solarPanels
- financialAnalyses
and feed the rest to an LLM model to predict the solar potential.

## Usage

```bash
go build -o solarcalc

export GOOGLE_API_KEY=<your_google_api_key>
./solarcalc "1600 Amphitheatre Parkway, Mountain View, CA" > output.json

# remove solarPotential.roofSegmentStats, solarPotential.solarPanelConfigs, solarPotential.solarPanels:
cat output.json | jq 'del(.solarPotential.roofSegmentStats, .solarPotential.solarPanelConfigs, .solarPotential.solarPanels, .solarPotential.financialAnalyses)'
```

# TODO:
- [ ] convert GeoTiff (https://developers.google.com/maps/documentation/solar/geotiff) to COG and then to PNG - see
    - https://www.cogeo.org/
    - https://github.com/lukeroth/gdal (https://github.com/lukeroth/gdal/blob/master/examples/tiff/tiff.go)
    - https://github.com/airbusgeo/godal
    - https://github.com/airbusgeo/cogger
    - https://github.com/SvenPfiffner/GeoTiffConverter
    - https://github.com/googlemaps-samples/js-solar-potential/blob/main/src/routes/layer.ts
