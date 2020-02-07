package schematic

const (
	// text string anchor positions
	ANCHOR_START  = 1
	ANCHOR_MIDDLE = 2
	ANCHOR_END    = 3

	// rotation and morroring options
	ROTATE_0    = 0
	ROTATE_90   = 1
	ROTATE_180  = 2
	ROTATE_270  = 3
	ROTATE_0M   = 4
	ROTATE_90M  = 5
	ROTATE_180M = 6
	ROTATE_270M = 7

	// font selections
	FONT_MONO     = 4
	FONT_MONO_I   = 5
	FONT_MONO_B   = 6
	FONT_MONO_IB  = 7
	FONT_SANS     = 8
	FONT_SANS_I   = 9
	FONT_SANS_B   = 10
	FONT_SANS_IB  = 11
	FONT_SERIF    = 12
	FONT_SERIF_I  = 13
	FONT_SERIF_B  = 14
	FONT_SERIF_IB = 15

	// standard color values
	COLOR_DEFAULT    = 0x0000
	COLOR_GRAPHICS   = 0x0001
	COLOR_TEXT       = 0x0002
	COLOR_PART_LINES = 0x0003
	COLOR_PART_TEXT  = 0x0004
	COLOR_PIN_LINES  = 0x0005
	COLOR_PIN_TEXT   = 0x0006
	COLOR_WIRES      = 0x0007
	COLOR_BUSSES     = 0x0008
	RGB_BLACK        = 0x1000
	RGB_RED          = 0x1F00
	RGB_GREEN        = 0x10F0
	RGB_BLUE         = 0x100F
)

type Schematic struct {
	Settings    SettingsContainer
	Definitions DefinitionsContainer
	Background  BackgroundContainer
	Pages       []*Page
}

type SettingsContainer struct {
	LengthUnit   string // "cm" "inch"
	TicksPerUnit int    // usually 120/inch or 48/cm
	PageWidth    int    // e.g. 1320x1020 for 11x8.5in
	PageHeight   int
	Properties   map[string]string
}

type DefinitionsContainer struct {
	Symbols map[string]*SymbolDefinition
}

type BackgroundContainer struct {
	Graphics    []*GraphicMark
	Symbols     []*Symbol
	Annotations []*Annotation
}

type Page struct {
	Title    string
	Wires    []*Wire
	Busses   []*Bus
	Graphics []*GraphicMark
	Symbols  []*Symbol
}

type SymbolDefinition struct {
	Type               string // "Graphics" "Part" "Net"
	Pins               []*Pin
	Graphics           []*GraphicMark
	Annotations        []*Annotation
	AnnotationDefaults []*Annotation
}

type Route struct {
	X0          int
	Y0          int
	X1          int
	Y1          int
	Color       int           `json:",omitempty"`
	Width       int           `json:",omitempty"`
	Annotations []*Annotation `json:",omitempty"`
}

type Wire struct {
	Route
}

type Bus struct {
	Route
}

type GraphicMark struct {
	Type   string // "Line" "Curve" "Text"
	X0     int
	Y0     int
	X1     int    `json:",omitempty"`
	Y1     int    `json:",omitempty"`
	CX0    int    `json:",omitempty"`
	CY0    int    `json:",omitempty"`
	CX1    int    `json:",omitempty"`
	CY1    int    `json:",omitempty"`
	Width  int    `json:",omitempty"`
	Color  int    `json:",omitempty"`
	Rotate int    `json:",omitempty"`
	Size   int    `json:",omitempty"`
	Anchor int    `json:",omitempty"`
	Font   int    `json:",omitempty"`
	Text   string `json:",omitempty"`
}

type Symbol struct {
	Def         string
	X0          int
	Y0          int
	Rotate      int           `json:",omitempty"`
	Annotations []*Annotation `json:",omitempty"`
}

type Annotation struct {
	Key    string
	Value  string
	Vis    bool `json:",omitempty"`
	DX     int  `json:",omitempty"`
	DY     int  `json:",omitempty"`
	Rotate int  `json:",omitempty"`
	Size   int  `json:",omitempty"`
	Anchor int  `json:",omitempty"`
	Font   int  `json:",omitempty"`
	Color  int  `json:",omitempty"`
}

type Pin struct {
	Route
}
