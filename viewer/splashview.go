
package main

import (
	"fmt"

	"github.com/jlassahn/gogui"
)

func CreateSplashWindow() (*SplashView, error) {

	ret := SplashView{}

	menu := gogui.CreateMenu()
	submenu := gogui.CreateTextMenuItem("Application")
		item := gogui.CreateTextMenuItem("About")
		submenu.AddMenuItem(item)
		// FIXME add separator
		item = gogui.CreateTextMenuItem("Quit")
		item.HandleMenuSelect(ret.quitHandler)
		submenu.AddMenuItem(item)
		menu.AddMenuItem(submenu)

	submenu = gogui.CreateTextMenuItem("File")
		item = gogui.CreateTextMenuItem("Open...")
		item.HandleMenuSelect(ret.openHandler)
		submenu.AddMenuItem(item)
		menu.AddMenuItem(submenu)

	window := gogui.CreateWindow(gogui.WINDOW_TOPMOST | gogui.WINDOW_TITLED)
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
	button.HandleClick(ret.quitHandler)
	window.AddChild(button)

	button = gogui.CreateTextButton("Open...")
	button.SetPosition(
		gogui.Pos(0, 220),
		gogui.Pos(100, -10-2*btnHeight),
		gogui.Pos(0, 310),
		gogui.Pos(100, -10-btnHeight))
	button.HandleClick(ret.openHandler)
	window.AddChild(button)

	window.Show()

	return &ret, nil
}

type SplashView struct {
	window gogui.Window
}

func (view *SplashView) Close() error {
	view.window.Destroy();
	return nil
}

func (view *SplashView) closeHandler() {
	QuitApp()
}

func (view *SplashView) quitHandler() {
	QuitApp()
}

func (view *SplashView) openHandler() {
	fmt.Println("FIXME got open button")
	name := RunOpenDialog()
	if name == "" {
		return
	}

	_,err := CreateSchematicViewFromFile(name)
	if err != nil {
		//FIXME display error
		return
	}

	view.window.Destroy();
}


