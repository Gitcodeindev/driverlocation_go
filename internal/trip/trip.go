package trip

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type Driver struct {
	// Определение полей
}

type Offer struct {
	Trip     *Trip   // Добавьте поле Trip
	Driver   *Driver // Добавьте поле Driver
	ID       int64   // Добавьте поле ID
	Accepted bool    // Добавьте поле Accepted
}

type Trip struct {
	ID       int64
	DriverID int64
	Start    time.Time
	End      time.Time
	Status   string
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

type TripService struct {
	repo                TripRepository
	notificationService NotificationService
	offerRepo           OfferRepo
	tripRepo            TripRepository
}

func (s *TripService) CreateTrip(ctx context.Context, newTrip *Trip) error {
	if newTrip == nil {
		return errors.New("trip details are required")
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
	defer ticker.Stop() // Остановка таймера при выходе из функции

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
	trip, err := s.tripRepo.GetTrip(ctx, tripID)
	if err != nil {
		return err
	}

	driver := Driver{} // Объявляем переменную driver
	offer := Offer{
		Trip:   trip,
		Driver: &driver, // Передаем указатель на driver
		ID:     0,       // Добавьте запятую после ID
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
			return fmt.Errorf("context error: %v", ctx.Err())
		case <-timeout:
			return fmt.Errorf("timeout waiting for driver to accept offer")
		case <-ticker.C:
			offer, err := s.offerRepo.GetByID(ctx, offer.ID)
			if err != nil {
				return err
			}

			if offer != nil && offer.Accepted {
				trip.Status = "Completed"
				return s.repo.UpdateTrip(ctx, trip)
			}
		}
	}
}
