# [PatchMango API Server]
MIT License
Copyright (c) 2025 Eric R Outen
This serves as a framework for an Embedded Device Updater using Go as a server-side API.
## Usage:
./patchmango [OPTIONS]
## Options:
### -port int
port to listen on (default 8080)
### -verbose
enable verbose output
## API Endpoints:
GET  /patchmango/api/versions         - List all available versions
GET  /patchmango/api/current_version  - Get the current version
GET  /patchmango/api/get_file         - Download a file
## Example Usage:
### Start server:
./patchmango -port=8080 -verbose</p>
### From the Remote Device...
### Make a request to view all Versions:
curl -H "Authorization: Bearer your-secret-token" \</p>
"http://localhost:8080/patchmango/api/versions"
### Make a request to view the Current Version:
curl -H "Authorization: Bearer your-secret-token" \
"http://localhost:8080/patchmango/api/current_version"
### Download and Execute a File:
curl -H "Authorization: Bearer your-secret-token" \
"http://localhost:8080/patchmango/api/get_file?file=filename" | bash