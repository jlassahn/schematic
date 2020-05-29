
package main

import (
	"fmt"
	"strconv"

	"github.com/jlassahn/gogui"
	//"github.com/jlassahn/schematic"
	"github.com/jlassahn/schematic/appui"
)

type PropertyBox struct {

	EditState *appui.SchemEditState
	FrameBox gogui.Box
	ColorInput gogui.TextInput
	WidthInput gogui.TextInput
}

func CreatePropertyBox(state *appui.SchemEditState) *PropertyBox {

	ret := PropertyBox{}
	ret.EditState = state

	ret.FrameBox = gogui.CreateBox()

	ret.ColorInput = gogui.CreateTextLineInput()
	ret.ColorInput.SetText("0x0")
	ret.ColorInput.HandleChange(ret.handleColorChange)
	ret.ColorInput.SetPosition(
		gogui.Pos(0,0),
		gogui.Pos(0,50),
		gogui.Pos(100,0),
		gogui.Pos(0,50 + ret.ColorInput.GetBestHeight()))
	ret.WidthInput = gogui.CreateTextLineInput()
	ret.WidthInput.SetText("0")
	ret.WidthInput.HandleChange(ret.handleWidthChange)
	ret.WidthInput.SetPosition(
		gogui.Pos(0,0),
		gogui.Pos(0,50 + ret.ColorInput.GetBestHeight()),
		gogui.Pos(100,0),
		gogui.Pos(0,50 + 2*ret.ColorInput.GetBestHeight()))
	ret.FrameBox.AddChild(ret.ColorInput)
	ret.FrameBox.AddChild(ret.WidthInput)
	ret.EditState.NewLineWidth = 0
	ret.EditState.NewLineColor = 0x0

	return &ret
}

func (prop *PropertyBox) SetPosition(left, top, right, bottom gogui.Position) {
	prop.FrameBox.SetPosition(left, top, right, bottom)
}

func (prop *PropertyBox) handleColorChange(txt string) {

	val, err := strconv.ParseInt(txt, 16, 32)
	if err != nil {
		val, err = strconv.ParseInt(txt, 0, 32)
	}

	if err != nil {
		prop.ColorInput.SetText(fmt.Sprintf("0x%X", prop.EditState.NewLineColor))
		return
	}

	prop.EditState.NewLineColor = int(val)
}

func (prop *PropertyBox) handleWidthChange(txt string) {

	val, err := strconv.ParseInt(txt, 0, 32)

	if err != nil {
		prop.WidthInput.SetText(fmt.Sprintf("%d", prop.EditState.NewLineWidth))
		return
	}

	prop.EditState.NewLineWidth = int(val)
}

