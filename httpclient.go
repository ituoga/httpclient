package httpclient

import (
	"encoding/json"
	"io"
	"net/http"
)

// Http client
// default http client is used
var Do = http.DefaultClient.Do

// Get requests a URL and decodes the response body to struct.
func Get[R any](url string) (*R, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	var r *R
	err = json.NewDecoder(response.Body).Decode(&r)
	return r, err
}

// Post requests a URL with a body and decodes the response body to struct.
func Post[R any](url string, contentType string, body io.Reader) (*R, error) {
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", contentType)
	response, err := Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	var r *R
	err = json.NewDecoder(response.Body).Decode(&r)
	return r, err
}
