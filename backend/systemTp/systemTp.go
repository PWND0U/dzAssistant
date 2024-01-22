package systemTp

import (
	"context"
	_ "embed"
	"github.com/energye/systray"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed icon.ico
var icon []byte

type SystemTp struct {
	ctx context.Context
}

func NewSystemTp() *SystemTp {
	return &SystemTp{}
}
func (a *SystemTp) systemTray() {
	systray.SetIcon(icon)

	show := systray.AddMenuItem("主窗口", "显示主窗口")
	show.Click(func() { runtime.WindowShow(a.ctx) })
	show.SetIcon(icon)
	systray.AddSeparator()
	exit := systray.AddMenuItem("退出", "退出程序")
	exit.Click(func() { runtime.Quit(a.ctx) })

	systray.SetOnClick(func(menu systray.IMenu) { runtime.WindowShow(a.ctx) })
	systray.SetOnRClick(func(menu systray.IMenu) { menu.ShowMenu() })
	systray.SetTooltip("这是一个很好用的工具")
}
func (a *SystemTp) Startup(ctx context.Context) {
	a.ctx = ctx
	systray.Run(a.systemTray, func() {})
}
