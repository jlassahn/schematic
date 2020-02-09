
package main

import (
	"github.com/jlassahn/gogui"
)

var fonts []gogui.Font

type View interface {
	Close() error
}

func InitShared() {

	gogui.Init()

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
	}

}

func ExitShared() {
	gogui.Exit()
}

func QuitApp() error {

	gogui.StopEventLoop(0)
	return nil
}

