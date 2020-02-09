
package main

import (
	"fmt"

	"github.com/jlassahn/gogui"
	"github.com/jlassahn/schematic"
)

func CreateSchematicViewFromFile(filename string) (*SchematicView, error) {

	schem, err := schematic.LoadSchematic(filename)
	if err != nil {
		return nil, err
	}

	ret := SchematicView{}
	ret.schem = schem
	ret.zoom = 1.0
	ret.currentPage = 1

	menu := gogui.CreateMenu()
	submenu := gogui.CreateTextMenuItem("Application")
		item := gogui.CreateTextMenuItem("About")
		submenu.AddMenuItem(item)
		// FIXME add separator
		item = gogui.CreateTextMenuItem("Quit")
		item.HandleMenuSelect(ret.fakeMenuHandler)
		submenu.AddMenuItem(item)
		menu.AddMenuItem(submenu)

	submenu = gogui.CreateTextMenuItem("File")
		item = gogui.CreateTextMenuItem("Open...")
		submenu.AddMenuItem(item)
		menu.AddMenuItem(submenu)

	//gogui.SetMainMenu(menu) //FIXME remove from GUI API
	window := gogui.CreateWindow()
	window.SetMenu(menu)

	ret.window = window
	window.SetPosition(
		gogui.Pos(50, 0),
		gogui.Pos(10, 0),
		gogui.Pos(75, 0),
		gogui.Pos(50, 0))
	window.HandleClose(ret.closeHandler)

	scroll := gogui.CreateScrollBox()
	scroll.SetPosition(
		gogui.Pos(0,100),
		gogui.Pos(0,50),
		gogui.Pos(100,0),
		gogui.Pos(100,0))

	dx := int(float64(schem.Settings.PageWidth)*ret.zoom + 40.0)
	dy := int(float64(schem.Settings.PageHeight)*ret.zoom + 40.0)
	scroll.SetContentSize(dx, dy)
	scroll.HandleRedraw(ret.drawHandler)
	window.AddChild(scroll)

	btnPos := 50
	button := gogui.CreateTextButton("Button 1")
	btnHeight := button.GetBestHeight()
	fmt.Printf("button height = %d\n", btnHeight)
	button.SetPosition(
		gogui.Pos(0, 0),
		gogui.Pos(0, btnPos),
		gogui.Pos(0, 100),
		gogui.Pos(0, btnPos+btnHeight))
	button.HandleClick(ret.clickHandler)
	window.AddChild(button)

	btnPos += btnHeight
	button = gogui.CreateTextButton("Button 2")
	fmt.Printf("button height = %d\n", btnHeight)
	button.SetPosition(
		gogui.Pos(0, 0),
		gogui.Pos(0, btnPos),
		gogui.Pos(0, 100),
		gogui.Pos(0, btnPos+btnHeight))
	button.HandleClick(ret.clickHandler)
	window.AddChild(button)

	btnPos += btnHeight
	button = gogui.CreateTextButton("Button 3")
	fmt.Printf("button height = %d\n", btnHeight)
	button.SetPosition(
		gogui.Pos(0, 0),
		gogui.Pos(0, btnPos),
		gogui.Pos(0, 100),
		gogui.Pos(0, btnPos+btnHeight))
	button.HandleClick(ret.clickHandler)
	window.AddChild(button)

	window.Show()

	return &ret, nil
}

type SchematicView struct {
	zoom float64
	schem *schematic.Schematic
	window gogui.Window
	currentPage int
}

func (view *SchematicView) clickHandler() {
	fmt.Println("CLICK")
}

func (view *SchematicView) closeHandler() {
	QuitApp()
}

func (view *SchematicView) fakeMenuHandler() {
	fmt.Println("MENU")
}

func (view *SchematicView) drawHandler(gfx gogui.Graphics) {

	dc := DrawingContext{gfx, view.zoom, 20.0}
	width := view.schem.Settings.PageWidth
	height := view.schem.Settings.PageHeight

	dc.DrawOutline(width, height)
	schematic.DrawGrid(dc, width, height)
	view.schem.DrawPage(dc, view.currentPage)
}

