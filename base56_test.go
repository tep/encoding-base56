package base56

import (
	"strconv"
	"testing"
)

func TestEncode(t *testing.T) {
	for _, tc := range []*encodeTestcase{{0, ""}, {1, "1"}, {100, "1n"}, {1540840132, "2nfvKM"}} {
		t.Run(strconv.Itoa(int(tc.in)), tc.test)
	}
}

type encodeTestcase struct {
	in   uint64
	want string
}

func (tc *encodeTestcase) test(t *testing.T) {
	if got := Encode(tc.in); got != tc.want {
		t.Errorf("Encode(%d) := %q; Wanted %q", tc.in, got, tc.want)
	}
}

func TestDecode(t *testing.T) {
	for _, tc := range []*decodeTestcase{{"", 0, nil}, {"1", 1, nil}, {"1n", 100, nil}, {"2nfvKM", 1540840132, nil}, {"NoGood", 0, ErrNotBase56}} {
		t.Run(tc.in, tc.test)
	}
}

type decodeTestcase struct {
	in   string
	want uint64
	err  error
}

func (tc *decodeTestcase) test(t *testing.T) {
	if got, err := Decode(tc.in); got != tc.want || err != tc.err {
		t.Errorf("Decode(%q) := (%d, %v); Wanted (%d, %v)", tc.in, got, err, tc.want, tc.err)
	}
}
