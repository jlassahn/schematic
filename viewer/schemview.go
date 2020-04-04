
package main

import (
	"fmt"

	"github.com/jlassahn/gogui"
	"github.com/jlassahn/schematic"
)

// FIXME destroy

func CreateSchematicViewFromFile(filename string) (*SchematicView, error) {

	schem, err := schematic.LoadSchematic(filename)
	if err != nil {
		return nil, err
	}

	ret := SchematicView{}

	window := gogui.CreateWindow(gogui.WINDOW_NORMAL)
	ret.window = window

	window.SetPosition(
		gogui.Pos(50, 0),
		gogui.Pos(10, 0),
		gogui.Pos(75, 0),
		gogui.Pos(50, 0))
	window.HandleClose(ret.closeHandler)

	ret.schemBox = CreateSchematicBox(schem)
	ret.schemBox.Scroll.SetPosition( // FIXME encapsulate
		gogui.Pos(0,100),
		gogui.Pos(0,50),
		gogui.Pos(100,0),
		gogui.Pos(100,0))

	window.AddChild(ret.schemBox.Scroll)


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
		item.SetShortcut("+")
		item.HandleMenuSelect(ret.schemBox.ZoomIn)
		submenu.AddMenuItem(item)

		item = gogui.CreateTextMenuItem("Zoom Out")
		item.HandleMenuSelect(ret.schemBox.ZoomOut)
		item.SetShortcut("-")
		submenu.AddMenuItem(item)
		menu.AddMenuItem(submenu)

	window.SetMenu(menu)


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
	window gogui.Window

	schemBox *SchematicBox
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
	view.Close()
}

type MouseMode interface {
	MouseMove(x int, y int)
	MouseDown(x int, y int, btn int)
	MouseUp(x int, y int, btn int)
	Cancel()
}

type SchematicBox struct {
	Scroll gogui.ScrollBox
	zoom float64
	schem *schematic.Schematic
	currentPage int

	mouseMode MouseMode
}

func CreateSchematicBox(schem *schematic.Schematic) *SchematicBox {
	ret := SchematicBox{}
	ret.zoom = 1.0
	ret.schem = schem
	ret.currentPage = 1

	ret.Scroll = gogui.CreateScrollBox()
	dx := int(float64(schem.Settings.PageWidth)*ret.zoom + 40.0)
	dy := int(float64(schem.Settings.PageHeight)*ret.zoom + 40.0)
	ret.Scroll.SetContentSize(dx, dy)
	ret.Scroll.SetBackgroundColor(gogui.Color{128, 128, 128, 255})
	ret.Scroll.HandleRedraw(ret.drawHandler)
	ret.Scroll.HandleMouseMove(ret.mouseMoveHandler)
	ret.Scroll.HandleMouseDown(ret.mouseDownHandler)
	ret.Scroll.HandleMouseUp(ret.mouseUpHandler)
	ret.Scroll.HandleMouseEnter(ret.mouseEnterHandler)
	ret.Scroll.HandleMouseLeave(ret.mouseLeaveHandler)

	return &ret

}

func (box *SchematicBox) ZoomIn() {
	box.zoomTo(box.zoom*2)
}

func (box *SchematicBox) ZoomOut() {
	box.zoomTo(box.zoom/2)
}

func (box *SchematicBox) drawHandler(gfx gogui.Graphics) {

	fmt.Println("DRAW")

	dc := DrawingContext{gfx, box.zoom, 20.0}
	width := box.schem.Settings.PageWidth
	height := box.schem.Settings.PageHeight

	dc.DrawOutline(width, height)
	schematic.DrawGrid(dc, width, height)
	box.schem.DrawPage(dc, box.currentPage)
}

func (box *SchematicBox) zoomTo(zm float64) {

	fmt.Println("Zoom")

	old_zoom := box.zoom
	old_dx := int(float64(box.schem.Settings.PageWidth)*box.zoom + 40.0)
	old_dy := int(float64(box.schem.Settings.PageHeight)*box.zoom + 40.0)
	box.zoom = zm
	dx := int(float64(box.schem.Settings.PageWidth)*box.zoom + 40.0)
	dy := int(float64(box.schem.Settings.PageHeight)*box.zoom + 40.0)

	visible_width := box.Scroll.GetVisibleWidth()
	visible_height := box.Scroll.GetVisibleHeight()
	visible_left := box.Scroll.GetVisibleLeft()
	visible_top := box.Scroll.GetVisibleTop()
	if visible_width > old_dx {
		visible_width = old_dx
	}
	if visible_height > old_dy {
		visible_height = old_dy
	}

	center_x := float64(visible_left) + float64(visible_width)/2
	center_x = (center_x - 40) / old_zoom
	center_x = center_x*zm + 40
	center_y := float64(visible_top) + float64(visible_height)/2
	center_y = (center_y - 40) / old_zoom
	center_y = center_y*zm + 40
	left := int(center_x - float64(visible_width)/2)
	top := int(center_y - float64(visible_height)/2)
	if left < 0 {
		left = 0
	}
	if top < 0 {
		top = 0
	}

	box.Scroll.SetContentSize(dx, dy)
	box.Scroll.SetVisibleLeftTop(left, top)

	box.Scroll.ForceRedraw()
}

func (box *SchematicBox) mouseMoveHandler(x int, y int) {
	fmt.Printf("MOUSE MOVE %v, %v\n", x, y)
}

func (box *SchematicBox) mouseDownHandler(x int, y int, btn int) {
	fmt.Printf("MOUSE DOWN %v, %v, %v\n", x, y, btn)
}

func (box *SchematicBox) mouseUpHandler(x int, y int, btn int) {
	fmt.Printf("MOUSE UP %v, %v, %v\n", x, y, btn)
}

func (box *SchematicBox) mouseEnterHandler() {
	fmt.Printf("MOUSE ENTER\n")
}

func (box *SchematicBox) mouseLeaveHandler() {
	fmt.Printf("MOUSE LEAVE\n")
}

