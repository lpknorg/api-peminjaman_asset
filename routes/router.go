package routes

import (
    "github.com/gorilla/mux"
    "api-asset2/controllers"
    // "api-asset2/middleware"
)

func SetupRouter() *mux.Router {

    router := mux.NewRouter()

    router.HandleFunc("/login", controllers.Logins).Methods("POST")

    api := router.PathPrefix("/api").Subrouter()
    // api.Use(middleware.JwtAuthentication)
    api.HandleFunc("/permintaan", controllers.GetPermintaans).Methods("GET")
    api.HandleFunc("/permintaan/{id}", controllers.GetPermintaan).Methods("GET")
    api.HandleFunc("/permintaan", controllers.CreatePermintaan).Methods("POST")
    api.HandleFunc("/permintaan/{id}", controllers.UpdatePermintaan).Methods("PUT")
    api.HandleFunc("/permintaan/{id}", controllers.DeletePermintaan).Methods("DELETE")



    api.HandleFunc("/perlengkapan", controllers.GetPerlengkapans).Methods("GET")

    return api
}
