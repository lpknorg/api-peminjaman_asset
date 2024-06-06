package main

import (
    "log"
    "net/http"
    "api-asset2/routes"
    "api-asset2/config"    
)

func main() {
    config.InitDB()
    router := routes.SetupRouter()

    log.Println("Server started at :9990")
    log.Fatal(http.ListenAndServe(":9990", router))
}
