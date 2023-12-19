package repository

import (
	"context"
	_ "database/sql"

	"github.com/Gitcodeindev/driverlocation_go/internal/model"
)

type DriverRepository interface {
	Create(ctx context.Context, driver *model.Driver) error
	Update(ctx context.Context, driver *model.Driver) error
	GetByID(ctx context.Context, id int64) (*model.Driver, error)
}
