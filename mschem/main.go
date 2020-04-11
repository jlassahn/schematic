
package main

import (
	"github.com/jlassahn/gogui"
	"github.com/jlassahn/schematic"
	"github.com/jlassahn/schematic/appui"
)


func main() {

	appui.Init()
	defer appui.Exit()

	mainui := CreateMainUI()
	appui.StartApp(mainui)

	gogui.RunEventLoop()
}

func CreateMainUI() appui.MainUI {
	ret := MSchemUI{}

	return &ret
}

type MSchemUI struct {
}

func (ui *MSchemUI) CreateLibraryWindow(lib *schematic.Library) appui.Window {
	return nil
}


