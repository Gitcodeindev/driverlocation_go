package location

import "github.com/Gitcodeindev/driverlocation_go/cmd/driver"

type Location struct {
	ID        int64   `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	DriverID  int64   `json:"driver_id"`
	Timestamp int64   `json:"timestamp"`
}

type LocationRepository interface {
	GetLocation(driverID int64) (*Location, error)
	UpdateLocation(location *Location) error
}

type LocationService struct {
	Repo LocationRepository
}

func NewLocationService(repo *main.LocationRepository) *LocationService {
	return &LocationService{
		Repo: repo,
	}
}
