package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ebdonato/go-deploy-cloud-run/pkg/cep"
	"github.com/ebdonato/go-deploy-cloud-run/pkg/weather"
	"github.com/ebdonato/go-deploy-cloud-run/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type webServerSuite struct {
	suite.Suite
	router http.Handler
}

func (suite *webServerSuite) SetupSuite() {
	apiKey := util.GetEnvVariable("WEATHER_API_KEY")

	weatherApi := weather.InstanceWeatherApi(apiKey)
	viaCep := cep.InstanceViaCep()

	r := NewWebServer(weatherApi, viaCep)

	suite.router = r
}

func (suite *webServerSuite) TestWebServerRunning() {
	ts := httptest.NewServer(suite.router)
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), http.StatusOK, resp.StatusCode)
}

func (suite *webServerSuite) TestWebServerValidCEP() {
	ts := httptest.NewServer(suite.router)
	defer ts.Close()

	url := fmt.Sprintf("%s/%s", ts.URL, "13330350")
	resp, err := http.Get(url)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), http.StatusOK, resp.StatusCode)
}

func (suite *webServerSuite) TestWebServerInvalidCEP() {
	ts := httptest.NewServer(suite.router)
	defer ts.Close()

	url := fmt.Sprintf("%s/%s", ts.URL, "123")
	resp, err := http.Get(url)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), http.StatusUnprocessableEntity, resp.StatusCode)
}

func (suite *webServerSuite) TestWebServerCEPNotFound() {
	ts := httptest.NewServer(suite.router)
	defer ts.Close()

	url := fmt.Sprintf("%s/%s", ts.URL, "00000000")
	resp, err := http.Get(url)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), http.StatusNotFound, resp.StatusCode)
}

func TestWebServerSuite(t *testing.T) {
	suite.Run(t, new(webServerSuite))
}
