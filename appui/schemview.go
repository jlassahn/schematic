
package appui

import (
	"github.com/jlassahn/schematic"
)

// FIXME pass SchemViewState instead of lots of params
func CreateSchemView(win Window, schembox SchemBox, schem *schematic.Schematic) SchemView{

	ret := schemView{}

	ret.window = win
	ret.schembox = schembox
	ret.schem = schem

	ret.modeSelect = CreateModeSelect(&ret) // FIXME viewer modeSelect should be more limited than editor modeSelect
	ret.schembox.SetMode(ret.modeSelect)

	return &ret
}

func CreateSchemEdit(state *SchemEditState) SchemEdit {

	ret := schemEdit{}

	ret.editState = state
	// FIXME include state by reference instead of copying
	ret.window = state.MainWindow
	ret.schembox = state.SchemBox
	ret.schem = state.Schem
	ret.undoBuffer = CreateUndoBuffer(state.Schem)

	ret.modeSelect = CreateModeSelect(&ret)
	ret.modeLine = CreateModeLine(&ret, state)
	ret.schembox.SetMode(ret.modeSelect)

	return &ret
}

type schemView struct {
	window Window
	schembox SchemBox
	schem *schematic.Schematic

	modeSelect MouseMode
}

type schemEdit struct {
	schemView
	editState *SchemEditState
	undoBuffer UndoBuffer

	modeLine MouseMode
}


func (sch *schemView) Close() {
	if sch.window.Close() == nil {
		windowListRemove(sch.window)
		splashWindow = nil
	}

}

func (sch *schemView) ZoomIn() {
	sch.schembox.ZoomIn()
}

func (sch *schemView) ZoomOut() {
	sch.schembox.ZoomOut()
}

func (sch *schemView) GetPageNumber() int {
	return sch.schembox.GetPageNumber()
}

func (sch *schemEdit) ModeSelect() {
	sch.schembox.SetMode(sch.modeSelect)
}

func (sch *schemEdit) ModeMakeLines() {
	sch.schembox.SetMode(sch.modeLine)
}

func (sch *schemEdit) Undo() {
	sch.undoBuffer.Undo()
	sch.schembox.Redraw()
}

func (sch *schemEdit) Redo() {
	sch.undoBuffer.Redo()
	sch.schembox.Redraw()
}

func (sch *schemEdit) GetEditOverlay() *schematic.Overlay {
	return sch.schembox.GetEditOverlay()
}

func (sch *schemEdit) GetUndoBuffer() UndoBuffer {
	return sch.undoBuffer
}

