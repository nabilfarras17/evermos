package soldier

import (
	"context"
	"github.com/evermos/war/pkg/util"
	"github.com/evermos/war/pkg/weapon"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	// log "github.com/sirupsen/logrus"
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
		ID:         uuid.NewV4(),
		IdentifyID: util.RandomString(10),
		Name:       request.Name,
		CreatedAt:  time.Now(),
	}
	gun, err := s.weaponService.CreateGun(ctx, request.GunName)
	if err != nil {
		return
	}
	soldier.Gun = gun
	return s.repository.InsertSoldier(ctx, soldier)
}

func (s *Service) GetSoldierByIdentifyID(ctx context.Context, input string) (res Soldier, err error) {
	soldier := s.repository.GetSoldierByIdentifyID(ctx, input)
	if soldier == nil {
		err = errors.Errorf("Soldier with identifyId: %v is not found!", input)
		return
	}
	res = *soldier
	return
}

func (s *Service) LoadBullets(ctx context.Context, identifyID string, request LoadBulletRequest) (res weapon.Gun, err error) {
	soldier := s.repository.GetSoldierByIdentifyID(ctx, identifyID)
	if soldier == nil {
		err = errors.Errorf("Soldier with identifyId: %v is not found!", identifyID)
		return
	}

	if soldier.Gun.IsVerified() {
		err = errors.Errorf("Gun with serialNumber: %v is already verified", soldier.Gun.SerialNumber)
		return
	}

	gun, err := s.weaponService.LoadBullets(ctx, soldier.Gun, request.Bullets)
	if err != nil {
		return
	}
	soldier.Gun = gun
	s.repository.UpdateSoldier(ctx, *soldier)
	res = gun
	return
}

func (s *Service) FireGun(ctx context.Context, identifyID string) (isVerified bool, err error) {
	soldier := s.repository.GetSoldierByIdentifyID(ctx, identifyID)
	if soldier == nil {
		err = errors.Errorf("Soldier with identifyId: %v is not found!", identifyID)
		return
	}
	for _, magazine := range soldier.Gun.Magazines {
		if magazine.IsVerified {
			isVerified = true
			return
		}
	}
	isVerified = false
	return
}
