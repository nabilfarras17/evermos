package weapon

import (
	uuid "github.com/satori/go.uuid"
)

type Gun struct {
	ID           uuid.UUID  `json:"id"`
	SerialNumber string     `json:"serialNumber"`
	Name         string     `json:"name"`
	Magazines    []Magazine `json:"magazines"`
}

type Magazine struct {
	ID                 uuid.UUID `json:"id"`
	Bullets            []Bullet  `json:"_"`
	CurrentTotalBullet int       `json:"currentTotalBullet"`
	MaxTotalBullet     int       `json:"maxTotalBullet"`
}

type Bullet struct{}
