// pkg/http/router.go
package http

import (
	"learn-back/internal/service"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRoutes(router *mux.Router, airportSvc *service.AirportService, aircraftSvc *service.AircraftService) {
	router.HandleFunc("/airports", GetAllAirportsHundlers(airportSvc)).Methods(http.MethodGet)
	router.HandleFunc("/aircrafts", GetAllAircraftsHundlers(aircraftSvc)).Methods(http.MethodGet)
}
