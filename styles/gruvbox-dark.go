package styles

import (
	"github.com/alecthomas/chroma"
)

// GruvboxDark Variables
const (
	bg0h = "#1d2021"
	bg0  = "#282828"
	bg0s = "#32302f"
	bg1  = "#3c3836"
	bg2  = "#504945"
	bg3  = "#665c54"
	bg4  = "#7c6f64"

	gray = "#928374"

	fg0 = "#fbf1c7"
	fg1 = "#ebdbb2"
	fg2 = "#d5c4a1"
	fg3 = "#bdae93"
	fg4 = "#a89984"

	red      = "#cc241d"
	redLt    = "#fb4934"
	green    = "#98971a"
	greenLt  = "#b8bb26"
	yellow   = "#d79921"
	yellowLt = "#fabd2f"
	blue     = "#458588"
	blueLt   = "#83a598"
	purple   = "#b16286"
	purpleLt = "#d3869b"
	aqua     = "#689d6a"
	aquaLt   = "#8ec07c"
	orange   = "#d65d0e"
	orangeLt = "#fe8019"
)

// GruvboxDark Style
var GruvboxDark = Register(chroma.MustNewStyle("gruvbox-dark", chroma.StyleEntries{
	chroma.Comment:    gray,
	chroma.Error:      red + " bold",
	chroma.Keyword:    red,
	chroma.Text:       fg1,
	chroma.Background: " bg:" + bg0,
}))
