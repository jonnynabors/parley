package http

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/jonnynabors/parley/internal/models"
)

// Client handles HTTP requests
type Client struct {
	httpClient *http.Client
}

// NewClient creates a new HTTP client
func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// SendRequest sends an HTTP request and returns the response
func (c *Client) SendRequest(req models.Request) (*models.Response, error) {
	start := time.Now()

	// Create HTTP request
	httpReq, err := http.NewRequest(req.Method, req.URL, bytes.NewBufferString(req.Body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Add headers
	for key, value := range req.Headers {
		httpReq.Header.Set(key, value)
	}

	// Send request
	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return &models.Response{
			Error: err.Error(),
			Time:  time.Since(start),
		}, err
	}
	defer httpResp.Body.Close()

	// Read response body
	bodyBytes, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Build response headers map
	headers := make(map[string]string)
	for key, values := range httpResp.Header {
		if len(values) > 0 {
			headers[key] = values[0]
		}
	}

	return &models.Response{
		StatusCode: httpResp.StatusCode,
		Status:     httpResp.Status,
		Headers:    headers,
		Body:       string(bodyBytes),
		Time:       time.Since(start),
	}, nil
}
