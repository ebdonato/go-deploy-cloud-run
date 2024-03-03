package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/ebdonato/go-deploy-cloud-run/pkg/cep"
	"github.com/ebdonato/go-deploy-cloud-run/pkg/weather"
	"github.com/ebdonato/go-deploy-cloud-run/util"
	"github.com/go-chi/chi/v5"
)

func NewWebServer() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", handlerHealth)
	r.Get("/{cep}", handlerCEP)

	return r
}

type CEPResponse struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

func handlerHealth(w http.ResponseWriter, r *http.Request) {
	defer log.Println("App health checked")
	w.Write([]byte("Server is running"))
}

func handlerCEP(w http.ResponseWriter, r *http.Request) {
	cepParams := strings.TrimSpace(r.URL.Path[1:])

	if !util.IsValidCEP(cepParams) {
		message := "Invalid CEP"
		log.Println(message)
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(message))
		return
	}

	var viaCep cep.CEP = cep.InstanceViaCep()

	location, err := viaCep.FindLocation(cepParams)
	if err != nil {
		message := "CEP not Found"
		log.Println(message)
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(message))
		return
	}

	var openWeather weather.Weather = weather.InstanceOpenWeather()

	temperature, err := openWeather.GetTemperature(location)
	if err != nil {
		message := "Internal Server Error"
		log.Println(message)
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(message))
		return
	}

	response := CEPResponse{
		TempC: temperature.Celsius,
		TempF: temperature.Fahrenheit,
		TempK: temperature.Kelvin,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
