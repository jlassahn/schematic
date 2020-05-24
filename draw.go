package schematic

import (
	"fmt"
)

/* FIXME remove
var STANDARD_COLORS = [5]int{
	0x000,
	0x00F,
	0x080,
	0xF00,
	0x000,
}

var STANDARD_WIDTHS = [5]int{
	1,
	1,
	3,
	1,
	1,
}
*/

func DrawGrid(dc DrawingContext, width int, height int) {
	for x := 12; x < width; x += 12 {
		dc.Line(x, 0, x, height, 0xCCC, 1)
	}
	for y := 12; y < height; y += 12 {
		dc.Line(0, y, width, y, 0xCCC, 1)
	}
}

func (overlay *Overlay) Draw(dc DrawingContext) {
	// FIXME implement
}

func (schem *Schematic) DrawPage(dc DrawingContext, pg int) {

	page := schem.Pages[pg-1]

	for _, wire := range page.Wires {
		wire.Draw(dc)
	}

	for _, bus := range page.Busses {
		bus.Draw(dc)
	}

	for _, graph := range page.Graphics {
		DrawGraphics(dc, graph, 0, 0, 0)
	}

	for _, sym := range page.Symbols {
		sym.Draw(dc, &schem.Definitions)
	}

	/* FIXME later
	for k,v := range page.wirePointMap {
		if len(v) > 2 {
			fmt.Printf("connect (%v, %v)\n", k.x, k.y)
			dc.Line(k.x, k.y, k.x, k.y, 0x00F, 3) //FIXME parameterize color and size
		}
	}
	*/
}

func (wire *Wire) Draw(dc DrawingContext) {

	col := dc.MapColor(wire.Color, WIRE)
	width := dc.MapWidth(wire.Width, WIRE)
	dc.Line(wire.X0, wire.Y0, wire.X1, wire.Y1, col, width)

	for _, ann := range wire.Annotations {
		DrawAnnotation(dc, ann, wire.X0, wire.Y0, 0)
	}
}

func (bus *Bus) Draw(dc DrawingContext) {

	col := dc.MapColor(bus.Color, BUS)
	width := dc.MapWidth(bus.Width, BUS)
	dc.Line(bus.X0, bus.Y0, bus.X1, bus.Y1, col, width)

	for _, ann := range bus.Annotations {
		DrawAnnotation(dc, ann, bus.X0, bus.Y0, 0)
	}
}

func (sym *Symbol) Draw(dc DrawingContext, defs *DefinitionsContainer) {

	def := defs.Symbols[sym.Def]
	def.Draw(dc, sym.X0, sym.Y0, sym.Rotate)

	for _, ann := range sym.Annotations {
		DrawAnnotation(dc, ann, sym.X0, sym.Y0, sym.Rotate)
	}
}

func (sd *SymbolDefinition) Draw(dc DrawingContext, xc int, yc int, rot int) {

	for _, pin := range sd.Pins {
		DrawPin(dc, pin, xc, yc, rot)
	}

	for _, graph := range sd.Graphics {
		DrawGraphics(dc, graph, xc, yc, rot)
	}

	for _, ann := range sd.Annotations {
		DrawAnnotation(dc, ann, xc, yc, rot)
	}
}

func DrawGraphics(dc DrawingContext, graph *GraphicMark, x0 int, y0 int, rot int) {

	switch graph.Type {
	case "Line":
		DrawGraphLine(dc, graph, x0, y0, rot)

	case "Curve":
		DrawGraphCurve(dc, graph, x0, y0, rot)

	case "Text":
		DrawGraphText(dc, graph, x0, y0, rot)

	default:
		fmt.Println("GRAPH UNKNOWN") //FIXME error handling
	}
}

func DrawGraphLine(dc DrawingContext, graph *GraphicMark, xc int, yc int, rot int) {
	x0, y0 := Rotate(rot, graph.X0, graph.Y0)
	x1, y1 := Rotate(rot, graph.X1, graph.Y1)

	col := dc.MapColor(graph.Color, GRAPHICS)
	width := dc.MapWidth(graph.Width, GRAPHICS)
	dc.Line(x0+xc, y0+yc, x1+xc, y1+yc, col, width)
}

func DrawGraphCurve(dc DrawingContext, graph *GraphicMark, xc int, yc int, rot int) {
	x0, y0 := Rotate(rot, graph.X0, graph.Y0)
	cx0, cy0 := Rotate(rot, graph.CX0, graph.CY0)
	x1, y1 := Rotate(rot, graph.X1, graph.Y1)
	cx1, cy1 := Rotate(rot, graph.CX1, graph.CY1)

	col := dc.MapColor(graph.Color, GRAPHICS)
	width := dc.MapWidth(graph.Width, GRAPHICS)
	dc.Curve(x0+xc, y0+yc, cx0+xc, cy0+yc, cx1+xc, cy1+yc, x1+xc, y1+yc, col, width)
}

func DrawGraphText(dc DrawingContext, graph *GraphicMark, xc int, yc int, rot int) {

	x0, y0 := Rotate(rot, graph.X0, graph.Y0)
	col := dc.MapColor(graph.Color, GRAPHICS)

	font := graph.Font
	if font < 4 || font > 15 {
		font = 4
	}

	size := graph.Size
	if size <= 0 {
		size = 16
	}

	flip := (graph.Rotate & 4) != 0
	angle := (graph.Rotate & 3)
	if (rot & 4) != 0 {
		flip = !flip
		angle = 4 - angle
	}
	angle = (angle + rot) & 3
	angle = angle * 90

	anchor := graph.Anchor
	if anchor < 1 || anchor > 3 {
		anchor = 2
	}

	if flip {
		anchor = 4 - anchor
	}

	//FIXME if two or three directions
	if angle == 180 {
		angle = 0
		anchor = 4 - anchor
		y0 += size * 8 / 10 //FIXME hack to estimate ascender height
	}

	//FIXME if two directions
	if angle == 90 {
		angle = 270
		anchor = 4 - anchor
		x0 += size * 8 / 10 //FIXME hack to estimate ascender height
	}

	dc.Text(graph.Text, x0+xc, y0+yc, angle, anchor, col, size, font)
}

func DrawAnnotation(dc DrawingContext, ann *Annotation, xc int, yc int, rot int) {

	if !ann.Vis {
		return
	}

	//FIXME duplicating a whole bunch of stuff from DrawGaphText
	x0, y0 := Rotate(rot, ann.DX, ann.DY)
	col := dc.MapColor(ann.Color, GRAPHICS) //FIXME is this really GRAPHICS?

	font := ann.Font
	if font < 4 || font > 15 {
		font = 4
	}

	size := ann.Size
	if size <= 0 {
		size = 16
	}

	flip := (ann.Rotate & 4) != 0
	angle := (ann.Rotate & 3)
	if (rot & 4) != 0 {
		flip = !flip
		angle = 4 - angle
	}
	angle = (angle + rot) & 3
	angle = angle * 90

	anchor := ann.Anchor
	if anchor < 1 || anchor > 3 {
		anchor = 2
	}

	if flip {
		anchor = 4 - anchor
	}

	//FIXME if two or three directions
	if angle == 180 {
		angle = 0
		anchor = 4 - anchor
		y0 += size * 8 / 10 //FIXME hack to estimate ascender height
	}

	//FIXME if two directions
	if angle == 90 {
		angle = 270
		anchor = 4 - anchor
		x0 += size * 8 / 10 //FIXME hack to estimate ascender height
	}

	dc.Text(ann.Value, x0+xc, y0+yc, angle, anchor, col, size, font)
}

func DrawPin(dc DrawingContext, pin *Pin, xc int, yc int, rot int) {

	x0, y0 := Rotate(rot, pin.X0, pin.Y0)
	x1, y1 := Rotate(rot, pin.X1, pin.Y1)
	col := dc.MapColor(pin.Color, PIN)
	width := dc.MapWidth(pin.Width, PIN)
	dc.Line(x0+xc, y0+yc, x1+xc, y1+yc, col, width)

	for _, ann := range pin.Annotations {
		DrawAnnotation(dc, ann, x0+xc, y0+yc, rot)
	}
}

/* FIXME remove
func MapColor(col int, mode int) int {
	if col&0x1000 == 0x1000 {
		return col & 0xFFF
	}

	if col > 0 && col < MAX_MODE {
		return STANDARD_COLORS[col]
	}
	return STANDARD_COLORS[mode]
}

func MapWidth(w int, mode int) int {
	if w <= 0 {
		return STANDARD_WIDTHS[mode]
	}
	return w
}
*/

func Rotate(rot int, x int, y int) (int, int) {

	switch rot {
	default:
		return x, y
	case ROTATE_90:
		return -y, x
	case ROTATE_180:
		return -x, -y
	case ROTATE_270:
		return y, -x

	case ROTATE_0M:
		return -x, y
	case ROTATE_90M:
		return -y, -x
	case ROTATE_180M:
		return x, -y
	case ROTATE_270M:
		return y, x
	}
}
