
package appui

import (
	"math"

	"github.com/jlassahn/gogui"
	"github.com/jlassahn/schematic"
)

func GetDrawingContext(gfx gogui.Graphics, settings *schematic.DrawingSettings, width int, height int, zoom float64, padding float64) schematic.DrawingContext {

	drv := drawDriver {
		gfx,
		width,
		height,
		zoom,
		padding,
	}

	drv.DrawOutline(width, height)
	return schematic.WrapDrawingDriver(&drv, settings)
}

	//dc.DrawOutline(width, height)

type drawDriver struct {
	gogui.Graphics
	width int
	height int
	zoom float64
	pad float64
}

func (dc *drawDriver) setColor(col int) {
	r := (col >> 8) & 0xF
	g := (col >> 4) & 0xF
	b := col& 0xF

	r = r | (r<<4)
	g = g | (g<<4)
	b = b | (b<<4)
	dc.SetStrokeColor(gogui.Color{uint8(r), uint8(g), uint8(b), 255})
}

func (dc *drawDriver) coord(x int) float64 {
	return float64(x)*dc.zoom + dc.pad
}

func (dc *drawDriver) Line(x0 int, y0 int, x1 int, y1 int, color int, width int) {

	dc.setColor(color)
	dc.SetLineWidth(float64(width)*dc.zoom)

	dc.StartPath(dc.coord(x0), dc.coord(y0))
	dc.LineTo(dc.coord(x1), dc.coord(y1))
	dc.StrokePath()

}

func (dc *drawDriver) Text(txt string, x0 int, y0 int, rotate int, anchor int, color int, size int, font int) {

	if font >= 16 || font < 0 {
		font = 0
	}

	dc.setColor(color)
	dc.SetFont(fonts[font], float64(size)*dc.zoom)
	width := dc.MeasureText(txt)

	if anchor == schematic.ANCHOR_START {
		width = 0
	} else if anchor != schematic.ANCHOR_END {
		width = width*0.5
	}

	rad := float64(rotate)*2*3.1415926/360;

	x := dc.coord(x0) - width*math.Cos(rad)
	y := dc.coord(y0) - width*math.Sin(rad)

	dc.DrawText(x, y, float64(rotate), txt)
}

func (dc *drawDriver) Curve(x0 int, y0 int, cx0 int, cy0 int, cx1 int, cy1 int, x1 int, y1 int, color int, width int) {

	dc.setColor(color)
	dc.SetLineWidth(float64(width)*dc.zoom)
	dc.StartPath(dc.coord(x0), dc.coord(y0))
	dc.CurveTo(dc.coord(cx0), dc.coord(cy0), dc.coord(cx1), dc.coord(cy1), dc.coord(x1), dc.coord(y1))
	dc.StrokePath()
}

func (dc *drawDriver) DrawOutline(width int, height int) {
	//dc.SetFillColor(gogui.Color{128, 128, 128, 255})
	//dc.FillCanvas()
	dc.SetFillColor(gogui.Color{255, 255, 255, 255})
	dc.StartPath(dc.coord(0), dc.coord(0))
	dc.LineTo(dc.coord(width), dc.coord(0))
	dc.LineTo(dc.coord(width), dc.coord(height))
	dc.LineTo(dc.coord(0), dc.coord(height))
	dc.ClosePath()
	dc.FillPath()
}

