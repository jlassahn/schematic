
package main

import (
	"github.com/jlassahn/gogui"
	"github.com/jlassahn/schematic"
	"github.com/jlassahn/schematic/appui"
)

type SchematicWindow struct {
	window gogui.Window
	schem *schematic.Schematic
	schemview appui.SchemView
	schembox appui.SchemBox
}

func (ui *ViewerUI) CreateSchematicWindow(schem *schematic.Schematic) appui.Window {

	ret := SchematicWindow{}
	ret.schem = schem
	ret.window = gogui.CreateWindow(gogui.WINDOW_NORMAL)
	ret.schembox = appui.CreateSchemBox(ret.schem)
	ret.schemview = appui.CreateSchemView(&ret, ret.schembox, ret.schem)

	el := ret.schembox.Box()
	el.SetPosition(
		gogui.Pos(0,100),
		gogui.Pos(0,50),
		gogui.Pos(100,0),
		gogui.Pos(100,0))
	ret.window.AddChild(el)

	ret.window.SetPosition(
		gogui.Pos(50, 0),
		gogui.Pos(10, 0),
		gogui.Pos(75, 0),
		gogui.Pos(50, 0))
	ret.window.HandleClose(ret.schemview.Close)

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

	submenu = gogui.CreateTextMenuItem(XLT("View"))
		item = gogui.CreateTextMenuItem(XLT("Zoom In"))
		item.SetShortcut("+")
		item.HandleMenuSelect(ret.schemview.ZoomIn)
		submenu.AddMenuItem(item)

		item = gogui.CreateTextMenuItem(XLT("Zoom Out"))
		item.HandleMenuSelect(ret.schemview.ZoomOut)
		item.SetShortcut("-")
		submenu.AddMenuItem(item)
		menu.AddMenuItem(submenu)

	ret.window.SetMenu(menu)

	ret.window.Show()

	return &ret
}

func (win *SchematicWindow) Close() error {
	//FIXME destroy contents
	win.window.Destroy()
	return nil
}

