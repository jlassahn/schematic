
package main

import (
	"github.com/jlassahn/gogui"
)

var splashView *SplashView = nil

func CreateSplashWindow() (*SplashView, error) {

	if splashView != nil {
		return splashView, nil
	}

	ret := SplashView{}

	menu := gogui.CreateMenu()
	submenu := menu.GetApplicationMenu()
		item := gogui.CreateTextMenuItem("About")
		submenu.AddMenuItem(item)
		submenu.AddSeparator()
		item = gogui.CreateTextMenuItem("Quit")
		item.HandleMenuSelect(QuitApp)
		item.SetShortcut("q")
		submenu.AddMenuItem(item)

	submenu = gogui.CreateTextMenuItem("File")
		item = gogui.CreateTextMenuItem("Open...")
		item.HandleMenuSelect(RunOpenDialog)
		submenu.AddMenuItem(item)
		menu.AddMenuItem(submenu)

	window := gogui.CreateWindow(gogui.WINDOW_TOPMOST | gogui.WINDOW_TITLED)
	window.SetTitle("mschematic Viewer")
	window.SetMenu(menu)

	ret.window = window
	window.SetPosition(
		gogui.Pos(50, -160),
		gogui.Pos(50, -80),
		gogui.Pos(50, 160),
		gogui.Pos(50, 80))
	window.HandleClose(ret.closeHandler)

	button := gogui.CreateTextButton("Quit")
	btnHeight := button.GetBestHeight()
	button.SetPosition(
		gogui.Pos(0, 220),
		gogui.Pos(100, -10-btnHeight),
		gogui.Pos(0, 310),
		gogui.Pos(100, -10))
	button.HandleClick(QuitApp)
	window.AddChild(button)

	button = gogui.CreateTextButton("Open...")
	button.SetPosition(
		gogui.Pos(0, 220),
		gogui.Pos(100, -10-2*btnHeight),
		gogui.Pos(0, 310),
		gogui.Pos(100, -10-btnHeight))
	button.HandleClick(RunOpenDialog)
	window.AddChild(button)

	window.Show()

	splashView = &ret
	ViewListAdd(&ret)
	return &ret, nil
}

func CloseSplashView() {
	if splashView != nil {
		splashView.Close()
		splashView = nil
	}
}

type SplashView struct {
	window gogui.Window
}

func (view *SplashView) Close() error {
	view.window.Destroy()
	ViewListRemove(view)
	return nil
}

func (view *SplashView) closeHandler() {
	view.Close()
}

