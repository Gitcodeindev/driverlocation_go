package location

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (ls *LocationService) GetLocation(driverID int64) (string, error) {
	// Реализация метода
	return "", nil
}

func (ls *LocationService) UpdateLocation(driverID int64, location string) error {
	// Реализация метода
	return nil
}

func RegisterRoutes(router *mux.Router, service *LocationService) {
	router.HandleFunc("/location/{id}", getLocationHandler(service)).Methods("GET")
	router.HandleFunc("/location/{id}", updateLocationHandler(service)).Methods("POST")
}

func getLocationHandler(service *LocationService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			http.Error(w, "Invalid driver ID", http.StatusBadRequest)
			return
		}

		location, err := service.GetLocation(id)
		if err != nil {
			http.Error(w, "Failed to get location", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(location)
	}
}

func updateLocationHandler(service *LocationService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			http.Error(w, "Invalid driver ID", http.StatusBadRequest)
			return
		}

		var location string
		err = json.NewDecoder(r.Body).Decode(&location)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		err = service.UpdateLocation(id, location)
		if err != nil {
			http.Error(w, "Failed to update location", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode("Location updated")
	}
}
