package trip

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"log"
	"time"
)

func main() {
	fmt.Println("Hello, world!")
}

type Driver struct {
}

type Offer struct {
	Trip            *Trip
	Driver          *Driver
	ID              int64
	Accepted        bool
	Source          interface{}
	Type            interface{}
	DataContentType interface{}
	Data            interface{}
	Time            any
}

type Trip struct {
	ID              int64
	DriverID        int64
	Start           time.Time
	End             time.Time
	Status          string
	Source          interface{}
	Type            interface{}
	DataContentType interface{}
	Time            interface{}
	Data            interface{}
	IsStarted       bool
}

type DriverRepo struct {
	drivers []Driver
}

func (r DriverRepo) GetByID() {

	return
}

type Service struct {
	repo                Repository
	notificationService NotificationService
	offerRepo           OfferRepo
	driverRepo          DriverRepo
}

type Repository interface {
	StartTrip(ctx context.Context, tripID int64) error
	EndTrip(ctx context.Context, tripID int64) error
	GetTrip(ctx context.Context, tripID int64) (*Trip, error)
	UpdateTrip(ctx context.Context, trip *Trip) error
	DeleteTrip(ctx context.Context, tripID int64) error
	ListTrips(ctx context.Context, driverID int64) ([]*Trip, error)
	CreateTrip(ctx context.Context, trip *Trip) error
	GetNewTrips(ctx context.Context, lastSeenID int64) ([]*Trip, error)
}

type MockNotificationService struct {
	NotifiedTrips []int64
}

func (m *MockNotificationService) NotifyTripCreated(_ context.Context, trip *Trip) error {
	m.NotifiedTrips = append(m.NotifiedTrips, trip.ID)
	log.Printf("Уведомление о создании поездки: ID поездки - %d\n", trip.ID)
	return nil
}

type NotificationService interface {
	NotifyTripCreated(ctx context.Context, trip *Trip) error
}

type LocationService interface {
	GetAvailableDrivers() ([]*Driver, error)
}

type OfferRepo interface {
	Create(ctx context.Context, offer *Offer) error
	GetByID(ctx context.Context, offerID int64) (*Offer, error)
	Update(ctx context.Context, offer *Offer) error
}

type YourTripRepositoryImplementation struct {
	db    *sql.DB
	trips []*Trip
}

func (r *YourTripRepositoryImplementation) StartTrip(tripID int64) error {
	trip := r.findTripById(tripID)
	if trip == nil {
		return errors.New("cannot find trip with given ID")
	}

	trip.Status = "started"
	trip.Start = time.Now()

	err := r.updateTrip(trip)
	if err != nil {
		return errors.New("error setting trip to start") // укажите подходящее сообщение об ошибке здесь
	}

	return nil
}

func (r *YourTripRepositoryImplementation) findTripById(id int64) *Trip {
	for _, trip := range r.trips {
		if trip.ID == id {
			return trip
		}
	}

	return nil
}

func (r *YourTripRepositoryImplementation) updateTrip(updatedTrip *Trip) error {
	for i, trip := range r.trips {
		if trip.ID == updatedTrip.ID {
			r.trips[i] = updatedTrip
			return nil
		}
	}

	return errors.New("не удалось найти поездку для обновления")
}

func (r *YourTripRepositoryImplementation) ListTrips() ([]*Trip, error) {
	if len(r.trips) == 0 {
		return nil, errors.New("пока нет поездок")
	}
	return r.trips, nil
}

func (r *YourTripRepositoryImplementation) GetTrip(tripID int64) (*Trip, error) {
	for _, trip := range r.trips {
		if trip.ID == tripID {
			return trip, nil
		}
	}
	return nil, errors.New("поездка не найдена")
}

func (r *YourTripRepositoryImplementation) GetNewTrips() ([]*Trip, error) {
	var newTrips []*Trip

	for _, trip := range r.trips {
		if trip.Status == "Created" {
			newTrips = append(newTrips, trip)
		}
	}

	if len(newTrips) == 0 {
		return nil, errors.New("новые поездки не найдены")
	}

	return newTrips, nil
}

func (r *YourTripRepositoryImplementation) EndTrip(tripID int64) error {
	for _, trip := range r.trips {
		if trip.ID == tripID {
			trip.Status = "Ended"
			return nil
		}
	}

	return errors.New("не удалось найти поездку с указанным ID")
}

func (r *YourTripRepositoryImplementation) CreateTrip(newTrip *Trip) error {
	if newTrip.ID == 0 || newTrip.DriverID == 0 {
		return errors.New("неверные данные поездки")
	}

	r.trips = append(r.trips, newTrip)

	return nil
}

func (r *YourTripRepositoryImplementation) DeleteTrip(tripID int64) error {
	for index, trip := range r.trips {
		if trip.ID == tripID {
			r.trips = append(r.trips[:index], r.trips[index+1:]...)
			return nil
		}
	}
	return errors.New("не удалось найти поездку с указанным ID")
}

func NewTripService(repo Repository, notificationService *MockNotificationService, offerRepo *main.OfferRepository) (*Service, error) {
	if repo == nil {
		return nil, errors.New("repository is nil")
	}
	if notificationService == nil {
		return nil, errors.New("notification service is nil")
	}
	if offerRepo == nil {
		return nil, errors.New("offer repository is nil")
	}

	return &Service{
		repo:                repo,
		notificationService: notificationService,
	}, nil
}

func (s *Service) AcceptTrip() error {
	ctx := context.TODO()
	tripID := int64(0) // Convert tripID to int64
	trip, err := s.repo.GetTrip(ctx, tripID)
	if err != nil {
		return err
	}

	if trip == nil {
		return errors.New("поездка не найдена")
	}

	driverID := int64(0) // Convert driverID to int64
	if driverID == 0 {
		return errors.New("водитель не найден")
	}

	trip.Status = "Accepted"
	trip.DriverID = driverID

	err = s.repo.UpdateTrip(ctx, trip)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) StartTrip() error {
	ctx := context.TODO()
	tripID := int64(0) // Convert tripID to int64
	trip, err := s.repo.GetTrip(ctx, tripID)
	if err != nil {
		return err
	}

	if trip == nil {
		return errors.New("поездка не найдена")
	}

	trip.Status = "In progress"
	trip.Start = time.Now()

	err = s.repo.UpdateTrip(ctx, trip)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) EndTrip() error {
	ctx := context.TODO()
	tripID := int64(0) // Convert tripID to int64
	trip, err := s.repo.GetTrip(ctx, tripID)
	if err != nil {
		return err
	}

	if trip == nil {
		return errors.New("поездка не найдена")
	}

	trip.Status = "Ended"
	trip.End = time.Now()

	err = s.repo.UpdateTrip(ctx, trip)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) CreateTrip(ctx context.Context, newTrip *Trip) error {
	if newTrip == nil {
		return errors.New("требуются детали поездки")
	}

	newTrip.Status = "Created"
	newTrip.Start = time.Now()

	err := s.repo.CreateTrip(ctx, newTrip)
	if err != nil {
		return err
	}

	err = s.notificationService.NotifyTripCreated(ctx, newTrip)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) LongPollTrips(ctx context.Context, lastSeenID int64) ([]*Trip, error) {
	timeout := time.After(30 * time.Second)
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-timeout:
			return nil, nil
		case <-ticker.C:
			trips, err := s.repo.GetNewTrips(ctx, lastSeenID)
			if err != nil {
				return nil, err
			}
			if len(trips) > 0 {
				return trips, nil
			}
		}
	}
}

func (s *Service) OfferEndTrip(ctx context.Context, tripID int64) error {
	trip, err := s.repo.GetTrip(ctx, tripID)
	if err != nil {
		return err
	}

	driver := Driver{}
	offer := Offer{
		Trip:   trip,
		Driver: &driver,
		ID:     0,
	}
	err = s.offerRepo.Create(ctx, &offer)
	if err != nil {
		return err
	}

	timeout := time.After(30 * time.Second)
	ticker := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("ошибка контекста: %v", ctx.Err())
		case <-timeout:
			return fmt.Errorf("таймаут ожидания принятия предложения водителем")
		case <-ticker.C:
			offer, err := s.offerRepo.GetByID(ctx, offer.ID)
			if err != nil {
				return err
			}

			if offer != nil && offer.Accepted {
				trip.Status = "Завершено"
				return s.repo.UpdateTrip(ctx, trip)
			}
		}
	}
}

func (s *Service) AcceptDriverTrip() error {
	ctx := context.TODO()
	tripOfferID := int64(123) // Replace 123 with the actual tripOfferID value and convert it to int64
	tripOffer, err := s.offerRepo.GetByID(ctx, tripOfferID)
	if err != nil {
		return err
	}

	var driverID int64 // Declare the driverID variable
	if tripOffer == nil {
		return errors.New("предложение не найдено")
	}

	if driverID == 0 {
		return errors.New("водитель не найден")
	}

	if tripOffer.Accepted {
		return errors.New("предложение уже принято")
	}

	s.driverRepo.GetByID()
	if err != nil {
		return err
	}

	tripOffer.Accepted = true
	err = s.offerRepo.Update(ctx, tripOffer)
	if err != nil {
		return err
	}

	return nil
}
