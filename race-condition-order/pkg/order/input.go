package order

import (
	"net/url"
)

type SaveOrderRequest struct {
	PhoneNumber string                 `json:"phoneNumber"`
	Created     string                 `json:"created"`
	Total       float64                `json:"total"`
	Items       []SaveOrderItemRequest `json:"items"`
}

func (s *SaveOrderRequest) validate() url.Values {
	errs := url.Values{}

	if s.PhoneNumber == "" {
		errs.Add("phoneNumber", "The phoneNumber field is required!")
	}

	if s.Created == "" {
		errs.Add("created", "The created field is required!")
	}

	if s.Total == 0 {
		errs.Add("total", "The total field is required!")
	}

	if len(s.Items) == 0 {
		errs.Add("items", "The items field is empty!")
	}

	for _, orderItemRequest := range s.Items {
		validated := orderItemRequest.validate()
		if len(validated) > 0 {
			return validated
		}
	}
	return errs
}

type SaveOrderItemRequest struct {
	SKU      string `json:"sku"`
	Quantity int    `json:"quantity"`
}

func (s *SaveOrderItemRequest) validate() url.Values {
	errs := url.Values{}
	if s.SKU == "" {
		errs.Add("sku", "The sku field is required!")
	}

	if s.Quantity == 0 {
		errs.Add("quantity", "The quantity field is required!")
	}
	return errs
}
