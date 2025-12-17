# patchmango
Embedded Device Updater using Go as a server-side API

### Testing
<p>-> go run main.go</p>
<p>curl -H "Authorization: Bearer your-secret-token-here-make-it-long-and-random" "http://localhost:8080/patchmango/api/versions"</p>
<p>TODO - Need Test Suite</p>
Test without auth header, test with incorrect auth header, test with bad token, test with good token