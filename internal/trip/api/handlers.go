package api

import (
	_ "context"
	"net/http"
	"strconv"

	"github.com/Gitcodeindev/driverlocation_go/internal/trip"
	"github.com/gorilla/mux"
)

type TripHandlers struct {
	service *trip.Service
}

type TripService struct {
}

func (s *TripService) AcceptTrip() error {
	return nil
}

func (s *TripService) StartTrip() error {
	return nil
}

func (s *TripService) EndTrip() error {
	return nil
}

func (h *TripHandlers) AcceptTrip(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_, _ = strconv.ParseInt(vars["tripId"], 10, 64)
	_, _ = strconv.ParseInt(vars["driverId"], 10, 64)

	err := h.service.AcceptTrip()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("Поездка принята"))
	if err != nil {
		return
	}
}

func (h *TripHandlers) StartTrip(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_, _ = strconv.ParseInt(vars["tripId"], 10, 64)

	err := h.service.StartTrip()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("Поездка началась"))
	if err != nil {
		return
	}
}

func (h *TripHandlers) EndTrip(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_, _ = strconv.ParseInt(vars["tripId"], 10, 64)

	err := h.service.EndTrip()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("Поездка завершена"))
	if err != nil {
		return
	}
}
