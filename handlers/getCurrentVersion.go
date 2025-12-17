package handlers

import (
    "encoding/json"
    "log"
    "net/http"
    "os"
    "patchmango/models"
    "patchmango/config"
)

func GetCurrentVersion(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    data, err := os.ReadFile(config.VersionsListFile)
    if err != nil {
        log.Printf("Error reading versions list: %v", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    var versionsList models.VersionsListStruct
    if err := json.Unmarshal(data, &versionsList); err != nil {
        log.Printf("Error parsing versions list: %v", err)
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    // Check if there are any versions
    if len(versionsList.Versions) == 0 {
        http.Error(w, "No versions available", http.StatusNotFound)
        return
    }

    currentVersion := versionsList.Versions[len(versionsList.Versions)-1]

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(currentVersion)
}