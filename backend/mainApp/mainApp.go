package mainApp

import (
	"context"
	"dzassistant/backend/config"
)

// MainApp struct
type MainApp struct {
	ctx context.Context
}

// NewMainApp creates a new App application struct
func NewMainApp() *MainApp {
	return &MainApp{}
}

// Startup is called when the app starts. The context is saved
func (a *MainApp) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *MainApp) GetAppName() string {
	return config.AppName
}
