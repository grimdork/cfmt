package cfmt_test

import (
	"testing"

	"github.com/grimdork/cfmt"
)

func TestPos(t *testing.T) {
	w, h, err := cfmt.TermSize()
	if err != nil {
		t.Errorf("Error getting terminal size: %s", err.Error())
		t.FailNow()
	}

	t.Logf("W: %d\tH: %d", w, h)
}

func TestColour(t *testing.T) {
	cfmt.Print("")
	cfmt.Print("%blue %bgwhite Blue%reset\tNormal\n")
	cfmt.Print("%bold%red%bgmagenta")
	cfmt.Print("Red bold%reset\tNormal\n")
}

func TestStyling(t *testing.T) {
	cfmt.Print("%bold")
	cfmt.Print("Bold%reset\n")
	cfmt.Print("%fuzzy")
	cfmt.Print("Fuzzy%reset\n")
	cfmt.Print("%italic")
	cfmt.Print("Italic%reset\n")
	cfmt.Print("%under")
	cfmt.Print("Underscore%reset\n")
	cfmt.Print("%blink")
	cfmt.Print("Blink%reset\n")
	cfmt.Print("%fast")
	cfmt.Print("Epileptic%reset\n")
	cfmt.Print("%reverse")
	cfmt.Print("Reverse%reset\n")
	cfmt.Print("%conceal")
	cfmt.Print("Concealed%reset\n")
	cfmt.Print("%strike")
	cfmt.Print("Strikethrough%reset\n")
}

func TestFmt(t *testing.T) {
	cfmt.Printf("%red%t%reset\n", false)
	cfmt.Printf("%green%t%reset\n", true)
}

func TestClearEOL(t *testing.T) {
	cfmt.Print("%save")
	cfmt.Print("One%rest Two\n")
}
