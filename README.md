# [PatchMango API Server]
<p>MIT License</p>
<p>Copyright (c) 2025 Eric R Outen</p>
<p>This serves as a framework for an Embedded Device Updater using Go as a server-side API.</p>
## Usage:
<p>./patchmango [OPTIONS]</p>
## Options:
### -port int
<p>port to listen on (default 8080)</p>
### -verbose
<p>enable verbose output</p>
## API Endpoints:
  		<p>GET  /patchmango/api/versions         - List all available versions<p>
  		<p>GET  /patchmango/api/current_version  - Get the current version</p>
  		<p>GET  /patchmango/api/get_file         - Download a file</p>
		<h2>Example Usage:</h2>
	  	<h3>Start server:</h3>
    	<p>./patchmango -port=8080 -verbose</p>
  		<h3>From the Remote Device...</h3>
  		<h3>Make a request to view all Versions:</h3>
    	<p>curl -H "Authorization: Bearer your-secret-token" \</p>
        <p>"http://localhost:8080/patchmango/api/versions"</p>
  		<h3>Make a request to view the Current Version:</h3>
    	<p>curl -H "Authorization: Bearer your-secret-token" \</p>
        <p>"http://localhost:8080/patchmango/api/current_version"</p>
  		<h3>Download and Execute a File:</h3>
    	<p>curl -H "Authorization: Bearer your-secret-token" \</p>
        <p>"http://localhost:8080/patchmango/api/get_file?file=filename" | bash</p>