
package main

import (
	"fmt"

	"github.com/jlassahn/gogui"
)

func handlerOpenFile(name string) error {
	fmt.Println("FILE OPEN "+name)
	_,err := CreateSchematicViewFromFile(name)
	return err
}

func main() {

	InitShared()
	defer ExitShared()

	gogui.HandleAppOpenFile(handlerOpenFile)

	myView,_ := CreateSplashWindow()
	//CreateSchematicViewFromFile("tschem.json")
	_ = myView

	ret := gogui.RunEventLoop()
	fmt.Println(ret)
}

