package handler

// Init Routes
func (rh *RootHandler) InitRoutes() {
	// Soldier Route
	rh.router.HandleFunc("/api/soldier/create", rh.soldierHandler.CreateSoldierHandler).Methods("POST")
	rh.router.HandleFunc("/api/soldier/{identifyId}", rh.soldierHandler.GetSoldierByIdentifyIDHandler).Methods("GET")
	rh.router.HandleFunc("/api/soldier/load-bullets/{identifyId}", rh.soldierHandler.LoadBulletsHandler).Methods("POST")
	rh.router.HandleFunc("/api/soldier/fire/{identifyId}", rh.soldierHandler.FireHandler).Methods("POST")
}
