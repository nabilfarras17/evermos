package soldier

import (
	"github.com/evermos/war/pkg/weapon"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Soldier struct {
	ID         uuid.UUID  `json:"id"`
	IdentifyID string     `json:"identifyId"`
	Name       string     `json:"name"`
	Gun        weapon.Gun `json:"gun"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  *time.Time `json:"updatedAt"`
}
