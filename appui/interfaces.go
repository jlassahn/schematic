
package appui

import (
	"github.com/jlassahn/gogui"
	"github.com/jlassahn/schematic"
)



type Window interface {
	Close() error
}

type MainUI interface {

	CreateSchematicWindow(schem *schematic.Schematic) Window
	CreateLibraryWindow(lib *schematic.Library) Window
	CreateSplashWindow() Window
}

type MouseMode interface {
	MouseMove(x int, y int)
	MouseDown(x int, y int, btn int)
	MouseUp(x int, y int, btn int)
	Cancel()
}

type SchemBox interface {
	Box() gogui.ScrollBox
	ZoomIn()
	ZoomOut()
	SetMode(mode MouseMode)
}

type SchemView interface {
	Close()
	ZoomIn()
	ZoomOut()
	// FIXME SaveAs, Print, etc
}

type SchemEdit interface {
	SchemView
	ModeSelect()
	ModeMakeLines()
}

