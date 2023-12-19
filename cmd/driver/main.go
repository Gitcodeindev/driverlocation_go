package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Gitcodeindev/driverlocation_go/internal/location"
	"github.com/Gitcodeindev/driverlocation_go/internal/trip"
	"github.com/gorilla/mux"
)

type LocationService struct {
}

func NewLocationService(locationRepo *LocationRepository) *LocationService {
	return &LocationService{}
}

type Location struct {
	ID        int64
	Latitude  float64
	Longitude float64
}

type LocationRepository struct {
	db *sql.DB
}

func (repo *LocationRepository) GetLocation(id int64) (*Location, error) {
	location1 := &Location{}
	query := "SELECT * FROM locations WHERE id=$1"
	err := repo.db.QueryRow(query, id).Scan(&location1.ID, &location1.Latitude, &location1.Longitude)
	if err != nil {
		return nil, err
	}
	return location1, nil
}

func (repo *LocationRepository) UpdateLocation(loc *Location) error {
	query := "UPDATE locations SET latitude=$1, longitude=$2 WHERE id=$3"
	_, err := repo.db.Exec(query, loc.Latitude, loc.Longitude, loc.ID)
	if err != nil {
		return err
	}
	return nil
}

type yourTripRepositoryImplementation struct {
	db *sql.DB
}

func (r *yourTripRepositoryImplementation) StartTrip(ctx context.Context, tripID int64) error {
	trip1, err := r.GetTrip(ctx, tripID)
	if err != nil {
		return err
	}

	if trip1.IsStarted {
		return errors.New("поездка уже началась")
	}

	trip1.IsStarted = true
	err = r.UpdateTrip(ctx, trip1)
	if err != nil {
		return err
	}

	return nil
}

func (r *yourTripRepositoryImplementation) CreateTrip(ctx context.Context, t *trip.Trip) error {
	query := `INSERT INTO trips (id, source, type, datacontenttype, time, data) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, t.ID, t.Source, t.Type, t.DataContentType, time.Now(), t.Data)
	if err != nil {
		return fmt.Errorf("не удалось создать поездку: %w", err)
	}
	return nil
}

func (r *yourTripRepositoryImplementation) DeleteTrip(ctx context.Context, tripID int64) error {
	query := `DELETE FROM trips WHERE id=$1`

	result, err := r.db.ExecContext(ctx, query, tripID)
	if err != nil {
		return fmt.Errorf("ошибка при удалении поездки: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("ошибка при проверки количества удаленных строк: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("не найдена поездка с id: %v", tripID)
	}

	return nil
}

func (r *yourTripRepositoryImplementation) EndTrip(ctx context.Context, tripID int64) error {
	query := `UPDATE trips SET status='Ended' WHERE id=$1`

	result, err := r.db.ExecContext(ctx, query, tripID)
	if err != nil {
		return fmt.Errorf("ошибка при завершении поездки: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("ошибка при проверки количества обновленных строк: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("не найдена поездка с id: %v", tripID)
	}

	return nil
}

func (r *yourTripRepositoryImplementation) GetNewTrips(ctx context.Context, lastSeenID int64) ([]*trip.Trip, error) {
	query := `SELECT * FROM trips WHERE id > $1`

	rows, err := r.db.QueryContext(ctx, query, lastSeenID)
	if err != nil {
		return nil, fmt.Errorf("ошибка при выполнении запроса к базе данных: %w", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var trips []*trip.Trip
	for rows.Next() {
		var t trip.Trip
		if err := rows.Scan(&t); err != nil {
			return nil, fmt.Errorf("ошибка при сканировании строки из базы данных: %w", err)
		}
		trips = append(trips, &t)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка в результате запроса: %w", err)
	}

	return trips, nil
}

func (r *yourTripRepositoryImplementation) GetTrip(ctx context.Context, tripID int64) (*trip.Trip, error) {
	query := `SELECT * FROM trips WHERE id=$1`
	row := r.db.QueryRowContext(ctx, query, tripID)
	t := &trip.Trip{}
	err := row.Scan(&t.ID, &t.Source, &t.Type, &t.DataContentType, &t.Time, &t.Data) // здесь вы должны просканировать все поля вашей структуры trip
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (r *yourTripRepositoryImplementation) ListTrips(ctx context.Context, driverID int64) ([]*trip.Trip, error) {
	query := `SELECT * FROM trips WHERE driverID = $1`

	rows, err := r.db.QueryContext(ctx, query, driverID)
	if err != nil {
		return nil, fmt.Errorf("ошибка при выполнении запроса к базе данных: %w", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
		}
	}(rows)

	var trips []*trip.Trip
	for rows.Next() {
		var t trip.Trip
		if err := rows.Scan(&t); err != nil {
			return nil, fmt.Errorf("ошибка при сканировании строки из базы данных: %w", err)
		}
		trips = append(trips, &t)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка в результате запроса: %w", err)
	}

	return trips, nil
}

func (r *yourTripRepositoryImplementation) UpdateTrip(ctx context.Context, trip *trip.Trip) error {
	query := `UPDATE trips SET driverID=$1, start=$2, end=$3, status=$4 WHERE id=$5`

	_, err := r.db.ExecContext(ctx, query, trip.DriverID, trip.Start, trip.End, trip.Status, trip.ID)
	if err != nil {
		return fmt.Errorf("ошибка при обновлении поездки: %w", err)
	}

	return nil
}

func NewTripRepository() trip.Repository {
	return &yourTripRepositoryImplementation{}
}

func main() {
	db, err := sql.Open("postgres", "user=yourUserName dbname=yourDbName sslmode=yourSslMode")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
		}
	}(db)

	offerRepo := NewOfferRepository()
	tripRepo := NewTripRepository()
	notificationService := &trip.MockNotificationService{}
	tripService, err := trip.NewTripService(tripRepo, notificationService, offerRepo)
	if err != nil {
		log.Fatal(err)
	}
	locationRepo := &LocationRepository{}
	Service := location.NewLocationService(locationRepo)

	r := mux.NewRouter()

	tripRoutes := r.PathPrefix("/trips").Subrouter()
	tripRoutes.HandleFunc("/{tripId}/accept", acceptTripHandler(tripService)).Methods("POST")
	tripRoutes.HandleFunc("/{tripId}/start", startTripHandler(tripService)).Methods("POST")
	tripRoutes.HandleFunc("/{tripId}/end", endTripHandler(tripService)).Methods("POST")

	driverRoutes := r.PathPrefix("/driver").Subrouter()
	driverRoutes.HandleFunc("/startTrip", startDriverTripHandler(tripService)).Methods("POST")
	driverRoutes.HandleFunc("/endTrip", endDriverTripHandler(tripService)).Methods("POST")
	driverRoutes.HandleFunc("/acceptTrip", acceptDriverTripHandler(tripService)).Methods("POST")

	r.HandleFunc("/trips", getTripsHandler(Service)).Methods("GET")

	log.Println("Запуск сервера на порту:", os.Getenv("DRIVER_PORT"))
	if err := http.ListenAndServe(":"+os.Getenv("DRIVER_PORT"), r); err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
}

type OfferRepository struct {
	db *sql.DB
}

func NewOfferRepository() *OfferRepository {
	return &OfferRepository{}
}

func startTripHandler(tripService *trip.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		_, err := strconv.ParseInt(vars["идВодителя"], 10, 64)
		if err != nil {
			http.Error(w, "Неверный ID водителя", http.StatusBadRequest)
			return
		}

		err = tripService.StartTrip()
		if err != nil {
			http.Error(w, "Не удалось начать поездку", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte("Поездка успешно начата!"))
		if err != nil {
			http.Error(w, "Ошибка отправки ответа", http.StatusInternalServerError)
			return
		}
	}
}

func endTripHandler(tripService *trip.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		tripIdStr, ok := vars["tripId"]
		if !ok {
			http.Error(w, "Необходим параметр tripId", http.StatusBadRequest)
			return
		}

		_, err := strconv.ParseInt(tripIdStr, 10, 64)
		if err != nil {
			http.Error(w, "Параметр tripId должен быть числом", http.StatusBadRequest)
			return
		}

		err = tripService.EndTrip()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func acceptTripHandler(tripService *trip.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		tripIdStr, ok := vars["tripId"]
		if !ok {
			http.Error(w, "Необходим параметр tripId", http.StatusBadRequest)
			return
		}

		_, err := strconv.ParseInt(tripIdStr, 10, 64)
		if err != nil {
			http.Error(w, "Параметр tripId должен быть числом", http.StatusBadRequest)
			return
		}

		var requestBody struct {
			DriverId int64 `json:"driverId"`
		}

		err = json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			http.Error(w, "Не смогли декодировать тело запроса: "+err.Error(), http.StatusBadRequest)
			return
		}

		err = tripService.AcceptTrip()
		if err != nil {
			// Если возникла ошибка, отправляем сообщение об ошибке
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func startDriverTripHandler(tripService *trip.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		tripIdStr, ok := vars["tripId"]
		if !ok {
			http.Error(w, "Необходим параметр tripId", http.StatusBadRequest)
			return
		}

		_, err := strconv.ParseInt(tripIdStr, 10, 64)
		if err != nil {
			http.Error(w, "Параметр tripId должен быть числом", http.StatusBadRequest)
			return
		}

		var requestBody struct {
			DriverId int64 `json:"driverId"`
		}

		err = json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			http.Error(w, "Не смогли декодировать тело запроса: "+err.Error(), http.StatusBadRequest)
			return
		}

		err = tripService.StartTrip()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func endDriverTripHandler(tripService *trip.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		tripIdStr, ok := vars["tripId"]
		if !ok {
			http.Error(w, "Необходим параметр tripId", http.StatusBadRequest)
			return
		}

		_, err := strconv.ParseInt(tripIdStr, 10, 64)
		if err != nil {
			http.Error(w, "Параметр tripId должен быть числом", http.StatusBadRequest)
			return
		}

		var requestBody struct {
			DriverId int64 `json:"driverId"`
		}

		err = json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			http.Error(w, "Не смогли декодировать тело запроса: "+err.Error(), http.StatusBadRequest)
			return
		}

		err = tripService.EndTrip()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func acceptDriverTripHandler(tripService *trip.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		tripIdStr, ok := vars["tripId"]
		if !ok {
			http.Error(w, "Необходим параметр tripId", http.StatusBadRequest)
			return
		}

		_, err := strconv.ParseInt(tripIdStr, 10, 64)
		if err != nil {
			http.Error(w, "Параметр tripId должен быть числом", http.StatusBadRequest)
			return
		}

		var requestBody struct {
			DriverId int64 `json:"driverId"`
		}

		err = json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			http.Error(w, "Не смогли декодировать тело запроса: "+err.Error(), http.StatusBadRequest)
			return
		}

		err = tripService.AcceptDriverTrip()
		if err != nil {
			return
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func getTripsHandler(locationService *location.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Call the location service to get all trips
		trips, err := locationService.GetAllTrips()

		// If there's an error fetching the trips
		if err != nil {
			// Send a 500 Internal Server Error response
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Convert the trips to JSON
		data, err := json.Marshal(trips)
		if err != nil {
			// Send a 500 Internal Server Error response
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Write the JSON trips data to the response
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(data)
		if err != nil {
			// Log the error or handle it as you need
			log.Printf("Failed to write response: %v", err)
		}
	}
}
