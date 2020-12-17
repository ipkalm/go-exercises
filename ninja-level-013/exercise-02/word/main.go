// Package word provide custom funcs for
// working with string and words.
package word

import "strings"

// UseCount returns map, which contains word as key
// and value as number of repeating this word in a string.
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count returns the number of words from a input string.
func Count(s string) int {
	return len(strings.Fields(s))
}
