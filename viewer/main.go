
package main

import (
	"github.com/jlassahn/gogui"
)

// FIXME refactor with RunOpenDialog
func handlerOpenFile(name string) error {
	_,err := CreateSchematicViewFromFile(name)
	if err != nil {
		//FIXME display error
		return err
	}

	CloseSplashView()
	return err
}

func main() {

	InitShared()
	defer ExitShared()

	gogui.HandleAppOpenFile(handlerOpenFile)

	CreateSplashWindow()

	gogui.RunEventLoop()
}

