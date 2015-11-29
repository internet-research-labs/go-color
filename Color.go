package main

import (
	"bytes"
	"math/rand"
)

// Return a reversed string
func Reverse(s string) string {
	var result bytes.Buffer
	n := len(s)

	for i := range s {
		result.WriteRune(rune(s[n-i-1]))
	}

	return result.String()
}

// Return a random integer from min to max
func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

// Return a hex string of an integer
func intToHex(val uint8) string {
	var result bytes.Buffer

	if val == 0 {
		return "0"
	}

	for val > 0 {
		var char rune
		i := val % 16
		val /= 16

		switch {
		case 0 <= i && i <= 9:
			char = rune('0' + i)
		case 10 <= i && i <= 15:
			char = rune('A' + i - 10)
		default:
			char = '?'
		}

		result.WriteRune(char)
	}

	return Reverse(result.String())
}

// Return a string of length
func pad(val string, length int) string {
	var result bytes.Buffer
	n := len(val)
	for i := 0; i < length-n; i++ {
		result.WriteRune('0')
	}
	return result.String() + val
}
