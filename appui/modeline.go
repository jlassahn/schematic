
package appui

import (
	"fmt"
)


type ModeLine struct {
	schemview SchemView
}

func CreateModeLine(sch SchemView) MouseMode {

	ret := ModeLine{}
	ret.schemview = sch

	return &ret
}


func (mode *ModeLine) MouseMove(x int, y int) {
	fmt.Printf("LINE move %v, %v\n", x, y)
}

func (mode *ModeLine) MouseDown(x int, y int, btn int) {
	fmt.Println("LINE down")
}

func (mode *ModeLine) MouseUp(x int, y int, btn int) {
	fmt.Println("LINE up")
}

func (mode *ModeLine) Cancel() {
	fmt.Println("LINE cancel")
}

