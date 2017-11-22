package base56

import "fmt"

var chars = []byte("0123456789ABCEFGHJKLMNPRSTUVWXYZabcdefghjklmnpqrstuvwxyz")
var srahc = map[rune]uint64{}

func init() {
	for i, c := range chars {
		srahc[rune(c)] = uint64(i)
	}
}

func Decode(s string) (uint64, error) {
	var v uint64
	p := uint64(1)

	for i := len(s) - 1; i >= 0; i-- {
		r := rune(s[i])
		iv, ok := srahc[r]
		if !ok {
			return 0, fmt.Errorf("invalid base56 digit: %c", r)
		}

		v += iv * p
		p *= 56
	}

	return v, nil
}

func Encode(i uint64) string {
	b := make([]byte, 0)
	for i > 0 {
		r := i % 56
		i = i / 56
		b = append([]byte{chars[r]}, b...)
	}

	return string(b)
}
