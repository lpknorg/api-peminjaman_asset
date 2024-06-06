package routes

import (
    "github.com/gorilla/mux"
    "api-asset2/controllers"
)

func SetupRouter() *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/permintaan", controllers.GetPermintaans).Methods("GET")
    router.HandleFunc("/permintaan/{id}", controllers.GetPermintaan).Methods("GET")
    router.HandleFunc("/permintaan", controllers.CreatePermintaan).Methods("POST")
    router.HandleFunc("/permintaan/{id}", controllers.UpdatePermintaan).Methods("PUT")
    router.HandleFunc("/permintaan/{id}", controllers.DeletePermintaan).Methods("DELETE")

    router.HandleFunc("/perlengkapan", controllers.GetPerlengkapans).Methods("GET")

    return router
}
