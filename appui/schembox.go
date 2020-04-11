
package appui

import (
	"fmt"

	"github.com/jlassahn/gogui"
	"github.com/jlassahn/schematic"
)


type schematicBox struct {

	schem *schematic.Schematic

	scrollbox gogui.ScrollBox
	zoom float64
	currentPage int

	mouseMode MouseMode
}

func CreateSchemBox(schem *schematic.Schematic) *schematicBox {
	ret := schematicBox{}
	ret.zoom = 1.0
	ret.schem = schem
	ret.currentPage = 1

	ret.scrollbox = gogui.CreateScrollBox()
	dx := int(float64(schem.Settings.PageWidth)*ret.zoom + 40.0)
	dy := int(float64(schem.Settings.PageHeight)*ret.zoom + 40.0)
	ret.scrollbox.SetContentSize(dx, dy)
	ret.scrollbox.SetBackgroundColor(gogui.Color{128, 128, 128, 255})
	ret.scrollbox.HandleRedraw(ret.drawHandler)
	ret.scrollbox.HandleMouseMove(ret.mouseMoveHandler)
	ret.scrollbox.HandleMouseDown(ret.mouseDownHandler)
	ret.scrollbox.HandleMouseUp(ret.mouseUpHandler)
	ret.scrollbox.HandleMouseEnter(ret.mouseEnterHandler)
	ret.scrollbox.HandleMouseLeave(ret.mouseLeaveHandler)

	return &ret

}

func (box *schematicBox) Box() gogui.ScrollBox {
	return box.scrollbox
}

func (box *schematicBox) ZoomIn() {
	box.zoomTo(box.zoom*2)
}

func (box *schematicBox) ZoomOut() {
	box.zoomTo(box.zoom/2)
}

func (box *schematicBox) drawHandler(gfx gogui.Graphics) {

	fmt.Println("DRAW")

	dc := DrawingContext{gfx, box.zoom, 20.0}
	width := box.schem.Settings.PageWidth
	height := box.schem.Settings.PageHeight

	dc.DrawOutline(width, height)
	schematic.DrawGrid(dc, width, height)
	box.schem.DrawPage(dc, box.currentPage)
}

func (box *schematicBox) zoomTo(zm float64) {

	fmt.Println("Zoom")

	old_zoom := box.zoom
	old_dx := int(float64(box.schem.Settings.PageWidth)*box.zoom + 40.0)
	old_dy := int(float64(box.schem.Settings.PageHeight)*box.zoom + 40.0)
	box.zoom = zm
	dx := int(float64(box.schem.Settings.PageWidth)*box.zoom + 40.0)
	dy := int(float64(box.schem.Settings.PageHeight)*box.zoom + 40.0)

	visible_width := box.scrollbox.GetVisibleWidth()
	visible_height := box.scrollbox.GetVisibleHeight()
	visible_left := box.scrollbox.GetVisibleLeft()
	visible_top := box.scrollbox.GetVisibleTop()
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

	box.scrollbox.SetContentSize(dx, dy)
	box.scrollbox.SetVisibleLeftTop(left, top)

	box.scrollbox.ForceRedraw()
}

func (box *schematicBox) mouseMoveHandler(x int, y int) {
	fmt.Printf("MOUSE MOVE %v, %v\n", x, y)
}

func (box *schematicBox) mouseDownHandler(x int, y int, btn int) {
	fmt.Printf("MOUSE DOWN %v, %v, %v\n", x, y, btn)
}

func (box *schematicBox) mouseUpHandler(x int, y int, btn int) {
	fmt.Printf("MOUSE UP %v, %v, %v\n", x, y, btn)
}

func (box *schematicBox) mouseEnterHandler() {
	fmt.Printf("MOUSE ENTER\n")
}

func (box *schematicBox) mouseLeaveHandler() {
	fmt.Printf("MOUSE LEAVE\n")
}

