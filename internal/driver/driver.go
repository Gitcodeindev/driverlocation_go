package driver

import (
	"errors"
)

type Driver struct {
	ID        int64
	Name      string
	License   string
	Available bool
	Location  string
}

type Repository interface {
	Create(driver *Driver) error
	Update(driver *Driver) error
	GetByID(id int64) (*Driver, error)
	GetAll() ([]*Driver, error)
}

type LocationService interface {
	GetLocation(driverID int64) (string, error)
}

type Service struct {
	repo       Repository
	locService LocationService
}

func (s *Service) RegisterDriver(driver *Driver) error {
	if driver.Name == "" || driver.License == "" {
		return errors.New("недопустимые данные водителя")
	}
	return s.repo.Create(driver)
}

func (s *Service) UpdateDriver(driver *Driver) error {
	if driver.Name == "" || driver.License == "" {
		return errors.New("недопустимые данные водителя")
	}
	return s.repo.Update(driver)
}

func (s *Service) GetDrivers() ([]*Driver, error) {
	drivers, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	for i, driver := range drivers {
		location, err := s.locService.GetLocation(driver.ID)
		if err != nil {
			return nil, err
		}
		drivers[i].Location = location
	}
	return drivers, nil
}

func (s *Service) StartTrip(driverID int64) error {
	driver, err := s.repo.GetByID(driverID)
	if err != nil {
		return err
	}
	if !driver.Available {
		return errors.New("водитель не доступен для начала поездки")
	}

	driver.Available = false

	if err := s.repo.Update(driver); err != nil {
		return err
	}

	return nil
}