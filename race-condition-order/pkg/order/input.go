package order

type SaveOrderRequest struct {
	Created string                 `json:"created"`
	Total   float64                `json:"total"`
	Items   []SaveOrderItemRequest `json:"items"`
}

type SaveOrderItemRequest struct {
	SKU      string `json:"sku"`
	Quantity int    `json:"quantity"`
}
