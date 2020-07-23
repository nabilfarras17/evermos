package order

type OrderHandler struct {
	service Service
}

func NewOrderHandler(service Service) OrderHandler {
	return OrderHandler{}
}
