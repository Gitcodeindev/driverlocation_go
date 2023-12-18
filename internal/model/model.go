package model

import "time"

// представляет сущность водителя в системе
type Driver struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	License     string    `json:"license"`
	Available   bool      `json:"available"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	ContactInfo string    `json:"contact_info"`
	Rating      float64   `json:"rating"`
}

// создает новый экземпляр водителя с начальными значениями
func NewDriver(name, license string) *Driver {
    return &Driver{
        Name:      name,
        License:   license,
        Available: true,
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }
}

// обновляет данные водителя
func (d *Driver) Update(name, license, contactInfo string, rating float64, available bool) {
	d.Name = name
	d.License = license
	d.Available = available
	d.ContactInfo = contactInfo
	d.Rating = rating
	d.UpdatedAt = time.Now()
}

