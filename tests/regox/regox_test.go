package regox_test

import (
	"regox/internals/regox"
	"testing"
)

func TestRegox(t *testing.T) {
	checkForMe("", "", regox.MATCHED, t)
	checkForMe("r", "", regox.UNMATCHED, t)
	checkForMe("", "r", regox.UNMATCHED, t)
	checkForMe("regox", "regox", regox.MATCHED, t)
	checkForMe("regex", "regox", regox.UNMATCHED, t)
	
	checkForMe("reg?x", "regox", regox.MATCHED, t)
	checkForMe("?", "r", regox.MATCHED, t)
	
	checkForMe("*", "regox.go", regox.MATCHED, t)
	checkForMe("****", "", regox.MATCHED, t)
	checkForMe("regox.*", "regox.go", regox.MATCHED, t)
	checkForMe("*******reg?x.*******", "regex.go", regox.MATCHED, t)
	
	checkForMe("reg[eo]x.go", "regex.go", regox.MATCHED, t)
	checkForMe("reg[eo]x.go", "regOx.go", regox.UNMATCHED, t)
	checkForMe("bh[a-o]rat.go", "bharat.go", regox.MATCHED, t)
	checkForMe("bh[a-e-o]rat.go", "bharat.go", regox.MATCHED, t)
	checkForMe("reg[e-o]x.go", "regOx.go", regox.UNMATCHED, t)
	checkForMe("1[?-]2", "1-2", regox.MATCHED, t)
	checkForMe("1[-?]2", "1*2", regox.UNMATCHED, t)
	checkForMe("1[]?-]2", "1]2", regox.MATCHED, t)
}

func checkForMe(pattern, text, expected string, t *testing.T) {
	actual := regox.Match(0, 0, pattern, text)
	if actual != expected {
		t.Fatalf("After matching [%s] to [%s], got %s but expected %s", pattern, text, actual, expected)
	}
}
