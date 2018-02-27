package main

import (
	"fmt"
)

// determine if a given string has only unique characters
// do not use additional data structure
func uniqueChars2(s string) bool {
	for i := range s {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return false
			}
		}
	}
	return true
}

// determine if a given string has only unique characters
// use additional data structure
func uniqueChars(s string) bool {
	m := map[rune]int{}
	for _, v := range s {
		// discount space character
		if v == 32 {
			continue
		}
		count, ok := m[v]
		if !ok {
			m[v] = 1
			continue
		}
		m[v] = count + 1
	}
	for i := range m {
		if m[i] > 1 {
			return false
		}
	}
	return true
}

func main() {
	s := "mahctur"
	fmt.Println(uniqueChars(s))
	fmt.Println(uniqueChars2(s))
}
