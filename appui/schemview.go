
package appui

import (
)

func CreateSchemView(win Window, schembox SchemBox) SchemView {

	ret := schemView{}
	ret.window = win
	ret.schembox = schembox

	ret.modeSelect = CreateModeSelect(&ret) // FIXME viewer modeSelect should be more limited than editor modeSelect
	ret.schembox.SetMode(ret.modeSelect)

	return &ret
}

func CreateSchemEdit(win Window, schembox SchemBox) SchemEdit {

	ret := schemEdit{}
	ret.window = win
	ret.schembox = schembox

	ret.modeSelect = CreateModeSelect(&ret)
	ret.modeLine = CreateModeLine(&ret)
	ret.schembox.SetMode(ret.modeSelect)

	return &ret
}

type schemView struct {
	window Window
	schembox SchemBox
	modeSelect MouseMode
}

type schemEdit struct {
	schemView
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


func (sch *schemEdit) ModeSelect() {
}

func (sch *schemEdit) ModeMakeLines() {
}

