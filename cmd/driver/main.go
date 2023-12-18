package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Gitcodeindev/driverlocation_go/internal/location"
	"github.com/Gitcodeindev/driverlocation_go/internal/trip"
	"github.com/gorilla/mux"
)

type LocationRepository struct {
}

func (r *LocationRepository) GetLocation(id int64) (*location.Location, error) {
	return nil, nil
}

type MockNotificationService struct{}

func (m *MockNotificationService) NotifyTripCreated(trip *trip.Trip) error {
	return nil
}

func main() {
	tripRepo := trip.NewRepository()
	notificationService := &MockNotificationService{}
	tripService := trip.NewTripService(tripRepo, notificationService)

	r := mux.NewRouter()

	r.HandleFunc("/trips/{tripId}/accept", acceptTripHandler(tripService)).Methods("POST")
	r.HandleFunc("/trips/{tripId}/start", startTripHandler(tripService)).Methods("POST")
	r.HandleFunc("/trips/{tripId}/end", endTripHandler(tripService)).Methods("POST")

	locationRepo := &LocationRepository{}
	locationService := location.NewLocationService(locationRepo)

	r.HandleFunc("/trips", getTripsHandler(locationService)).Methods("GET")

	r.HandleFunc("/driver/startTrip", startDriverTripHandler(tripService)).Methods("POST")
	r.HandleFunc("/driver/endTrip", endDriverTripHandler(tripService)).Methods("POST")
	r.HandleFunc("/driver/acceptTrip", acceptDriverTripHandler(tripService)).Methods("POST")

	log.Println("Starting server on port:", os.Getenv("DRIVER_PORT"))
	if err := http.ListenAndServe(":"+os.Getenv("DRIVER_PORT"), r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
func startTripHandler(tripService *trip.TripService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func endTripHandler(tripService *trip.TripService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func acceptTripHandler(tripService *trip.TripService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func startDriverTripHandler(tripService *trip.TripService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func endDriverTripHandler(tripService *trip.TripService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func acceptDriverTripHandler(tripService *trip.TripService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func getTripsHandler(locationService *location.LocationService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
