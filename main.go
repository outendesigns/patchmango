package main

import (
    "log"
    "flag"
    "fmt"
    "net/http"
    "net"
    "os"
    "patchmango/handlers"
)

// Get preferred outbound IP of this machine
func GetOutboundIP() net.IP {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    localAddr := conn.LocalAddr().(*net.UDPAddr)

    return localAddr.IP
}

func updateHelpUsage() {
    flag.Usage = func() {
        fmt.Fprintf(os.Stderr, "PatchMango API Server\n\n")
        fmt.Fprintf(os.Stderr, "  MIT License\n")
        fmt.Fprintf(os.Stderr, "  Copyright (c) 2025 Eric R Outen\n\n")
        fmt.Fprintf(os.Stderr, "This serves as a framework for an Embedded Device Updater using Go as a server-side API.\n\n")
        fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS]\n\n", os.Args[0])
        fmt.Fprintf(os.Stderr, "Options:\n")
        flag.PrintDefaults()

        fmt.Fprintf(os.Stderr, "\nAPI Endpoints:\n")
        fmt.Fprintf(os.Stderr, "  GET  /patchmango/api/versions         - List all available versions\n")
        fmt.Fprintf(os.Stderr, "  GET  /patchmango/api/current_version  - Get the current version\n")
        fmt.Fprintf(os.Stderr, "  GET  /patchmango/api/get_file         - Download a file\n")
        
        fmt.Fprintf(os.Stderr, "\nExample Usage:\n")
        fmt.Fprintf(os.Stderr, "  Start server:\n")
        fmt.Fprintf(os.Stderr, "    %s -port=8080 -verbose\n\n", os.Args[0])
        fmt.Fprintf(os.Stderr, "  Make a request to view all Versions:\n")
        fmt.Fprintf(os.Stderr, "    curl -H \"Authorization: Bearer your-secret-token\" \\\n")
        fmt.Fprintf(os.Stderr, "         \"http://localhost:8080/patchmango/api/versions\"\n\n")
        fmt.Fprintf(os.Stderr, "  Make a request to view the Current Version:\n")
        fmt.Fprintf(os.Stderr, "    curl -H \"Authorization: Bearer your-secret-token\" \\\n")
        fmt.Fprintf(os.Stderr, "         \"http://localhost:8080/patchmango/api/current_version\"\n\n")
        fmt.Fprintf(os.Stderr, "  Download and Execute a File:\n")
        fmt.Fprintf(os.Stderr, "    curl -H \"Authorization: Bearer your-secret-token\" \\\n")
        fmt.Fprintf(os.Stderr, "         \"http://localhost:8080/patchmango/api/get_file?file=filename\" | bash\n\n")
    }
}

func main() {
    // Define Flags
    port := flag.Int("port", 8080, "port to listen on")
    verbose := flag.Bool("verbose", false, "enable verbose output")

    // Update Help Usage
    updateHelpUsage()

    // Parse the flags
    flag.Parse()

    // Define API endpoints
    http.HandleFunc("/patchmango/api/versions", handlers.CheckAuthorization(handlers.GetVersions))
    http.HandleFunc("/patchmango/api/current_version", handlers.CheckAuthorization(handlers.GetCurrentVersion))
    http.HandleFunc("/patchmango/api/get_file", handlers.CheckAuthorization(handlers.GetFile))
    
    //port := config.Port
    //if p := os.Getenv("PORT"); p != "" {
    //    port = p
    //}

    localAddr := GetOutboundIP()
    if *verbose {
        log.Printf("Server starting on port %d", *port)
        log.Printf("IP address = %s", localAddr)
    }
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}