
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

	submenu = gogui.CreateTextMenuItem("View")
		item = gogui.CreateTextMenuItem("Zoom In")
		item.HandleMenuSelect(ret.zoomIn)
		submenu.AddMenuItem(item)

		item = gogui.CreateTextMenuItem("Zoom Out")
		item.HandleMenuSelect(ret.zoomOut)
		submenu.AddMenuItem(item)
		menu.AddMenuItem(submenu)

	window := gogui.CreateWindow(gogui.WINDOW_NORMAL)
	window.SetMenu(menu)

	ret.window = window
	window.SetPosition(
		gogui.Pos(50, 0),
		gogui.Pos(10, 0),
		gogui.Pos(75, 0),
		gogui.Pos(50, 0))
	window.HandleClose(ret.closeHandler)

	scroll := gogui.CreateScrollBox()
	ret.scroll = scroll
	scroll.SetPosition(
		gogui.Pos(0,100),
		gogui.Pos(0,50),
		gogui.Pos(100,0),
		gogui.Pos(100,0))

	dx := int(float64(schem.Settings.PageWidth)*ret.zoom + 40.0)
	dy := int(float64(schem.Settings.PageHeight)*ret.zoom + 40.0)
	scroll.SetContentSize(dx, dy)
	scroll.SetBackgroundColor(gogui.Color{128, 128, 128, 255})
	scroll.HandleRedraw(ret.drawHandler)
	window.AddChild(scroll)

	btnPos := 50
	button := gogui.CreateTextButton("Button 1")
	btnHeight := button.GetBestHeight()
	button.SetPosition(
		gogui.Pos(0, 0),
		gogui.Pos(0, btnPos),
		gogui.Pos(0, 100),
		gogui.Pos(0, btnPos+btnHeight))
	button.HandleClick(ret.clickHandler)
	window.AddChild(button)

	btnPos += btnHeight
	button = gogui.CreateTextButton("Button 2")
	button.SetPosition(
		gogui.Pos(0, 0),
		gogui.Pos(0, btnPos),
		gogui.Pos(0, 100),
		gogui.Pos(0, btnPos+btnHeight))
	button.HandleClick(ret.clickHandler)
	window.AddChild(button)

	btnPos += btnHeight
	button = gogui.CreateTextButton("Button 3")
	button.SetPosition(
		gogui.Pos(0, 0),
		gogui.Pos(0, btnPos),
		gogui.Pos(0, 100),
		gogui.Pos(0, btnPos+btnHeight))
	button.HandleClick(ret.clickHandler)
	window.AddChild(button)

	window.Show()

	ViewListAdd(&ret)
	return &ret, nil
}

type SchematicView struct {
	zoom float64
	schem *schematic.Schematic
	window gogui.Window
	scroll gogui.ScrollBox
	currentPage int
}

func (view *SchematicView) Close() error {
	view.window.Destroy()
	ViewListRemove(view)
	return nil
}

func (view *SchematicView) clickHandler() {
	fmt.Println("CLICK")
}

func (view *SchematicView) closeHandler() {
	view.window.Destroy()
	ViewListRemove(view)
}

func (view *SchematicView) drawHandler(gfx gogui.Graphics) {

	fmt.Println("DRAW")

	dc := DrawingContext{gfx, view.zoom, 20.0}
	width := view.schem.Settings.PageWidth
	height := view.schem.Settings.PageHeight

	dc.DrawOutline(width, height)
	schematic.DrawGrid(dc, width, height)
	view.schem.DrawPage(dc, view.currentPage)
}

func (view *SchematicView) zoomIn() {
	fmt.Println("Zoom IN")
	//FIXME reorganize
	view.zoom = view.zoom*2
	dx := int(float64(view.schem.Settings.PageWidth)*view.zoom + 40.0)
	dy := int(float64(view.schem.Settings.PageHeight)*view.zoom + 40.0)
	view.scroll.SetContentSize(dx, dy)
	view.scroll.ForceRedraw()
}

func (view *SchematicView) zoomOut() {
	fmt.Println("Zoom OUT")
	view.zoom = view.zoom/2
	dx := int(float64(view.schem.Settings.PageWidth)*view.zoom + 40.0)
	dy := int(float64(view.schem.Settings.PageHeight)*view.zoom + 40.0)

	fmt.Printf("scrollpos %d, %d  %d, %d\n",
		view.scroll.GetVisibleWidth(),
		view.scroll.GetVisibleHeight(),
		view.scroll.GetVisibleLeft(),
		view.scroll.GetVisibleTop());

	view.scroll.SetContentSize(dx, dy)

	fmt.Printf("scrollpos %d, %d  %d, %d\n",
		view.scroll.GetVisibleWidth(),
		view.scroll.GetVisibleHeight(),
		view.scroll.GetVisibleLeft(),
		view.scroll.GetVisibleTop());

	view.scroll.ForceRedraw()
}

