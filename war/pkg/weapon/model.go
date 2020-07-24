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

func (g *Gun) IsVerified() bool {
	for _, magazine := range g.Magazines {
		if magazine.IsVerified {
			return true
		}
		return false
	}
	return false
}

type Magazine struct {
	ID             uuid.UUID `json:"id"`
	GunID          uuid.UUID `json:"gunId"`
	IsVerified     bool      `json:"isVerified"`
	Bullets        []Bullet  `json:"-"`
	MaxTotalBullet int       `json:"maxTotalBullet"`
}

type Bullet struct{}
