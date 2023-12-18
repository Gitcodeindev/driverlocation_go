package driver

import (
	"context"
	"errors"
)

type Driver struct {
	ID        int64
	Name      string
	License   string
	Available bool
	Location  string
}

type DriverRepository interface {
	Create(driver *Driver) error
	Update(driver *Driver) error
	GetByID(id int64) (*Driver, error)
	GetAll() ([]*Driver, error)
}

type LocationService interface {
	GetLocation(driverID int64) (string, error)
}

type DriverService struct {
	driverRepo      DriverRepository
	locationService LocationService
}

func (s *DriverService) RegisterDriver(ctx context.Context, driver *Driver) error {
	if driver.Name == "" || driver.License == "" {
		return errors.New("invalid driver data")
	}
	return s.driverRepo.Create(driver)
}

func (s *DriverService) UpdateDriver(ctx context.Context, driver *Driver) error {
	if driver.Name == "" || driver.License == "" {
		return errors.New("invalid driver data")
	}
	return s.driverRepo.Update(driver)
}

func (s *DriverService) GetDrivers(ctx context.Context) ([]*Driver, error) {
	drivers, err := s.driverRepo.GetAll()
	if err != nil {
		return nil, err
	}
	for i, driver := range drivers {
		location, err := s.locationService.GetLocation(driver.ID)
		if err != nil {
			return nil, err
		}
		drivers[i].Location = location
	}
	return drivers, nil
}

func (s *DriverService) StartTrip(ctx context.Context, driverID int64, tripID int64) error {
	driver, err := s.driverRepo.GetByID(driverID)
	if err != nil {
		return err
	}
	if !driver.Available {
		return errors.New("driver is not available to start a trip")
	}
	return nil
}
