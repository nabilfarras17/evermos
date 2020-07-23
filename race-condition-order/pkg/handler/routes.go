package handler

// Init Routes
func (rh *RootHandler) InitRoutes() {
	rh.router.HandleFunc("/order/save", rh.orderHandler.SaveOrder).Methods("POST")
}
