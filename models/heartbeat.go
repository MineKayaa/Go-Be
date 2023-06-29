package models

import (
	"time"
)

type Heartbeat struct {
	DeviceID  int
	Device    Device
	UpdatedAt time.Time
}
