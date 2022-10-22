package base56

import (
	"fmt"
	"testing"
)

func TestEncodeDecode(t *testing.T) {
	var cases []ecTestcase
	cases = append(cases, mkTestcases(0, "", "", "")...)
	cases = append(cases, mkTestcases(1, "1", "3", "3")...)
	cases = append(cases, mkTestcases(100, "1n", "3N", "3n")...)
	cases = append(cases, mkTestcases(1540840132, "2nfvKM", "4NFVkn", "4nfvLN")...)

	for _, tc := range cases {
		t.Run("Enc:"+tc.name(), tc.testEncode)
		t.Run("Dec:"+tc.name(), tc.testDecode)
	}
}

func (tc *ecTestcase) testEncode(t *testing.T) {
	if got := tc.enc.Encode(tc.num); got != tc.str {
		t.Errorf("Encode(%d) := %q; Wanted %q", tc.num, got, tc.str)
	}
}

func (tc *ecTestcase) testDecode(t *testing.T) {
	if got, err := tc.enc.Decode(tc.str); err != nil || got != tc.num {
		t.Errorf("Encode(%q) := (%d, %v); Wanted (%d, nil)", tc.str, got, err, tc.num)
	}
}

type ecTestcase struct {
	enc *Encoding
	num uint64
	str string
}

func (tc *ecTestcase) name() string {
	var n string
	switch tc.enc {
	case Std:
		n = "Std"
	case Alt:
		n = "Alt"
	case Py3:
		n = "Py3"
	default:
		n = "<unknown>"
	}

	return fmt.Sprintf("%s:%d", n, tc.num)
}

func mkTestcases(i uint64, std, alt, py3 string) []ecTestcase {
	return []ecTestcase{
		{Std, i, std},
		{Alt, i, alt},
		{Py3, i, py3},
	}
}
