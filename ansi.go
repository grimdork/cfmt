package cfmt

const (
	//
	// Foreground
	//

	black   = "\x1b[30;1m"
	red     = "\x1b[31;1m"
	green   = "\x1b[32;1m"
	yellow  = "\x1b[33;1m"
	blue    = "\x1b[34;1m"
	magenta = "\x1b[35;1m"
	cyan    = "\x1b[36;1m"
	white   = "\x1b[37;1m"

	grey         = "\x1b[90;1m"
	lightRed     = "\x1b[91;1m"
	lightGreen   = "\x1b[92;1m"
	lightYellow  = "\x1b[93;1m"
	lightBlue    = "\x1b[94;1m"
	lightMagenta = "\x1b[95;1m"
	lightCyan    = "\x1b[96;1m"
	lightWhite   = "\x1b[97;1m"

	//
	// Background
	//

	bgBlack   = "\x1b[40;5m"
	bgRed     = "\x1b[41;5m"
	bgGreen   = "\x1b[42;5m"
	bgYellow  = "\x1b[43;5m"
	bgBlue    = "\x1b[44;5m"
	bgMagenta = "\x1b[45;5m"
	bgCyan    = "\x1b[46;5m"
	bgWhite   = "\x1b[47;5m"

	bgGrey         = "\x1b[100;5m"
	bgLightRed     = "\x1b[101;5m"
	bgLightGreen   = "\x1b[102;5m"
	bgLightYellow  = "\x1b[103;5m"
	bgLightBlue    = "\x1b[104;5m"
	bgLightMagenta = "\x1b[105;5m"
	bgLightCyan    = "\x1b[106;5m"
	bgLightWhite   = "\x1b[107;5m"

	//
	// Other options
	//

	reset         = "\x1b[0m"
	bold          = "\x1b[1m"
	fuzzy         = "\x1b[2m"
	italic        = "\x1b[3m"
	underscore    = "\x1b[4m"
	blink         = "\x1b[5m"
	fastBlink     = "\x1b[6m"
	reverse       = "\x1b[7m"
	concealed     = "\x1b[8m"
	strikethrough = "\x1b[9m"

	clearEOL    = "\x1b[K"
	savePos     = "\x1b[s"
	restorePos  = "\x1b[u"
	clearScreen = "\x1b[2J"
	hideCursor  = "\x1b[?25l"
	showCursor  = "\x1b[?25h"
)
