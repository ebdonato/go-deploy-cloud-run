package weather

import (
	"testing"

	"github.com/ebdonato/go-deploy-cloud-run/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type weatherApiSuite struct {
	suite.Suite
	weatherApi WeatherApi
}

func (suite *weatherApiSuite) SetupSuite() {
	apiKey := util.GetEnvVariable("WEATHER_API_KEY")
	suite.weatherApi = *InstanceWeatherApi(apiKey)
}

func (suite *weatherApiSuite) TestGetLocationInfo() {
	location := "Indaiatuba,SP"
	_, err := suite.weatherApi.GetTemperature(location)
	assert.NoError(suite.T(), err)

	location = "Cachoeiro de Itapemirim,ES"
	_, err = suite.weatherApi.GetTemperature(location)
	assert.NoError(suite.T(), err)

	location = "X"
	_, err = suite.weatherApi.GetTemperature(location)
	assert.Error(suite.T(), err)
}

func TestWeatherApiSuite(t *testing.T) {
	suite.Run(t, new(weatherApiSuite))
}
