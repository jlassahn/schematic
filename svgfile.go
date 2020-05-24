package schematic

import (
	"fmt"
	"io"
	"os"
)

var FONT_NAMES = [16]string{
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
}

func (schem *Schematic) ExportSVG(settings *DrawingSettings, filename string, page int) error {

	drv, err := newSVGFile(filename, schem.Settings.PageWidth, schem.Settings.PageHeight)
	if err != nil {
		return err
	}

	ctx := WrapDrawingDriver(drv, settings)

	DrawGrid(ctx, schem.Settings.PageWidth, schem.Settings.PageHeight)
	schem.DrawPage(ctx, page)

	finalize(drv)
	return nil
}

type SVGFile struct {
	io.WriteCloser
}

func newSVGFile(name string, width int, height int) (SVGFile, error) {

	fp, err := os.Create(name)
	if err != nil {
		return SVGFile{nil}, err
	}

	fmt.Fprintf(fp, "<?xml version=\"1.0\"?>\n")
	fmt.Fprintf(fp, "<svg width=\"%v\" height=\"%v\" version=\"1.1\" xmlns=\"http://www.w3.org/2000/svg\" stroke-linecap=\"round\" font-family=\"mono\">\n", width, height)
	fmt.Fprintln(fp,
		`	<style>
	@font-face { font-family: "mono_n"; src: local("DejaVu Sans Mono"); }
	@font-face { font-family: "mono_i"; src: local("DejaVu Sans Mono Oblique"); }
	@font-face { font-family: "mono_b"; src: local("DejaVu Sans Mono Bold"); }
	@font-face { font-family: "mono_bi"; src: local("DejaVu Sans Mono Bold Oblique"); }
	@font-face { font-family: "sans_n"; src: local("DejaVu Sans"); }
	@font-face { font-family: "sans_i"; src: local("DejaVu Sans Oblique"); }
	@font-face { font-family: "sans_b"; src: local("DejaVu Sans Bold"); }
	@font-face { font-family: "sans_bi"; src: local("DejaVu Sans Bold Oblique"); }
	@font-face { font-family: "serif_n"; src: local("DejaVu Serif"); }
	@font-face { font-family: "serif_i"; src: local("DejaVu Serif Italic"); }
	@font-face { font-family: "serif_b"; src: local("DejaVu Serif Bold"); }
	@font-face { font-family: "serif_bi"; src: local("DejaVu Serif Bold Italic"); }
	}
	</style>
`)

	return SVGFile{fp}, nil
}

func finalize(fp SVGFile) {

	fmt.Fprintf(fp, "</svg>\n")
	fp.Close()
}

func (svg SVGFile) Line(x0 int, y0 int, x1 int, y1 int, color int, width int) {
	fmt.Fprintf(svg, "<line x1=\"%v\" y1=\"%v\" x2=\"%v\" y2=\"%v\" stroke=\"#%.3X\" stroke-width=\"%v\" />\n",
		x0, y0, x1, y1, color, width)
}

func (svg SVGFile) Text(txt string, x0 int, y0 int, rotate int, anchor int, color int, size int, font int) {

	anchor_str := "middle"
	if anchor == ANCHOR_START {
		anchor_str = "start"
	} else if anchor == ANCHOR_END {
		anchor_str = "end"
	}

	fmt.Fprintf(svg, "<text transform=\"translate(%v,%v) rotate(%v)\" font-size=\"%v\" fill=\"#%.3X\" text-anchor=\"%v\" font-family=\"%v\">%v</text>\n",
		x0, y0, rotate, size, color, anchor_str, FONT_NAMES[font], txt)
}

func (svg SVGFile) Curve(x0 int, y0 int, cx0 int, cy0 int, cx1 int, cy1 int, x1 int, y1 int, color int, width int) {
	fmt.Fprintf(svg, "<path d=\"M %v %v C %v %v %v %v %v %v\" fill=\"none\" stroke=\"#%.3X\" stroke-width=\"%v\" />\n",
		x0, y0, cx0, cy0, cx1, cy1, x1, y1, color, width)
}
