package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	wordMap := make(map[string]int)

	for _, word := range strings.Fields(s) {
		value, exists := wordMap[word]

		if exists {
			wordMap[word] = value + 1
		} else {
			wordMap[word] = 1
		}
	}
	return wordMap
}

func main() {
	wc.Test(WordCount)
}
