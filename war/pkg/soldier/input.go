package soldier

import (
	"net/url"
)

type CreateSoldierRequest struct {
	Name    string `json:"name"`
	GunName string `json:"gunName"`
}

func (c *CreateSoldierRequest) validate() url.Values {
	errs := url.Values{}
	if c.Name == "" {
		errs.Add("name", "The name field is required!")
	}

	if c.GunName == "" {
		errs.Add("gunName", "The gunName field is required!")
	}
	return errs
}

type LoadBulletRequest struct {
	Bullets int `json:"bullets"`
}

func (l *LoadBulletRequest) validate() url.Values {
	errs := url.Values{}
	if l.Bullets == 0 {
		errs.Add("bullets", "The bullets field is required!")
	}
	return errs
}
