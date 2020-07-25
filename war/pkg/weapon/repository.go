package weapon

import (
	"context"
	"github.com/pkg/errors"
)

var gunMap map[string]Gun

type Repository struct{}

func NewRepository() Repository {
	gunMap = make(map[string]Gun, 0)
	return Repository{}
}

func (r *Repository) InsertGun(ctx context.Context, gun Gun) (res Gun, err error) {
	if _, ok := gunMap[gun.SerialNumber]; ok {
		err = errors.Errorf("Gun with serialNumber: %v already exists!", gun.SerialNumber)
		return
	}
	gunMap[gun.SerialNumber] = gun
	return gun, nil
}

func (r *Repository) UpdateGun(ctx context.Context, gun Gun) (res Gun, err error) {
	if _, ok := gunMap[gun.SerialNumber]; !ok {
		err = errors.Errorf("Gun with serialNumber: %v not exists!", gun.SerialNumber)
		return
	}
	gunMap[gun.SerialNumber] = gun
	return gun, nil
}

func (r *Repository) GetGunBySerialNumber(ctx context.Context, serialNumber string) (res Gun, err error) {
	if gun, ok := gunMap[serialNumber]; ok {
		res = gun
		return
	} else {
		err = errors.Errorf("Gun with serialNumber: %v not exists!", serialNumber)
		return
	}
}
