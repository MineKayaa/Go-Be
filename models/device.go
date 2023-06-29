package models

import "time"

type Device struct {
	ID            uint
	Latitude      float64
	Longitude     float64
	TechnicianID  int
	IsAlertIssued bool
	Technician    Technician
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
