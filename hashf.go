package base56

import (
	"fmt"
	"hash/fnv"
)

// Hashf feeds its printf-like arguments into an FNV-1 hash (see "hash/fnv"),
// then returns the hashed 64 bit sum as a base56 encoded string.
func Hashf(format string, args ...interface{}) string {
	h := fnv.New64()
	fmt.Fprintf(h, format, args...)
	return Encode(h.Sum64())
}
