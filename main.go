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

    log.Println("Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
