
package appui

import (
)

func CreateSchemView(win Window, schembox SchemBox) SchemView {

	ret := schemView{}
	ret.window = win
	ret.schembox = schembox

	ret.modeselect = CreateModeSelect(&ret)
	ret.schembox.SetMode(ret.modeselect)

	return &ret
}

func CreateSchemEdit(win Window, schembox SchemBox) SchemEdit {

	ret := schemEdit{}
	ret.window = win
	ret.schembox = schembox

	ret.modeselect = CreateModeSelect(&ret)
	ret.schembox.SetMode(ret.modeselect)

	return &ret
}

type schemView struct {
	window Window
	schembox SchemBox
	modeselect MouseMode
}

type schemEdit struct {
	schemView
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


func (sch *schemEdit) ModeSelect() {
}

func (sch *schemEdit) ModeMakeLines() {
}

