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

// AcceptTrip - метод для принятия поездки
func (s *TripService) AcceptTrip(ctx context.Context, tripID, driverID int64) error {
	// реализация принятия поездки
	return nil
}

// StartTrip - метод для начала поездки
func (s *TripService) StartTrip(ctx context.Context, tripID int64) error {
	// реализация начала поездки
	return nil
}

// EndTrip - метод для завершения поездки
func (s *TripService) EndTrip(ctx context.Context, tripID int64) error {
	// реализация завершения поездки
	return nil
}

func NewTripHandlers(service *trip.TripService) *TripHandlers {
	return &TripHandlers{service: service}
}

func (h *TripHandlers) AcceptTrip(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tripID, _ := strconv.ParseInt(vars["tripId"], 10, 64)
	driverID, _ := strconv.ParseInt(vars["driverId"], 10, 64)

	err := h.service.AcceptTrip(r.Context(), tripID, driverID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Trip accepted"))
}

func (h *TripHandlers) StartTrip(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tripID, _ := strconv.ParseInt(vars["tripId"], 10, 64)

	err := h.service.StartTrip(r.Context(), tripID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Trip started"))
}

func (h *TripHandlers) EndTrip(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tripID, _ := strconv.ParseInt(vars["tripId"], 10, 64)

	err := h.service.EndTrip(r.Context(), tripID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Trip ended"))
}
