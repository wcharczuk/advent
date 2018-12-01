package log

// AnsiColor represents an ansi color code fragment.
type AnsiColor string

func (c AnsiColor) escaped() string {
	return "\033[" + string(c)
}

// Apply returns a string with the color code applied.
func (c AnsiColor) Apply(text string) string {
	return c.escaped() + text + ColorReset.escaped()
}

// Named colors
const (
	ColorBlack       AnsiColor = "30m"
	ColorRed         AnsiColor = "31m"
	ColorGreen       AnsiColor = "32m"
	ColorYellow      AnsiColor = "33m"
	ColorBlue        AnsiColor = "34m"
	ColorPurple      AnsiColor = "35m"
	ColorCyan        AnsiColor = "36m"
	ColorWhite       AnsiColor = "37m"
	ColorLightBlack  AnsiColor = "90m"
	ColorLightRed    AnsiColor = "91m"
	ColorLightGreen  AnsiColor = "92m"
	ColorLightYellow AnsiColor = "93m"
	ColorLightBlue   AnsiColor = "94m"
	ColorLightPurple AnsiColor = "95m"
	ColorLightCyan   AnsiColor = "96m"
	ColorLightWhite  AnsiColor = "97m"
	ColorGray        AnsiColor = ColorLightBlack
	ColorReset       AnsiColor = "0m"
)
