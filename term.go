package cfmt

import (
	"fmt"
	"os"

	"golang.org/x/sys/unix"
)

// TermSize returns the width and height of the default terminal.
func TermSize() (width, height int, err error) {
	f, err := os.Open("/dev/tty")
	if err != nil {
		return 0, 0, err
	}

	defer f.Close()
	sz, err := unix.IoctlGetWinsize(int(f.Fd()), unix.TIOCGWINSZ)
	if err != nil {
		return 0, 0, err
	}
	return int(sz.Col), int(sz.Row), nil
}

// Pos returns the appropriate escape code to move the cursor to the specified position in the terminal.
func Pos(x, y int) string {
	return fmt.Sprintf("\x1b[%d;%dH", y, x)
}
