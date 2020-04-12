
package main

import (
	"github.com/jlassahn/gogui"
	"github.com/jlassahn/schematic/appui"
)

func (ui *ViewerUI) CreateSplashWindow() appui.Window {

	ret := SplashWindow{}

	ret.window = gogui.CreateWindow(gogui.WINDOW_TOPMOST | gogui.WINDOW_TITLED)
	ret.window.SetTitle("mSchematic Viewer")

	ret.window.SetPosition(
		gogui.Pos(50, -160),
		gogui.Pos(50, -80),
		gogui.Pos(50, 160),
		gogui.Pos(50, 80))
	ret.window.HandleClose(appui.Quit)

	button := gogui.CreateTextButton(XLT("Quit"))
	btnHeight := button.GetBestHeight()
	button.SetPosition(
		gogui.Pos(0, 220),
		gogui.Pos(100, -10-btnHeight),
		gogui.Pos(0, 310),
		gogui.Pos(100, -10))
	button.HandleClick(appui.Quit)
	ret.window.AddChild(button)

	button = gogui.CreateTextButton(XLT("Open..."))
	button.SetPosition(
		gogui.Pos(0, 220),
		gogui.Pos(100, -10-2*btnHeight),
		gogui.Pos(0, 310),
		gogui.Pos(100, -10-btnHeight))
	button.HandleClick(appui.RunOpenDialog)
	ret.window.AddChild(button)

	menu := gogui.CreateMenu()
	submenu := menu.GetApplicationMenu()
		item := gogui.CreateTextMenuItem(XLT("About"))
		item.HandleMenuSelect(ShowAboutWindow)
		submenu.AddMenuItem(item)
		submenu.AddSeparator()
		item = gogui.CreateTextMenuItem(XLT("Quit"))
		item.HandleMenuSelect(appui.Quit)
		item.SetShortcut("q")
		submenu.AddMenuItem(item)

	submenu = gogui.CreateTextMenuItem(XLT("File"))
		item = gogui.CreateTextMenuItem(XLT("Open..."))
		item.HandleMenuSelect(appui.RunOpenDialog)
		submenu.AddMenuItem(item)
		menu.AddMenuItem(submenu)
	ret.window.SetMenu(menu)

	ret.window.Show()

	return &ret
}


type SplashWindow struct {
	window gogui.Window
}

func (win *SplashWindow) Close() error {
	win.window.Destroy()
	return nil
}

//FIXME fake
func XLT(txt string) string {
	return txt
}

