# cfmt
ANSI colours and basic movement.

## Functions

### Print()
Print() takes special fmt-style parameters in the form of keywords to style text/backgrounds and move the cursor.

Example:
```go
	cfmt.Print("%blue %bgwhite Blue%normal\tNormal\n")
```

Note that keywords may be followed by a space when followed by letters. This extra space won't be printed.

### Printf()
Printf() takes the keywords from Print and applies its styling, then parses it through fmt.Printf().
```go
	cfmt.Printf("%red%t%normal\n", false)
```

### Format()
Returns a formatted string directly, allowing you to further process it.

Example:
```go
	s := cfmt.Format("%green Green")
```

### TermSize()
This utility function returns the width and height of the current terminal in number of characters:
```go
	w,h,err := cfmt.TermSize()
	…
```

### Pos()
This function sets the position in the terminal window to print from. Best combined with %cls.

## Keywords

### Text colours
- black
- red
- green
- yellow
- blue
- magenta
- cyan
- white
- grey
- lred
- lgreen
- lyellow
- lblue
- lmagenta
- lcyan
- lwhite

### Background colours
- bgblack
- bgred
- bggreen
- bgyellow
- bgblue
- bgmagenta
- bgcyan
- bgwhite
- bggrey
- bglred
- bglgreen
- bglyellow
- bglblue
- bglmagenta
- bglcyan
- bglwhite

### Styling and specials
- reset: Back to normal colour and style
- bold
- fuzzy: May simply be grey
- italic
- under
- blink: May not work (thank goodness)
- fast: Fastblink; see blink
- reverse: Mighr appear as black on grey.
- conceal: May not do anything.
- strike

- clre: Clear to end of line
- save: Store cursor position
- rest: Restore cursor position; useful in progress printouts
- cls: Clear screen
- hide: Hide cursor
- show: Show cursor
