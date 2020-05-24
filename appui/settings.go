
package appui

import (
	"github.com/jlassahn/schematic"
)


type Settings struct {
	DrawSettings *schematic.DrawingSettings
}

var drawSettings = schematic.DrawingSettings {
	Colors: [schematic.MAX_MODE]int {
		0x000, // not used
		0x00F, // WIRE
		0x080, // BUS
		0xF00, // PIN
		0x000, // GRAPHICS
	},
	Widths: [schematic.MAX_MODE]int {
		1, // not used
		1, // WIRE
		3, // BUS
		1, // PIN
		1, // GRAPHICS
	},
}

var globalSettings = Settings {
	DrawSettings: &drawSettings,
}

var GlobalSettings = &globalSettings

