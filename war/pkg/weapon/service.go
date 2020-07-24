package weapon

import (
	"context"
	"github.com/satori/go.uuid"
)

type Service struct {
	repository Repository
}

func NewService() Service {
	weaponRepository := NewRepository()
	return Service{
		repository: weaponRepository,
	}
}

func (s *Service) CreateGun(ctx context.Context, name string) (res Gun, err error) {
	input := Gun{
		ID:           uuid.NewV4(),
		Name:         name,
		SerialNumber: "xxx",
	}
	return s.repository.InsertGun(ctx, input)
}
