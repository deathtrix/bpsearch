package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	avl "./avltree"
	"./config"
)

func main() {
	fmt.Printf("BPSearch v1.0.0\n\n")
	start := time.Now()

	// load settings
	// settings := config.Load()

	// load keywords from disk
	// tree := avl.LoadFromDisk()

	// keywordStore, _ := tree.Get("soy")
	// keyword, _ := keywordStore.(map[string]interface{})
	// score := keyword["http://www.intermod.ro"].(float32)
	// fmt.Printf("score: %-v\n", score)

	// keywordStore, _ := tree.Get("soy")
	// urls, _ := keywordStore.(map[string]interface{})
	// fmt.Printf("%.2f\n", urls["http://www.intermod.ro"])
	// for k, v := range urls {
	// 	fmt.Printf("url: %s, score: %.2f\n", k, v)
	// }

	// Start crawling and indexing in background
	// ch := make(chan string)
	// go indexer.Start(tree, ch)
	// crawler.Start("http://www.intermod.ro", ch)

	// time.Sleep(10000 * time.Millisecond)

	// Run the HTTP server
	fmt.Println("Listening on http://localhost:3333/")
	http.HandleFunc("/", handlerRoot)
	http.HandleFunc("/search/", handlerSearch)
	http.HandleFunc("/config-save/", handlerConfigSave)
	http.HandleFunc("/config-load/", handlerConfigLoad)
	http.ListenAndServe(":3333", nil)

	elapsed := time.Since(start)
	fmt.Printf("Time: %s", elapsed)
}

func handlerRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to BPSearch v1.0.0")
}

func handlerSearch(w http.ResponseWriter, r *http.Request) {
	// load keywords from disk
	tree := avl.LoadFromDisk()

	keyString := r.URL.Query().Get("keywords")
	keywords := strings.Split(keyString, " ")

	s := ""
	pages := map[string]float64{}
	for _, keyword := range keywords {
		keywordStore, _ := tree.Get(keyword)
		urls, _ := keywordStore.(map[string]interface{})
		for k, v := range urls {
			score, _ := v.(float64)
			pages[k] += score
		}
	}

	// TODO: sort pages by score
	// sort.Sort(ByLength(elems))

	for k := range pages {
		s += fmt.Sprintf("<a href=\"%s\">%s</a><br>", k, k)
	}

	fmt.Fprintf(w, s)
}

func handlerConfigSave(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	settings := config.Load()
	for k, v := range r.Form {
		settings[k] = strings.Join(v, "")
	}

	config.Save(settings)
}

func handlerConfigLoad(w http.ResponseWriter, r *http.Request) {
	settings := config.LoadJSON()
	fmt.Fprintf(w, settings)
}
