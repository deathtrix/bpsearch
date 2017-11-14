package indexer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"unicode"

	"../config"
	"../interfaces"
	"github.com/k4s/phantomgo"
)

// type Score struct {
// 	url   string
// 	score float64
// }
// type Scores []Score

// Scores struct
type Scores map[string]float64

// Keymap struct
type Keymap map[string]Scores

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

func getPageContent(urlStr string) map[string]int {
	resp, err := http.Get(urlStr)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	str := string(respBytes)
	cnt := wordCount(str)
	return cnt
}

func index(urlStr string, store interfaces.StoreInterface) {
	var scores map[string]float64
	scoresJSON := parseHTML(urlStr)
	err := json.Unmarshal(scoresJSON, &scores)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%-v\n", scores)

	for word, score := range scores {
		old, _ := store.Get(word)

		// var urlMap = make(map[string]interface{})
		// if old != nil {
		// 	urlMap = old.(map[string]interface{})
		// }
		// urlMap[urlStr] = score

		// use structs - problems with AVL serialization
		urlMap := Scores{}
		if t, ok := old.(Scores); ok {
			urlMap = t
		}
		urlMap[urlStr] = score

		store.Put(word, urlMap)
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

func parseHTML(urlStr string) []byte {
	// load settings
	settings := config.Load()

	p := phantomgo.NewPhantom()
	jsBytes, err := ioutil.ReadFile("parse.js")
	if err != nil {
		fmt.Println(err)
	}
	js := string(jsBytes)
	js = strings.Replace(js, "<<URL>>", urlStr, -1)
	js = strings.Replace(js, "<<SIZE_WEIGHT>>", settings["SIZE_WEIGHT"], -1)
	js = strings.Replace(js, "<<BOLD_WEIGHT>>", settings["BOLD_WEIGHT"], -1)
	js = strings.Replace(js, "<<H1_WEIGHT>>", settings["H1_WEIGHT"], -1)
	js = strings.Replace(js, "<<H2_WEIGHT>>", settings["H2_WEIGHT"], -1)
	js = strings.Replace(js, "<<H3_WEIGHT>>", settings["H3_WEIGHT"], -1)
	js = strings.Replace(js, "<<H4_WEIGHT>>", settings["H4_WEIGHT"], -1)
	js = strings.Replace(js, "<<NRP_WEIGHT>>", settings["NRP_WEIGHT"], -1)

	res, _ := p.Exec(js)
	output, _ := ioutil.ReadAll(res)
	return output
}
