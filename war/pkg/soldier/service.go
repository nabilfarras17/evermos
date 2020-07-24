package soldier

import (
	"context"
	"github.com/evermos/war/pkg/weapon"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	"time"
)

type Service struct {
	repository    Repository
	weaponService weapon.Service
}

func NewService(weaponService weapon.Service) Service {
	soldierRepository := NewRepository()
	return Service{
		repository:    soldierRepository,
		weaponService: weaponService,
	}
}

func (s *Service) CreateSoldier(ctx context.Context, request CreateSoldierRequest) (res Soldier, err error) {
	soldier := Soldier{
		ID:        uuid.NewV4(),
		PublicID:  "xxxx",
		Name:      request.Name,
		CreatedAt: time.Now(),
	}
	gun, err := s.weaponService.CreateGun(ctx, request.GunName)
	if err != nil {
		return
	}
	soldier.Gun = gun
	return s.repository.InsertSoldier(ctx, soldier)
}

func (s *Service) GetSoldierByPublicID(ctx context.Context, input string) (res Soldier, err error) {
	soldier := s.repository.GetSoldierByPublicID(ctx, input)
	if soldier == nil {
		err = errors.Errorf("Soldier with publicId: %v is not found!", input)
		return
	}
	res = *soldier
	return
}
