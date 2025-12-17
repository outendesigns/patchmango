package handlers

import (
    "fmt"
    "io"
    "log"
    "net/http"
    "os"
    "path/filepath"
    "patchmango/config"
)

func GetFile(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    filename := r.URL.Query().Get("file")
    if filename == "" {
        http.Error(w, "Missing file parameter", http.StatusBadRequest)
        return
    }

    filename = filepath.Base(filename)

    filePath := filepath.Join(config.FilesDir, filename)

    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        http.Error(w, "File not found", http.StatusNotFound)
        return
    }

    file, err := os.Open(filePath)
    if err != nil {
        log.Printf("Error opening file %s: %v", filename, err)
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }
    defer file.Close()
    
    w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
    w.Header().Set("Content-Type", "application/octet-stream")
    
    io.Copy(w, file)
}