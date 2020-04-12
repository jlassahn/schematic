
package appui

import (
	"fmt"
	"time"

	"github.com/jlassahn/gogui"
	"github.com/jlassahn/schematic"
)


type schematicBox struct {

	schem *schematic.Schematic

	scrollbox gogui.ScrollBox
	zoom float64
	xPoint, yPoint int
	currentPage int
	snapToGrid bool
	showRulers bool

	mouseMode MouseMode
}

func CreateSchemBox(schem *schematic.Schematic) *schematicBox {
	ret := schematicBox{}
	ret.zoom = 1.0
	ret.schem = schem
	ret.currentPage = 1
	ret.snapToGrid = true
	ret.showRulers = true

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

func (box *schematicBox) SetMode(mode MouseMode) {
	box.mouseMode = mode
}

func (box *schematicBox) drawHandler(gfx gogui.Graphics) {

	t0 := time.Now()

	dc := DrawingContext{gfx, box.zoom, 20.0}
	width := box.schem.Settings.PageWidth
	height := box.schem.Settings.PageHeight

	dc.DrawOutline(width, height)
	schematic.DrawGrid(dc, width, height)
	if box.showRulers {
		// FIXME configuration for display colors!
		dc.Line(0, box.yPoint, width, box.yPoint, 0x1FF0, 3)
		dc.Line(box.xPoint, 0, box.xPoint, height, 0x1FF0, 3)
	}

	box.schem.DrawPage(dc, box.currentPage)

	t1 := time.Now()
	fmt.Printf("DRAW dt = %v\n", t1.Sub(t0))
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
	center_x = (center_x - 20) / old_zoom
	center_x = center_x*zm + 20
	center_y := float64(visible_top) + float64(visible_height)/2
	center_y = (center_y - 20) / old_zoom
	center_y = center_y*zm + 20
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

func (box *schematicBox) translateMouseCoords(x int, y int) (int, int) {

	tx := int((float64(x) - 20) / box.zoom)
	ty := int((float64(y) - 20) / box.zoom)
	if tx < 0 {
		tx = 0
	}
	if tx > box.schem.Settings.PageWidth {
		tx = box.schem.Settings.PageWidth
	}

	if ty < 0 {
		ty = 0
	}
	if ty > box.schem.Settings.PageHeight {
		ty = box.schem.Settings.PageHeight
	}

	if box.snapToGrid {
		tx = 12*((tx + 6)/12)
		ty = 12*((ty + 6)/12)
	}

	return tx, ty
}

func (box *schematicBox) mouseMoveHandler(x int, y int) {

	tx, ty := box.translateMouseCoords(x, y)
	box.mouseMode.MouseMove(tx, ty)
	if tx != box.xPoint || ty != box.yPoint {
		box.xPoint = tx
		box.yPoint = ty
		box.scrollbox.ForceRedraw()
	}
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

