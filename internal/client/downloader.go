package client

import (
	"io"
	"net/http"
)

func (h *HTTPClient) downloadFile(url string) (io.ReadCloser, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errorConnection
	}

	// We don't close resp.Body! Don't forget to close it after use.
	return resp.Body, nil
}
