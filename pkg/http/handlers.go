package http

import (
	"encoding/json"
	"fmt"
	"learn-back/internal/service"
	"net/http"
)

func GetAllAirportsHandler(svc *service.AirportService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		airports, err := svc.GetAllAirports()
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in f", r)
			}
		}()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(airports)
	}
}

func GetAllAircraftsHandler(svc *service.AircraftService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		aircrafts, err := svc.GetAllAircrafts()
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in f", r)
			}
		}()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(aircrafts)
	}
}
