package driver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Handler struct {
	service Service
}

func (h *Handler) RegisterDriver(w http.ResponseWriter, r *http.Request) {
	var driver Driver
	err := json.NewDecoder(r.Body).Decode(&driver)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	err = h.service.RegisterDriver(&driver)
	if err != nil {
		http.Error(w, "Failed to register driver", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode("Водитель зарегистрирован")
	if err != nil {
		return
	}
}

func (h *Handler) UpdateDriver() {
}

func (h *Handler) GetDrivers() {
}

func (h *Handler) StartTrip(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	driverID, err := strconv.ParseInt(vars["driverId"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid driver id", http.StatusBadRequest)
		return
	}

	err = h.service.StartTrip(driverID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to start trip: %s", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode("Trip started")
	if err != nil {
		return
	}
}