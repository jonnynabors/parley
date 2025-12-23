package main

import (
	"context"

	"github.com/jonnynabors/parley/internal/http"
	"github.com/jonnynabors/parley/internal/models"
)

// App struct
type App struct {
	ctx        context.Context
	httpClient *http.Client
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		httpClient: http.NewClient(),
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// SendHTTPRequest sends an HTTP request
func (a *App) SendHTTPRequest(req models.Request) (*models.Response, error) {
	return a.httpClient.SendRequest(req)
}
