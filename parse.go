package cfmt

import (
	"fmt"
	"strings"
	"unicode"
)

// Print with colours, but no other formatting.
func Print(s string) {
	print(Format(s))
}

// Printf with colours uses fmt.Printf().
func Printf(s string, v ...interface{}) {
	f := Format(s)
	fmt.Printf(f, v...)
}

// Format returns a string with styling applied.
func Format(f string) string {
	dst := strings.Builder{}
	for len(f) > 0 {
		c := f[0]
		if c == '%' {
			var key string
			key, f = parseKeyword(f)
			switch key {
			// Text colour
			case "black":
				dst.Write([]byte(black))
			case "red":
				dst.Write([]byte(red))
			case "green":
				dst.Write([]byte(green))
			case "yellow":
				dst.Write([]byte(yellow))
			case "blue":
				dst.Write([]byte(blue))
			case "magenta":
				dst.Write([]byte(magenta))
			case "cyan":
				dst.Write([]byte(cyan))
			case "white":
				dst.Write([]byte(white))

			case "grey":
				dst.Write([]byte(grey))
			case "lred":
				dst.Write([]byte(lightRed))
			case "lgreen":
				dst.Write([]byte(lightGreen))
			case "lyellow":
				dst.Write([]byte(lightYellow))
			case "lblue":
				dst.Write([]byte(lightBlue))
			case "lmagenta":
				dst.Write([]byte(lightMagenta))
			case "lcyan":
				dst.Write([]byte(lightCyan))
			case "lwhite":
				dst.Write([]byte(lightWhite))

			// Background colour
			case "bgblack":
				dst.Write([]byte(bgBlack))
			case "bgred":
				dst.Write([]byte(bgRed))
			case "bggreen":
				dst.Write([]byte(bgGreen))
			case "bgyellow":
				dst.Write([]byte(bgYellow))
			case "bgblue":
				dst.Write([]byte(bgBlue))
			case "bgmagenta":
				dst.Write([]byte(bgMagenta))
			case "bgcyan":
				dst.Write([]byte(bgCyan))
			case "bgwhite":
				dst.Write([]byte(bgWhite))

			case "bggrey":
				dst.Write([]byte(bgGrey))
			case "bglred":
				dst.Write([]byte(bgLightRed))
			case "bglgreen":
				dst.Write([]byte(bgLightGreen))
			case "bglyellow":
				dst.Write([]byte(bgLightYellow))
			case "bglblue":
				dst.Write([]byte(bgLightBlue))
			case "bglmagenta":
				dst.Write([]byte(bgLightMagenta))
			case "bglcyan":
				dst.Write([]byte(bgLightCyan))
			case "bglwhite":
				dst.Write([]byte(bgLightWhite))

				// Other styling
			case "reset":
				dst.Write([]byte(reset))
			case "bold":
				dst.Write([]byte(bold))
			case "fuzzy":
				dst.Write([]byte(fuzzy))
			case "italic":
				dst.Write([]byte(italic))
			case "under":
				dst.Write([]byte(underscore))
			case "blink":
				dst.Write([]byte(blink))
			case "fast":
				dst.Write([]byte(fastBlink))
			case "reverse":
				dst.Write([]byte(reverse))
			case "conceal":
				dst.Write([]byte(concealed))
			case "strike":
				dst.Write([]byte(strikethrough))

			case "clre":
				dst.Write([]byte(clearEOL))
			case "save":
				dst.Write([]byte(savePos))
			case "rest":
				dst.Write([]byte(restorePos))
			case "cls":
				dst.Write([]byte(clearScreen))
			case "hide":
				dst.Write([]byte(hideCursor))
			case "show":
				dst.Write([]byte(showCursor))

			default:
				dst.Write([]byte("%"))
				dst.Write([]byte(key))
			}
		} else {
			dst.Write([]byte{f[0]})
			f = f[1:]
		}
	}

	return dst.String()
}

// parseKeyword returns the parsed keyword and the rest of the input string.
func parseKeyword(f string) (string, string) {
	var b strings.Builder
	if len(f) == 0 {
		return "", f
	}

	b.WriteByte(f[0])
	in := f[1:]
	loop := true
	for len(in) > 0 && loop {
		if !unicode.IsLetter(rune(in[0])) {
			loop = false
			if len(in) > 1 && in[0] == ' ' {
				in = in[1:]
			}
		} else {
			b.WriteByte(in[0])
			in = in[1:]
		}
	}
	return b.String()[1:], in

}
