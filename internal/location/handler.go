package location

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *Service) GetLocation(int64) (string, error) {
	return "", nil
}

func (s *Service) UpdateLocation() error {
	return nil
}

func updateLocationHandler(service *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		_, err := strconv.ParseInt(vars["id"], 10, 64)
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

		err = service.UpdateLocation()
		if err != nil {
			http.Error(w, "Failed to update location", http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode("Location updated")
		if err != nil {
			return
		}
	}
}
