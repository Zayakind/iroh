// pkg/http/router.go
package http

import (
	"learn-back/internal/service"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRoutes(router *mux.Router, airportSvc *service.AirportService, aircraftSvc *service.AircraftService) {
	router.HandleFunc("/airports", GetAllAirportsHandler(airportSvc)).Methods(http.MethodGet)
	router.HandleFunc("/aircrafts", GetAllAircraftsHandler(aircraftSvc)).Methods(http.MethodGet)
}
