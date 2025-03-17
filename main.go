package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/robertoduessmann/weather-api/config"
	"github.com/robertoduessmann/weather-api/controller"
	v2 "github.com/robertoduessmann/weather-api/controller/v2"
)

func main() {
	weather := mux.NewRouter()
	weather.Path("/weather/{city}").Methods(http.MethodGet).HandlerFunc(controller.CurrentWeather)

	weather.
		Path("/v2/weather/{city}").
		Queries("unit", "{unit}").
		Methods(http.MethodGet).
		HandlerFunc(v2.CurrentWeather)

	weather.
		Path("/v2/weather/{city}").
		Methods(http.MethodGet).
		HandlerFunc(v2.CurrentWeather)

	// Configure CORS with proper options
	corsMiddleware := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),    // Allow requests from any origin
		handlers.AllowedMethods([]string{"GET"}),  // Allow only GET requests
		handlers.AllowedHeaders([]string{"Content-Type", "Accept"}),
		handlers.ExposedHeaders([]string{}),
	)

	if err := http.ListenAndServe(":"+config.Get().Port, corsMiddleware(weather)); err != nil {
		log.Fatal(err)
	}
}
