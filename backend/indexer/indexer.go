package indexer

import (
	"fmt"
	"strings"

	"../interfaces"
)

// Indexer struct
type Indexer struct {
}

// Start indexing the text
func (i *Indexer) Start(str string, store interfaces.StoreInterface) {
	cnt := wordCount(str)
	for word := range cnt {
		// urlMap := make(map[string]int)
		var urlMap map[string]int
		urlMap, _ = store.Get(word)
		urlMap["url2"] = cnt[word]
		store.Put(word, urlMap)
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
