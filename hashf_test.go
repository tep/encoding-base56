package base56

import (
	"testing"
)

func TestHashf(t *testing.T) {
	cases := []hfTestcase{
		{Std, "ETcHHsXGBhs"},
		{Alt, "ftCiiSxhdHS"},
		{Py3, "FTcJJsXHDhs"},
	}

	for _, tc := range cases {
		t.Run(tc.name(), tc.test)
	}
}

type hfTestcase struct {
	enc  *Encoding
	want string
}

func (tc *hfTestcase) test(t *testing.T) {
	if got := tc.enc.Hashf("%s:%04X:%08X", "one", 2, 123456789); got != tc.want {
		t.Errorf("Hashf(\"%%s:%%04X:%%08X\", \"one\", 2, 123456789) == %q; wanted %q", got, tc.want)
	}
}

func (tc *hfTestcase) name() string {
	switch tc.enc {
	case Std:
		return "Std"
	case Alt:
		return "Alt"
	case Py3:
		return "Py3"
	default:
		panic("unknown encoding")
	}
}
