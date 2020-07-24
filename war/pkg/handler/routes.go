package handler

// Init Routes
func (rh *RootHandler) InitRoutes() {
	// Soldier Route
	rh.router.HandleFunc("/api/soldier/create", rh.soldierHandler.CreateSoldierHandler).Methods("POST")
}
