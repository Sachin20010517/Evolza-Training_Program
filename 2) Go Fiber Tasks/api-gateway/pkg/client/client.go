package client

import (
	"net/http"
	"time"
)

// HTTPClient is a reusable HTTP client with a timeout
var HTTPClient = &http.Client{Timeout: 10 * time.Second}
