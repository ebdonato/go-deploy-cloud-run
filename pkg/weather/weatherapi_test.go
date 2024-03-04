package weather

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type weatherApiSuite struct {
	suite.Suite
	weatherApi WeatherApi
}

func (suite *weatherApiSuite) SetupSuite() {
	suite.weatherApi = *InstanceWeatherApi()
}

func (suite *weatherApiSuite) TestGetLocationInfo() {
	location := "Indaiatuba,SP"
	_, err := suite.weatherApi.GetTemperature(location)
	assert.NoError(suite.T(), err)

	location = "X"
	_, err = suite.weatherApi.GetTemperature(location)
	assert.Error(suite.T(), err)
}

func TestWeatherApiSuite(t *testing.T) {
	suite.Run(t, new(weatherApiSuite))
}
