//
// Copyright 2018 Timothy E. Peoples
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
// IN THE SOFTWARE.
//

// Package base56 provides functions for encoding/decoding uint64 values as
// short, easily digestible, base56 strings.
//
// The set of 56 digits are the numerals 0 - 9 plus all upper and lower case
// (ASCII) characters except for [DIOQio] (each of which may be easily confused
// with the numerals 0 or 1).
package base56 // import "toolman.org/encoding/base56"

import "errors"

var chars = []byte("0123456789ABCEFGHJKLMNPRSTUVWXYZabcdefghjklmnpqrstuvwxyz")
var srahc = map[rune]uint64{}

func init() {
	for i, c := range chars {
		srahc[rune(c)] = uint64(i)
	}
}

// ErrNotBase56 is returned by Decode if it is given an invalid base56 value.
var ErrNotBase56 = errors.New("invalid base56 value")

// Decode takes a valid, base56 string -- as returned by Encode -- and returns
// its uint64 value, or zero and ErrNotBase56 if the base56 string is invalid.
// Valid base56 values are composed of digits in the range
// [0-9ABCE-HJ-NPR-Za-hj-np-z] -- which is the numerals 0 through 9 plus all
// upper and lower case ASCII letters except for [DIOQio].
func Decode(s string) (uint64, error) {
	var v uint64
	p := uint64(1)

	for i := len(s) - 1; i >= 0; i-- {
		r := rune(s[i])
		iv, ok := srahc[r]
		if !ok {
			return 0, ErrNotBase56
		}

		v += iv * p
		p *= 56
	}

	return v, nil
}

// Encode accepts a uint64 value and encodes it to a base56 string composed of
// digits as described by Decode.
func Encode(i uint64) string {
	b := make([]byte, 0)
	for i > 0 {
		r := i % 56
		i = i / 56
		b = append([]byte{chars[r]}, b...)
	}

	return string(b)
}
