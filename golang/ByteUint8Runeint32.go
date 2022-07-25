package main

import (
	"bytes"
	"unicode"
)

func CharMain() {
	s := "1234abcdABCD"
	// s[] is byte == uint
	// for i,e := range s,   e is rune = int32

	bytes_var := []byte(s)

	bytes_var = bytes.ToLower(bytes_var)
	rune_char := rune(bytes_var[0])
	_ = byte(rune_char)

	print(unicode.IsLower(rune_char) || unicode.IsNumber(rune_char))

}
