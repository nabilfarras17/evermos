package handler

// Init Routes
func (rh *RootHandler) InitRoutes() {
	// Order Route
	rh.router.HandleFunc("/api/order/save", rh.orderHandler.SaveOrder).Methods("POST")
	rh.router.HandleFunc("/api/order/{publicID}", rh.orderHandler.GetOrderByPublicID).Methods("GET")
	// Product Route
	rh.router.HandleFunc("/api/product/{sku}", rh.pcatalogueHandler.ResolveProductBySKU).Methods("GET")
}
