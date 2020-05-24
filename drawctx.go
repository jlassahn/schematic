package schematic

const (
	WIRE     = 1
	BUS      = 2
	PIN      = 3
	GRAPHICS = 4
	MAX_MODE = 5
)

type DrawingDriver interface {
	Line(x0 int, y0 int, x1 int, y1 int, color int, width int)
	Text(txt string, x0 int, y0 int, rotate int, anchor int, color int, size int, font int)
	Curve(x0 int, y0 int, cx0 int, cy0 int, cx1 int, cy1 int, x1 int, y1 int, color int, width int)
}

type DrawingContext interface {
	DrawingDriver
	MapColor(col int, mode int) int
	MapWidth(width int, mode int) int
}

type DrawingSettings struct {
	Colors [MAX_MODE]int
	Widths [MAX_MODE]int
}

type drawCtx struct {
	driver DrawingDriver
	settings *DrawingSettings
}

func WrapDrawingDriver(driver DrawingDriver, settings *DrawingSettings) DrawingContext {

	ret := drawCtx {
		driver: driver,
		settings: settings,
	}

	return &ret
}


func (ctx *drawCtx) Line(x0 int, y0 int, x1 int, y1 int, color int, width int) {
	ctx.driver.Line(x0, y0, x1, y1, color, width)
}

func (ctx *drawCtx) Text(txt string, x0 int, y0 int, rotate int, anchor int, color int, size int, font int) {
	ctx.driver.Text(txt, x0, y0, rotate, anchor, color, size, font)
}

func (ctx *drawCtx) Curve(x0 int, y0 int, cx0 int, cy0 int, cx1 int, cy1 int, x1 int, y1 int, color int, width int) {
	ctx.driver.Curve(x0, y0, cx0, cy0, cx1, cy1, x1, y1, color, width)
}

func (ctx *drawCtx) MapColor(col int, mode int) int {
	if col&0x1000 == 0x1000 {
		return col & 0xFFF
	}

	if col > 0 && col < MAX_MODE {
		return ctx.settings.Colors[col]
	}
	return ctx.settings.Colors[mode]
}

func (ctx *drawCtx) MapWidth(w int, mode int) int {
	if w <= 0 {
		return ctx.settings.Widths[mode]
	}
	return w
}

