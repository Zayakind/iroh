package http

import (
	"encoding/json"
	"learn-back/internal/service"
	"net/http"
)

func GetAllAirportsHundlers(svc *service.AirportService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		airports, err := svc.GetAllAirports()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(airports)
	}
}

func GetAllAircraftsHundlers(svc *service.AircraftService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		aircrafts, err := svc.GetAllAircrafts()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(aircrafts)
	}
}
