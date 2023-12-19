package api

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Gitcodeindev/driverlocation_go/internal/trip"
	"github.com/gorilla/mux"
)

type TripHandlers struct {
	service *trip.TripService
}

type TripService struct {
}

func NewTripHandlers(service *trip.TripService) *TripHandlers {
	return &TripHandlers{service: service}
}

func (s *TripService) AcceptTrip(ctx context.Context, tripID, driverID int64) error {
	return nil
}

func (s *TripService) StartTrip(ctx context.Context, tripID int64) error {
	return nil
}

func (s *TripService) EndTrip(ctx context.Context, tripID int64) error {
	return nil
}

func (h *TripHandlers) AcceptTrip(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tripID, _ := strconv.ParseInt(vars["tripId"], 10, 64)
	driverID, _ := strconv.ParseInt(vars["driverId"], 10, 64)

	err := h.service.AcceptTrip(tripID, driverID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Поездка принята"))
}

func (h *TripHandlers) StartTrip(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tripID, _ := strconv.ParseInt(vars["tripId"], 10, 64)

	err := h.service.StartTrip(tripID, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Поездка началась"))
}

func (h *TripHandlers) EndTrip(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tripID, _ := strconv.ParseInt(vars["tripId"], 10, 64)

	err := h.service.EndTrip(tripID, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Поездка завершена"))
}
