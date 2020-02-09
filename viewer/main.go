
package main

import (
	"fmt"

	"github.com/jlassahn/gogui"
)

func main() {

	InitShared()
	defer ExitShared()

	myView,_ := CreateSchematicViewFromFile("tschem.json")
	_ = myView

	ret := gogui.RunEventLoop()
	fmt.Println(ret)
}

