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
