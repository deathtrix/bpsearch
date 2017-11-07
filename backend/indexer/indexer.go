package indexer

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"unicode"

	"../interfaces"
)

// Result struct
type Result struct {
	url    string
	weight int
}

// Indexer struct
type Indexer []Result

// Start indexing the text
func Start(store interfaces.StoreInterface, ch <-chan string) {
	for {
		select {
		case urlStr := <-ch:
			index(urlStr, store)
			// case <-time.After(3 * time.Second):
			// 	break
		}
	}
}

func index(urlStr string, store interfaces.StoreInterface) {
	resp, err := http.Get(urlStr)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	// TODO: extract text using headless browser
	str := string(respBytes)

	cnt := wordCount(str)
	for word := range cnt {
		old, _ := store.Get(word)

		var urlMap = make(map[string]interface{})
		if old != nil {
			urlMap = old.(map[string]interface{})
		}
		urlMap["url2"] = cnt[word]

		// use Indexer struct - problems with AVL serialization
		// urlMap := Indexer{}
		// if t, ok := old.(Indexer); ok {
		// 	urlMap = t
		// }
		// urlMap = Indexer{Result{url: "url3", weight: cnt[word]}}

		store.Put(word, urlMap)
		fmt.Printf("%s - %d\n", word, cnt[word])
	}

	store.SaveToDisk()
}

func wordCount(s string) map[string]int {
	words := strings.Fields(s)
	counts := make(map[string]int, len(words))
	for _, word := range words {
		counts[strings.ToLower(word)]++
	}
	return counts
}

func stripSpaces(str string) string {
	var s string
	r := false
	for _, el := range str {
		if !unicode.IsSpace(el) {
			s = s + string(el)
			r = true
		} else {
			if r {
				s = s + string(el)
			}
			r = false
		}
	}
	return s

	// return strings.Map(func(r rune) rune {
	// 	if unicode.IsSpace(r) {
	// 		return -1
	// 	}
	// 	return r
	// }, str)
}
