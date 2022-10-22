

# base56 [![MIT License][mit-img]][mit] [![GitHub Release][release-img]][release] [![GoDoc][godoc-img]][godoc] [![Go Report Card][reportcard-img]][reportcard]

`import "toolman.org/encoding/base56"`

## Install

```sh
  go get toolman.org/encoding/base56
```
## Overview

Package base56 provides functions for encoding/decoding uint64 values as
short, easily digestible, base56 strings. For example, the `Std` encoding
transforms the `uint64` value `1540840132` into the string `"2nfvKM"`.

For compatibility with other implementions, this package supports three
separate and distinct base56 character sets through the package level,
Encoding variables `Std`, `Alt` and `Py3`. Each of these Encodings leverage
only 7-bit clean (ASCII) characters.

The `Std` Encoding employs the original character set used by this package and
is composed of the numerals 0-9 followed by all upper case characters except
for 'D' (`0x44`), 'I' (`0x49`), 'O' (`0x4f`), and 'Q' (`0x51`) and then all
lower case characters except for 'i' (`0x69`) and 'o' (`0x6f`).

The `Alt` Encoding is compatible with PHP and Java implementations and is
defined as the numerals 2-9 followed by all lower case characters except for
'l' (`0x6c`) and 'o' (`0x6f`) and then all upper case characters except for 'I'
(`0x49`) and 'O' (`0x4f`).

The `Py3` Encoding is compatible with the Python-3 implementation and is
defined as the numerals 2-9 followed by all upper case characters except for
'I' (`0x49`) and 'O' (`0x4f`) and then all lower case characters except for 'l'
(`0x6c`) and 'o' (`0x6f`).

Note, the `Alt` and `Py3` Encodings are identical except for the order of
character classes.  `Alt` is numerals->lowercase->uppercase while `Py3` is
numerals->uppercase->lowercase.



<!--- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - --->

[mit]:            https://github.com/tep/encoding-base56/blob/master/LICENSE
[mit-img]:        http://img.shields.io/badge/License-MIT-c41e3a.svg

[release]:        https://github.com/tep/encoding-base56/releases
[release-img]:    https://img.shields.io/github/release/tep/encoding-base56/all.svg

[godoc]:          https://godoc.org/toolman.org/encoding/base56
[godoc-img]:      https://godoc.org/toolman.org/encoding/base56?status.svg

[reportcard]:     https://goreportcard.com/report/toolman.org/encoding/base56
[reportcard-img]: https://goreportcard.com/badge/toolman.org/encoding/base56
