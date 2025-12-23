package models

import "time"

// Request represents an HTTP or GraphQL request
type Request struct {
	ID        string                 `json:"id"`
	Name      string                 `json:"name"`
	Method    string                 `json:"method"`
	URL       string                 `json:"url"`
	Headers   map[string]string      `json:"headers"`
	Body      string                 `json:"body"`
	Type      string                 `json:"type"` // "http" or "graphql"
	QueryVars map[string]interface{} `json:"query_vars,omitempty"`
}

// Response represents an HTTP or GraphQL response
type Response struct {
	StatusCode int               `json:"status_code"`
	Status     string            `json:"status"`
	Headers    map[string]string `json:"headers"`
	Body       string            `json:"body"`
	Time       time.Duration     `json:"time"`
	Error      string            `json:"error,omitempty"`
}

// Collection represents a collection of requests
type Collection struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Requests []Request `json:"requests"`
}
