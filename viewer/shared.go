
package main

import (
	"github.com/jlassahn/gogui"
)

var fonts []gogui.Font
var openDialog gogui.FileDialog

type View interface {
	Close() error
}

func InitShared() {

	gogui.Init()

	viewSet = map[View]bool {}

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

	openDialog = gogui.CreateOpenFileDialog()
}

func ExitShared() {
	openDialog.Destroy()
	gogui.Exit()
}

func QuitApp() {

	for view := range viewSet {
		view.Close()
	}
}

func RunOpenDialog() {

	var name string

	if openDialog.Run() {
		name = openDialog.GetFile()
	} else {
		return
	}

	_,err := CreateSchematicViewFromFile(name)
	if err != nil {
		//FIXME display error
		return
	}

	CloseSplashView()
}

var viewCount int = 0
var viewSet map[View]bool

func ViewListRemove(view View) {
	viewCount --
	delete(viewSet, view)
	if viewCount <= 0 {
		gogui.StopEventLoop(0)
	}
}

func ViewListAdd(view View) {
	viewCount ++
	viewSet[view] = true
}

