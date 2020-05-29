
package main

import (
	"github.com/jlassahn/gogui"
	"github.com/jlassahn/schematic"
	"github.com/jlassahn/schematic/appui"
)

type SchematicWindow struct {
	window gogui.Window
	schemedit appui.SchemEdit

	editState appui.SchemEditState
}

func (ui *MSchemUI) CreateSchematicWindow(schem *schematic.Schematic) appui.Window {

	ret := SchematicWindow{}
	ret.window = gogui.CreateWindow(gogui.WINDOW_NORMAL)

	ret.editState.MainWindow = &ret
	ret.editState.Schem = schem
	ret.editState.SchemBox = appui.CreateSchemBox(schem)

	ret.schemedit = appui.CreateSchemEdit(&ret.editState)

	el := ret.editState.SchemBox.Box()
	el.SetPosition(
		gogui.Pos(0,100),
		gogui.Pos(0,50),
		gogui.Pos(100,0),
		gogui.Pos(100,0))
	ret.window.AddChild(el)

	//FIXME fake
	prop := CreatePropertyBox(&ret.editState)
	prop.FrameBox.SetPosition(
		gogui.Pos(0,0),
		gogui.Pos(0,0),
		gogui.Pos(0,100),
		gogui.Pos(100,0))
	ret.window.AddChild(prop.FrameBox)

	/*
	txt := gogui.CreateTextLineInput()
	txt.SetText("fake value")
	txt.HandleChange(FakeTextHandler)
	txt.SetPosition(
		gogui.Pos(0,0),
		gogui.Pos(0,50),
		gogui.Pos(0,100),
		gogui.Pos(0,50 + txt.GetBestHeight()))
	ret.window.AddChild(txt)
	ret.editState.NewLineWidth = 5
	ret.editState.NewLineColor = 0x1F0F
	*/

	ret.window.SetPosition(
		gogui.Pos(50, 0),
		gogui.Pos(10, 0),
		gogui.Pos(75, 0),
		gogui.Pos(50, 0))
	ret.window.HandleClose(ret.schemedit.Close)

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

	submenu = gogui.CreateTextMenuItem(XLT("Edit"))
		// FIXME add selection checkmark to these menu items
		// FIXME add keyboard shortcuts
		item = gogui.CreateTextMenuItem(XLT("Select Elements"))
		item.HandleMenuSelect(ret.schemedit.ModeSelect)
		submenu.AddMenuItem(item)
		item = gogui.CreateTextMenuItem(XLT("Add Lines"))
		item.HandleMenuSelect(ret.schemedit.ModeMakeLines)
		submenu.AddMenuItem(item)
		submenu.AddSeparator()
		item = gogui.CreateTextMenuItem(XLT("Undo"))
		item.SetShortcut("z")
		item.HandleMenuSelect(ret.schemedit.Undo)
		submenu.AddMenuItem(item)
		item = gogui.CreateTextMenuItem(XLT("Redo"))
		item.SetShortcut("Z")
		item.HandleMenuSelect(ret.schemedit.Redo)
		submenu.AddMenuItem(item)
		menu.AddMenuItem(submenu)

	submenu = gogui.CreateTextMenuItem(XLT("View"))
		item = gogui.CreateTextMenuItem(XLT("Zoom In"))
		item.SetShortcut("+")
		item.HandleMenuSelect(ret.schemedit.ZoomIn)
		submenu.AddMenuItem(item)

		item = gogui.CreateTextMenuItem(XLT("Zoom Out"))
		item.HandleMenuSelect(ret.schemedit.ZoomOut)
		item.SetShortcut("-")
		submenu.AddMenuItem(item)
		menu.AddMenuItem(submenu)

	ret.window.SetMenu(menu)

	ret.window.Show()

	return &ret
}

func (win *SchematicWindow) Close() error {
	//FIXME check for unsaved stuff
	//FIXME destroy contents
	win.window.Destroy()
	return nil
}

