package repository

import (
	"context"
	"database/sql"

	"github.com/Gitcodeindev/driverlocation_go/internal/model"
)

type DriverRepository interface {
	Create(ctx context.Context, driver *model.Driver) error
	Update(ctx context.Context, driver *model.Driver) error
	GetByID(ctx context.Context, id int64) (*model.Driver, error)
}

func NewDriverRepository(db *sql.DB) DriverRepository {
	return &driverRepositoryImpl{
		db: db,
	}
}

type driverRepositoryImpl struct {
	db *sql.DB
}

func (r *driverRepositoryImpl) Create(ctx context.Context, driver *model.Driver) error {
	query := `INSERT INTO drivers (name, license) VALUES ($1, $2)`
	_, err := r.db.ExecContext(ctx, query, driver.Name, driver.License)
	if err != nil {
		return err
	}
	return nil
}

func (r *driverRepositoryImpl) Update(ctx context.Context, driver *model.Driver) error {
	query := `UPDATE drivers SET name=$1, license=$2, available=$3, contact_info=$4, rating=$5 WHERE id=$6`
	_, err := r.db.ExecContext(ctx, query, driver.Name, driver.License, driver.Available, driver.ContactInfo, driver.Rating, driver.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *driverRepositoryImpl) GetByID(ctx context.Context, id int64) (*model.Driver, error) {
	query := `SELECT id, name, license, available, contact_info, rating FROM drivers WHERE id=$1`
	row := r.db.QueryRowContext(ctx, query, id)

	driver := &model.Driver{}
	err := row.Scan(&driver.ID, &driver.Name, &driver.License, &driver.Available, &driver.ContactInfo, &driver.Rating)
	if err != nil {
		return nil, err
	}
	return driver, nil
}
