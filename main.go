package main

import (
    "log"
    "net/http"
    "os"
    "patchmango/handlers"
    "patchmango/config"
)

func main() {
    http.HandleFunc("/patchmango/api/versions", handlers.CheckAuthorization(handlers.GetVersions))
    http.HandleFunc("/patchmango/api/current_version", handlers.CheckAuthorization(handlers.GetCurrentVersion))
    http.HandleFunc("/patchmango/api/get_file", handlers.CheckAuthorization(handlers.GetFile))
    
    port := config.Port
    if p := os.Getenv("PORT"); p != "" {
        port = p
    }
    
    log.Printf("Server starting on port %s", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}