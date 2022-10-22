//
// Copyright 2022 Timothy E. Peoples
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
// For compatibility with other implementions, this package supports three
// separate and distinct base56 character sets through the package level,
// Encoding variables Std, Alt and Py3. Each of these Encodings leverage only
// 7-bit clean (ASCII) characters.
//
// The Std Encoding employs the original character set used by this package and
// is composed of the numerals 0-9 followed by all upper case characters except
// for 'D' (0x44), 'I' (0x49), 'O' (0x4f), and 'Q' (0x51) and then all lower
// case characters except for 'i' (0x69) and 'o' (0x6f).
//
// The Alt Encoding is compatible with PHP and Java implementations and is
// defined as the numerals 2-9 followed by all lower case characters except for
// 'l' (0x6c) and 'o' (0x6f) and then all upper case characters except for 'I'
// (0x49) and 'O' (0x4f).
//
// The Py3 Encoding is compatible with the Python-3 implementation and is
// defined as the numerals 2-9 followed by all upper case characters except for
// 'I' (0x49) and 'O' (0x4f) and then all lower case characters except for 'l'
// (0x6c) and 'o' (0x6f).
//
// Note, the Alt and Py3 Encodings are identical except for the order of
// character classes.  Alt is numerals->lowercase->uppercase while Py3 is
// numerals->uppercase->lowercase.
//
// For reference, here are links to the known implementations for other
// languages:
//
//     PHP....: http://rossduggan.ie/blog/codetry/base-56-integer-encoding-in-php/index.html
//     Java...: ??
//     Python.: https://github.com/jyn514/base56
//
package base56 // import "toolman.org/encoding/base56"

import "errors"

var (
	// Std is the standard character set traditionally used by this package.
	Std = charSet("0123456789ABCEFGHJKLMNPRSTUVWXYZabcdefghjklmnpqrstuvwxyz")

	// Alt is an alternative character set used by PHP and Java implementations.
	Alt = charSet("23456789abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ")

	// Py3 is the character set used by the Python base56 library.
	Py3 = charSet("23456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnpqrstuvwxyz")
)

// ErrNotBase56 is returned by Decode if it is provided an invalid base56 value
// for the associated Encoding.
var ErrNotBase56 = errors.New("invalid base56 value")

// Decode is a convenience wrapper around Std.Decode.
func Decode(s string) (uint64, error) {
	return Std.Decode(s)
}

// Encode is a convenience wrapper around Std.Encode.
func Encode(i uint64) string {
	return Std.Encode(i)
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// Encoding represents a distinct base56 character set.
type Encoding struct {
	runes []rune
	vmap  map[rune]uint64
}

// charSet is the unexported constructor for type Encoding. The provided
// string should be composed of 56 unique characters in LSB to MSB order.
// See the definitions of Std, Alt, and Py3 above.
func charSet(s string) *Encoding {
	e := &Encoding{runes: []rune(s), vmap: make(map[rune]uint64)}
	for i, r := range e.runes {
		e.vmap[r] = uint64(i)
	}
	return e
}

// Decode accepts a valid, base56 string -- as returned by Encode -- and
// returns its uint64 value, or zero and ErrNotBase56 if the base56 string
// is invalid.  Valid base56 values are composed of characters from the
// receiver's defined character set.
func (e *Encoding) Decode(s string) (uint64, error) {
	var v uint64
	p := uint64(1)

	for i := len(s) - 1; i >= 0; i-- {
		r := rune(s[i])
		iv, ok := e.vmap[r]
		if !ok {
			return 0, ErrNotBase56
		}

		v += iv * p
		p *= 56
	}

	return v, nil
}

// Encode accepts a uint64 value and encodes it to a base56 string composed
// of characters from the receiver's defined character set.
func (e *Encoding) Encode(i uint64) string {
	var b []rune

	for i > 0 {
		r := i % 56
		i = i / 56
		b = append([]rune{e.runes[r]}, b...)
	}

	return string(b)
}
