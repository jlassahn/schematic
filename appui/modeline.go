
package appui

import (
	"fmt"

	"github.com/jlassahn/schematic"
)


type ModeLine struct {
	schemEdit SchemEdit
}

func CreateModeLine(sch SchemEdit) MouseMode {

	ret := ModeLine{}
	ret.schemEdit = sch

	return &ret
}


func (mode *ModeLine) MouseMove(x int, y int) {
	fmt.Printf("LINE move %v, %v\n", x, y)

	ovr := mode.schemEdit.GetEditOverlay()
	if len(ovr.Graphics) > 0 {
		ovr.Graphics[0].X1 = x
		ovr.Graphics[0].Y1 = y
		//FIXME color, width
	}
}

func (mode *ModeLine) MouseDown(x int, y int, btn int) {

	ovr := mode.schemEdit.GetEditOverlay()
	if len(ovr.Graphics) > 0 {

		g := ovr.Graphics[0]
		ovr.Graphics = nil

		g.X1 = x
		g.Y1 = y

		cmd := CmdCreateGraphics {
			Page: mode.schemEdit.GetPageNumber(),
			NewGraphics: g,
		}
		undo := mode.schemEdit.GetUndoBuffer()
		undo.Do( []UndoElement{ &cmd } )

	} else {
		line := schematic.GraphicMark {
			Type: "Line",
			X0: x,
			Y0: y,
			X1: x,
			Y1: y,
			//FIXME color, width
		}
		ovr.Graphics = []*schematic.GraphicMark{ &line }
	}
}

func (mode *ModeLine) MouseUp(x int, y int, btn int) {
	fmt.Println("LINE up")
}

func (mode *ModeLine) Cancel() {
	fmt.Println("LINE cancel")

	ovr := mode.schemEdit.GetEditOverlay()
	ovr.Graphics = nil
}

