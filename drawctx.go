package schematic

type DrawingContext interface {
	Line(x0 int, y0 int, x1 int, y1 int, color int, width int)
	Text(txt string, x0 int, y0 int, rotate int, anchor int, color int, size int, font int)
	Curve(x0 int, y0 int, cx0 int, cy0 int, cx1 int, cy1 int, x1 int, y1 int, color int, width int)
}
