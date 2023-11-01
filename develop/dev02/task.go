package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func UnpackString(str string) string {
	var unpackedStr strings.Builder
	var count int
	for _, ch := range str {
		if unicode.IsDigit(ch) {
			count, _ = strconv.Atoi(string(ch))
		} else {
			if count > 0 {
				unpackedStr.WriteString(strings.Repeat(string(ch), count))
				count = 0
			} else {
				unpackedStr.WriteRune(ch)
			}
		}
	}
	return unpackedStr.String()
}

func main() {
	testCases := []string{"a4bc2d5e", "abcd", "45", ""}
	for _, tc := range testCases {
		fmt.Printf("Unpacking %s: %s\n", tc, UnpackString(tc))
	}
}
