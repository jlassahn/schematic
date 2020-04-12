
package appui

import (
	"fmt"
)


type ModeSelect struct {
	schemview SchemView
}

func CreateModeSelect(sch SchemView) MouseMode {

	ret := ModeSelect{}
	ret.schemview = sch

	return &ret
}


func (mode *ModeSelect) MouseMove(x int, y int) {
	fmt.Printf("SELECT move %v, %v\n", x, y)
}

func (mode *ModeSelect) MouseDown(x int, y int, btn int) {
	fmt.Println("SELECT down")
}

func (mode *ModeSelect) MouseUp(x int, y int, btn int) {
	fmt.Println("SELECT up")
}

func (mode *ModeSelect) Cancel() {
	fmt.Println("SELECT cancel")
}

