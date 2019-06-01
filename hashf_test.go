package base56

import (
	"testing"
)

func TestHashf(t *testing.T) {
	want := "ETcHHsXGBhs"
	if got := Hashf("%s:%04X:%08X", "one", 2, 123456789); got != want {
		t.Errorf("Hashf(\"%%s:%%04X:%%08X\", \"one\", 2, 123456789) == %q; wanted %q", got, want)
	}
}
