

# base56
`import "toolman.org/encoding/base56"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## Install

```sh
  go get toolman.org/encoding/base56
```

## <a name="pkg-overview">Overview</a>
Package base56 provides functions for encoding/decoding uint64 values as
short, easily digestible, base56 strings.

The set of 56 digits are the numerals 0 - 9 plus all upper and lower case
(ASCII) characters except for [DIOQio] (each of which may be easily confused
with the numerals 0 or 1).




## <a name="pkg-index">Index</a>
* [func Decode(s string) (uint64, error)](#Decode)
* [func Encode(i uint64) string](#Encode)


#### <a name="pkg-files">Package files</a>
[base56.go](/src/toolman.org/encoding/base56/base56.go) 





## <a name="Decode">func</a> [Decode](/src/target/base56.go?s=2013:2050#L37)
``` go
func Decode(s string) (uint64, error)
```
Decode takes a valid, base56 string -- as returned by Encode -- and returns
its uint64 value, or zero and an error if the base56 string is invalid.
Valid base56 values are composed of digits in the range
[0-9ABCE-HJ-NPR-Za-hj-np-z] (which is all ASCII numerals, upper and lower
case letters except for [DIOQio])



## <a name="Encode">func</a> [Encode](/src/target/base56.go?s=2389:2417#L57)
``` go
func Encode(i uint64) string
```
Encode accepts a uint64 value and encodes it to a base56 string composed of
digits as described by Decode.

