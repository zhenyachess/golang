/*
Two lines are received at the input, which contain numbers and "garbage" (extra characters and special signs).
It is necessary to remove the "garbage" from the lines and convert the resulting values into numbers.
It is allowed to use only certain packages: fmt, strconv, unicode, strings, bytes.
*/

package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func findNumberFromString(someshit string) int64 {
	var tempDigit int64
	var tempNumber int64 = 0
	var cntDigits int64 = 1

	for i := len(someshit) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(someshit[i])) {
			tempDigit, _ = strconv.ParseInt(string(someshit[i]), 10, 64)
			tempNumber += tempDigit * cntDigits
			cntDigits *= 10
		}
	}
	return tempNumber
}

func adding(mixed1 string, mixed2 string) int64 {
	return findNumberFromString(mixed1) + findNumberFromString(mixed2)
}

func main() {
	fmt.Println(adding("%8^80", "hhhhh20&&&&nd")) // result: 100
}
