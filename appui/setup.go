
package appui

import (
	"github.com/jlassahn/gogui"
	"github.com/jlassahn/schematic"
)

var fonts []gogui.Font
var openDialog gogui.FileDialog
var saveDialog gogui.FileDialog

var windowCount int = 0
var windowSet map[Window]bool
var splashWindow Window
var mainui MainUI


func Init() {

	gogui.Init()

	windowSet = map[Window]bool {}

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
	saveDialog = gogui.CreateSaveFileDialog()
}

func Exit() {
	openDialog.Destroy()
	gogui.Exit()
}

func StartApp(ui MainUI) {
	mainui = ui
	gogui.HandleAppOpenFile(handleOpenFile)

	splashWindow = mainui.CreateSplashWindow()
	windowListAdd(splashWindow)
}

func Quit() {
	for win := range windowSet {
		if win.Close() == nil {
			windowListRemove(win)
		}
	}
}

func RunOpenDialog() {

	var name string

	if openDialog.Run() {
		name = openDialog.GetFile()
	} else {
		return
	}

	handleOpenFile(name)
}

func NewEmptySchematic() {

	schem := schematic.Schematic {}

	schem.Settings.LengthUnit = "inch"
	schem.Settings.TicksPerUnit = 120
	schem.Settings.PageWidth = 1320
	schem.Settings.PageHeight = 1020

	page := schematic.Page {}
	schem.Pages = []*schematic.Page{&page}

	win := mainui.CreateSchematicWindow(&schem)
	windowListAdd(win)

	closeSplashWindow()
}

func RunSaveAsDialog(schem *schematic.Schematic) {

	var name string

	if saveDialog.Run() {
		name = saveDialog.GetFile()
	} else {
		return
	}

	err := schem.Save(name)
	if err != nil {
		//FIXME show error!
	}
}

func TryToClose(win Window) {
	if win.Close() == nil {
		windowListRemove(win)
	}
}

func windowListRemove(win Window) {
	windowCount --
	delete(windowSet, win)
	if windowCount <= 0 {
		gogui.StopEventLoop(0)
	}
}

func windowListAdd(win Window) {
	windowCount ++
	windowSet[win] = true
}

func handleOpenFile(name string) error {

	// FIXME schematic or library

	schem, err := schematic.LoadSchematic(name)
	if err != nil {
		return err
	}

	win := mainui.CreateSchematicWindow(schem)
	windowListAdd(win)

	closeSplashWindow()

	return nil
}

func closeSplashWindow() {

	if splashWindow == nil {
		return
	}

	if splashWindow.Close() == nil {
		windowListRemove(splashWindow)
		splashWindow = nil
	}
}

