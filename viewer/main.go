
package main

import (
	"fmt"
	"math"

	"github.com/jlassahn/gogui"
	"github.com/jlassahn/schematic"
)

var fonts []gogui.Font

var zoom = 0.5

var mySchem *schematic.Schematic

func clickHandler() {
	fmt.Println("CLICK")
}

func closeHandler() {
	gogui.StopEventLoop(0)
}

func fakeMenuHandler() {
	fmt.Println("MENU")
}

func drawHandler(gfx gogui.Graphics) {

	//default background is white
	gfx.FillCanvas()

	dc := DrawingContext{gfx, zoom}
	schematic.DrawGrid(dc, mySchem.Settings.PageWidth, mySchem.Settings.PageHeight)
	mySchem.DrawPage(dc, 1)
}

func main() {

	fonts = []gogui.Font{
		gogui.CreateFont("DejaVuSansMono", gogui.FONT_NORMAL),
		gogui.CreateFont("DejaVuSansMono", gogui.FONT_NORMAL),
		gogui.CreateFont("DejaVuSansMono", gogui.FONT_NORMAL),
		gogui.CreateFont("DejaVuSansMono", gogui.FONT_NORMAL),

		gogui.CreateFont("DejaVuSansMono", gogui.FONT_NORMAL),
		gogui.CreateFont("DejaVuSansMono-Oblique", gogui.FONT_NORMAL),
		gogui.CreateFont("DejaVuSansMono-Bold", gogui.FONT_NORMAL),
		gogui.CreateFont("DejaVuSansMono-BoldOblique", gogui.FONT_NORMAL),

		gogui.CreateFont("DejaVuSans", gogui.FONT_NORMAL),
		gogui.CreateFont("DejaVuSans-Oblique", gogui.FONT_NORMAL),
		gogui.CreateFont("DejaVuSans-Bold", gogui.FONT_NORMAL),
		gogui.CreateFont("DejaVuSans-BoldOblique", gogui.FONT_NORMAL),

		gogui.CreateFont("DejaVuSerif", gogui.FONT_NORMAL),
		gogui.CreateFont("DejaVuSerif-Italic", gogui.FONT_NORMAL),
		gogui.CreateFont("DejaVuSerif-Bold", gogui.FONT_NORMAL),
		gogui.CreateFont("DejaVuSerif-BoldItalic", gogui.FONT_NORMAL),

		/*
	"mono_n",
	"mono_n",
	"mono_n",
	"mono_n",
	"mono_n",
	"mono_i",
	"mono_b",
	"mono_bi",
	"sans_n",
	"sans_i",
	"sans_b",
	"sans_bi",
	"serif_n",
	"serif_i",
	"serif_b",
	"serif_bi",
	*/
	}

	mySchem,_ = schematic.LoadSchematic("tschem.json")

	gogui.Init()
	defer gogui.Exit()

	menu := gogui.CreateMenu()
	submenu := gogui.CreateTextMenuItem("Application")
		item := gogui.CreateTextMenuItem("About")
		submenu.AddMenuItem(item)
		// FIXME add separator
		item = gogui.CreateTextMenuItem("Quit")
		item.HandleMenuSelect(fakeMenuHandler)
		submenu.AddMenuItem(item)
		menu.AddMenuItem(submenu)

	submenu = gogui.CreateTextMenuItem("File")
		item = gogui.CreateTextMenuItem("Open...")
		submenu.AddMenuItem(item)
		menu.AddMenuItem(submenu)

	gogui.SetMainMenu(menu)

	window := gogui.CreateWindow()
	window.SetPosition(
		gogui.Pos(50, 0),
		gogui.Pos(10, 0),
		gogui.Pos(75, 0),
		gogui.Pos(50, 0))
	window.HandleClose(closeHandler)

	scroll := gogui.CreateScrollBox()
	scroll.SetPosition(
		gogui.Pos(0,100),
		gogui.Pos(0,50),
		gogui.Pos(100,0),
		gogui.Pos(100,0))

	dx := int(float64(mySchem.Settings.PageWidth)*zoom)
	dy := int(float64(mySchem.Settings.PageHeight)*zoom)
	scroll.SetContentSize(dx, dy)
	scroll.HandleRedraw(drawHandler)
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
	button.HandleClick(clickHandler)
	window.AddChild(button)

	btnPos += btnHeight
	button = gogui.CreateTextButton("Button 2")
	fmt.Printf("button height = %d\n", btnHeight)
	button.SetPosition(
		gogui.Pos(0, 0),
		gogui.Pos(0, btnPos),
		gogui.Pos(0, 100),
		gogui.Pos(0, btnPos+btnHeight))
	button.HandleClick(clickHandler)
	window.AddChild(button)

	btnPos += btnHeight
	button = gogui.CreateTextButton("Button 3")
	fmt.Printf("button height = %d\n", btnHeight)
	button.SetPosition(
		gogui.Pos(0, 0),
		gogui.Pos(0, btnPos),
		gogui.Pos(0, 100),
		gogui.Pos(0, btnPos+btnHeight))
	button.HandleClick(clickHandler)
	window.AddChild(button)

	window.Show()

	ret := gogui.RunEventLoop()
	fmt.Println(ret)
}


type DrawingContext struct {
	gogui.Graphics
	zoom float64
}

func (dc DrawingContext) setColor(col int) {
	r := (col >> 8) & 0xF
	g := (col >> 4) & 0xF
	b := col& 0xF

	r = r | (r<<4)
	g = g | (g<<4)
	b = b | (b<<4)
	dc.SetStrokeColor(gogui.Color{uint8(r), uint8(g), uint8(b), 255})
}

func (dc DrawingContext) Line(x0 int, y0 int, x1 int, y1 int, color int, width int) {

	dc.setColor(color)
	dc.SetLineWidth(float64(width)*dc.zoom)

	dc.StartPath(float64(x0)*dc.zoom, float64(y0)*dc.zoom)
	dc.LineTo(float64(x1)*dc.zoom, float64(y1)*dc.zoom)
	dc.StrokePath()

}

func (dc DrawingContext) Text(txt string, x0 int, y0 int, rotate int, anchor int, color int, size int, font int) {

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

	x := float64(x0)*dc.zoom - width*math.Cos(rad)
	y := float64(y0)*dc.zoom - width*math.Sin(rad)

	dc.DrawText(x, y, float64(rotate), txt)
}

func (dc DrawingContext) Curve(x0 int, y0 int, cx0 int, cy0 int, cx1 int, cy1 int, x1 int, y1 int, color int, width int) {

	dc.setColor(color)
	dc.SetLineWidth(float64(width)*dc.zoom)
	dc.StartPath(float64(x0)*dc.zoom, float64(y0)*dc.zoom)
	dc.CurveTo(float64(cx0)*dc.zoom, float64(cy0)*dc.zoom, float64(cx1)*dc.zoom, float64(cy1)*dc.zoom, float64(x1)*dc.zoom, float64(y1)*dc.zoom)
	dc.StrokePath()
}

