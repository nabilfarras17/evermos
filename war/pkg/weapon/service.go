package weapon

import (
	"context"
	"github.com/evermos/war/config"
	"github.com/evermos/war/pkg/util"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	// log "github.com/sirupsen/logrus"
	"math/rand"
)

type Service struct {
	config     *config.Config
	repository Repository
}

func NewService(conf *config.Config) Service {
	weaponRepository := NewRepository()
	return Service{
		config:     conf,
		repository: weaponRepository,
	}
}

func (s *Service) CreateGun(ctx context.Context, name string) (res Gun, err error) {
	input := Gun{
		ID:           uuid.NewV4(),
		Name:         name,
		SerialNumber: util.RandomString(12),
	}
	magazines := make([]Magazine, 0)
	magazines = append(magazines, Magazine{
		ID:             uuid.NewV4(),
		GunID:          input.ID,
		IsVerified:     false,
		MaxTotalBullet: s.config.MaxThresholdBullet,
	})
	input.Magazines = magazines
	return s.repository.InsertGun(ctx, input)
}

func (s *Service) GetGunBySerialNumber(ctx context.Context, serialNumber string) (res Gun, err error) {
	return s.repository.GetGunBySerialNumber(ctx, serialNumber)
}

func (s *Service) LoadBullets(ctx context.Context, gun Gun, bulletCount int) (res Gun, err error) {
	var (
		magazine     Magazine
		newMagazines []Magazine
	)

	if len(gun.Magazines) == 0 {
		err = errors.New("magazine in gun is empty!")
		return
	} else if len(gun.Magazines) == 1 {
		magazine = gun.Magazines[0]
	} else {
		randomPickMagazine := rand.Intn(len(gun.Magazines))
		magazine = gun.Magazines[randomPickMagazine]
	}

	for i := 0; i < bulletCount; i++ {
		if len(magazine.Bullets) < magazine.MaxTotalBullet {
			magazine.Bullets = append(magazine.Bullets, Bullet{})
		}
	}

	if len(magazine.Bullets) == magazine.MaxTotalBullet {
		magazine.IsVerified = true
	}

	for _, newMagazine := range gun.Magazines {
		if newMagazine.ID == magazine.ID {
			newMagazines = append(newMagazines, magazine)
			continue
		} else {
			newMagazines = append(newMagazines, newMagazine)
		}
	}
	gun.Magazines = newMagazines
	s.repository.UpdateGun(ctx, gun)
	res = gun
	return
}
