// Package morestrings implements additional functions to manipulate UTF-8
// encoded strings, beyond what is provided in the standard "strings" package.
package morestrings

import "fmt"

// ReverseRunes returns its argument string reversed rune-wise left to right.
func ReverseRunes(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
