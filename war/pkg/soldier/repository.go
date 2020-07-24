package soldier

import (
	"context"
	"github.com/pkg/errors"
)

var soldierMap map[string]Soldier

type Repository struct{}

func NewRepository() Repository {
	soldierMap = make(map[string]Soldier, 0)
	return Repository{}
}

func (r *Repository) InsertSoldier(ctx context.Context, input Soldier) (res Soldier, err error) {
	if _, ok := soldierMap[input.PublicID]; ok {
		err = errors.Errorf("Soldier with publicID: %v", input.PublicID)
		return
	}
	soldierMap[input.PublicID] = input
	return input, nil
}

func (r *Repository) GetSoldierByPublicID(ctx context.Context, input string) *Soldier {
	if soldier, ok := soldierMap[input]; ok {
		return &soldier
	}
	return nil
}
