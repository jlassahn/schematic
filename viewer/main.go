
package main

import (
	"github.com/jlassahn/gogui"
)

func handlerOpenFile(name string) error {
	_,err := CreateSchematicViewFromFile(name)
	return err
}

func main() {

	InitShared()
	defer ExitShared()

	gogui.HandleAppOpenFile(handlerOpenFile)

	CreateSplashWindow()

	gogui.RunEventLoop()
}

