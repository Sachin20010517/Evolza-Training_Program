package services

import (
	"api-gateway/pkg/client"
	"bytes"
	"io/ioutil"
	"net/http"
)

// BackendServiceURL defines the base URL for the user backend service
const BackendServiceURL = "http://user-backend-service"

// ForwardUserRequest forwards the API request to the user backend service
func ForwardUserRequest(method, path string, body []byte) (*http.Response, error) {
	// Create a new HTTP request
	req, err := http.NewRequest(method, BackendServiceURL+path, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	// Set appropriate headers
	req.Header.Set("Content-Type", "application/json")

	// Use the custom HTTP client to perform the request
	resp, err := client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Read and return the response
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Wrap the response
	response := &http.Response{
		StatusCode: resp.StatusCode,
		Body:       ioutil.NopCloser(bytes.NewBuffer(body)),
		Header:     resp.Header,
	}
	return response, nil
}
