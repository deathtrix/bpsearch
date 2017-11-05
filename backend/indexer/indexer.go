package indexer

import (
	"fmt"
	"strings"
)

// Start indexing the text
func Start(str string) {
	cnt := wordCount(str)
	for word := range cnt {
		fmt.Printf("%s - %d\n", word, cnt[word])
	}
}

func wordCount(s string) map[string]int {
	words := strings.Fields(s)
	counts := make(map[string]int, len(words))
	for _, word := range words {
		counts[strings.ToLower(word)]++
	}
	return counts
}
