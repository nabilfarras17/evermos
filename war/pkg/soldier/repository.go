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
	if _, ok := soldierMap[input.IdentifyID]; ok {
		err = errors.Errorf("Soldier with identify: %v is already exist!", input.IdentifyID)
		return
	}
	soldierMap[input.IdentifyID] = input
	return input, nil
}

func (r *Repository) GetSoldierByIdentifyID(ctx context.Context, input string) *Soldier {
	if soldier, ok := soldierMap[input]; ok {
		return &soldier
	}
	return nil
}

func (r *Repository) UpdateSoldier(ctx context.Context, input Soldier) (res Soldier, err error) {
	if _, ok := soldierMap[input.IdentifyID]; ok {
		soldierMap[input.IdentifyID] = input
		res = input
		return
	} else {
		err = errors.Errorf("Soldier with identify: %v is not exist!", input.IdentifyID)
		return
	}
}
