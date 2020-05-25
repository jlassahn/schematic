
package appui

import (
	//"fmt"

	"github.com/jlassahn/schematic"
)

type UndoElement interface {
	Do(schem *schematic.Schematic)
	Undo(schem *schematic.Schematic)
}

type UndoBuffer interface {
	Do(el []UndoElement)
	Undo()
	Redo()
}

func CreateUndoBuffer(schem *schematic.Schematic) UndoBuffer {
	return &undoBuffer{
		schem: schem,
	}
}

type undoBuffer struct {
	schem *schematic.Schematic
	buffer [][]UndoElement
	pos int
}

func (undo *undoBuffer) Do(el []UndoElement) {

	if undo.pos < len(undo.buffer) {
		undo.buffer = undo.buffer[0:undo.pos]
	}
	undo.buffer = append(undo.buffer, el)
	undo.pos ++

	// FIXME limit maximum history?

	doAll(el, undo.schem)
}

func (undo *undoBuffer) Undo() {
	if undo.pos == 0 {
		return
	}
	undo.pos --
	undoAll(undo.buffer[undo.pos], undo.schem)
}

func (undo *undoBuffer) Redo() {
	if undo.pos >= len(undo.buffer) {
		return
	}
	doAll(undo.buffer[undo.pos], undo.schem)
	undo.pos ++
}

func doAll(el []UndoElement, schem *schematic.Schematic) {

	for i:=0; i< len(el); i++ {
		el[i].Do(schem)
	}
}

func undoAll(el []UndoElement, schem *schematic.Schematic) {

	for i:=len(el) - 1; i>= 0; i-- {
		el[i].Undo(schem)
	}
}


type CmdCreateGraphics struct {
	Page int
	NewGraphics *schematic.GraphicMark
}

func (cmd *CmdCreateGraphics) Do(schem *schematic.Schematic) {

	// make a copy
	g := *cmd.NewGraphics

	schem.Pages[cmd.Page - 1].Graphics = append(
		schem.Pages[cmd.Page - 1].Graphics, &g)

}

func (cmd *CmdCreateGraphics) Undo(schem *schematic.Schematic) {

	n := len(schem.Pages[cmd.Page - 1].Graphics)
	schem.Pages[cmd.Page - 1].Graphics =
		schem.Pages[cmd.Page - 1].Graphics[0: n-1]
}

