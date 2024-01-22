package main

import (
	"context"
	"dzassistant/backend/assistant"
	"dzassistant/backend/config"
	"dzassistant/backend/mainApp"
	"dzassistant/backend/systemTp"
	"embed"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := mainApp.NewMainApp()
	systemTpApp := systemTp.NewSystemTp()
	assistantApp := assistant.NewAssistant()
	// Create application with options
	err := wails.Run(&options.App{
		Title:     config.AppName,
		Frameless: true,
		Width:     1024,
		Height:    768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Windows: &windows.Options{
			WebviewIsTransparent:                true,
			WindowIsTranslucent:                 true,
			DisableWindowIcon:                   false,
			IsZoomControlEnabled:                false,
			ZoomFactor:                          0,
			DisableFramelessWindowDecorations:   false,
			WebviewUserDataPath:                 "",
			WebviewBrowserPath:                  "",
			Theme:                               0,
			CustomTheme:                         nil,
			BackdropType:                        windows.Auto,
			Messages:                            nil,
			ResizeDebounceMS:                    0,
			OnSuspend:                           nil,
			OnResume:                            nil,
			WebviewGpuIsDisabled:                false,
			WebviewDisableRendererCodeIntegrity: false,
			EnableSwipeGestures:                 false,
		},
		Mac: &mac.Options{
			TitleBar:             nil,
			Appearance:           "",
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			Preferences:          nil,
			About:                nil,
			OnFileOpen:           nil,
			OnUrlOpen:            nil,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			app.Startup(ctx)
			systemTpApp.Startup(ctx)
			assistantApp.Startup(ctx)
		},
		Bind: []interface{}{
			app,
			systemTpApp,
			assistantApp,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
