package trip

import (
	"context"
	"errors"
	"fmt"
	"time"

	
)

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

type TripService struct {
	repo                TripRepository
	notificationService NotificationService
	offerRepo           OfferRepo
}

type TripRepository interface {
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

func (m *MockNotificationService) NotifyTripCreated(ctx context.Context, trip *Trip) error {
	// здесь должна быть реализация метода
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
}

func NewTripRepository() *yourTripRepositoryImplementation {

	return &yourTripRepositoryImplementation{}
}

type yourTripRepositoryImplementation struct {
}

func (r *yourTripRepositoryImplementation) StartTrip(ctx context.Context, tripID int64) error {
	return nil
}

func (r *yourTripRepositoryImplementation) ListTrips(ctx context.Context, driverID int64) ([]*Trip, error) {
	return nil, nil
}

func (r *yourTripRepositoryImplementation) GetTrip(ctx context.Context, tripID int64) (*Trip, error) {
	return nil, nil
}

func (r *yourTripRepositoryImplementation) GetNewTrips(ctx context.Context, lastSeenID int64) ([]*Trip, error) {
	return nil, nil
}

func (r *yourTripRepositoryImplementation) EndTrip(ctx context.Context, tripID int64) error {
	return nil
}

func (r *yourTripRepositoryImplementation) CreateTrip(ctx context.Context, trip *Trip) error {
	return nil
}

func (r *yourTripRepositoryImplementation) DeleteTrip(ctx context.Context, tripID int64) error {
	return nil
}

func NewTripService(repo TripRepository, notificationService *MockNotificationService, offerRepo OfferRepo) *TripService {
	return &TripService{
		repo:                repo,
		notificationService: notificationService,
		offerRepo:           offerRepo,
		}
}

func (s *TripService) AcceptTrip(tripID, driverID int64) error {
	return nil
}

func (s *TripService) StartTrip(tripID int64, id int64) error {
	return nil
}

func (s *TripService) EndTrip(tripID int64, id int64) error {
	return nil
}

func (s *TripService) CreateTrip(ctx context.Context, newTrip *Trip) error {
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

func (s *TripService) LongPollTrips(ctx context.Context, lastSeenID int64) ([]*Trip, error) {
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

func (s *TripService) OfferEndTrip(ctx context.Context, tripID int64) error {
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

func (s *TripService) AcceptDriverTrip(tripId int64, driverId int64) {
	return
}
