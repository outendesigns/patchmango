package main

import (
    "log"
    "net/http"
    "os"
    "patchmango/handlers"
)

func main() {
    http.HandleFunc("/patchmango/api/versions", handlers.GetVersions)
    http.HandleFunc("/patchmango/api/current_version", handlers.GetCurrentVersion)
    http.HandleFunc("/patchmango/api/get_file", handlers.GetFile)
    
    port := "8080"
    if p := os.Getenv("PORT"); p != "" {
        port = p
    }
    
    log.Printf("Server starting on port %s", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}