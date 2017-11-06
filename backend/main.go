package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	avl "./avltree"
	"./crawler"
	"./indexer"
	"./interfaces"
)

func main() {
	fmt.Printf("BPSearch v1.0.0\n\n")
	start := time.Now()

	// load keywords from disk
	tree := loadAVLFromDisk()

	val, _ := tree.Get("one")
	fmt.Printf("%-v\n", val)

	// Start crawling and indexing in background
	ch := make(chan string)
	crawler := new(crawler.Crawler)
	indexer := new(indexer.Indexer)
	go startIndexer(indexer, tree, ch)
	startCrawler("http://jeremywho.com", crawler, ch)

	// Run the HTTP server
	// fmt.Println("Listening on http://localhost:3333/")
	// http.HandleFunc("/", handlerRoot)
	// http.HandleFunc("/search/", handlerSearch)
	// http.ListenAndServe(":3333", nil)

	// Save keywords to disk
	saveAVLToDisk(tree)

	elapsed := time.Since(start)
	fmt.Printf("Time: %s", elapsed)
}

func startIndexer(indexer interfaces.IndexerInterface, store interfaces.StoreInterface, ch <-chan string) {
	for {
		select {
		case pageText := <-ch:
			indexer.Start(pageText, store)
			// case <-time.After(3 * time.Second):
			// 	break
		}
	}
}

func startCrawler(urlStr string, crawler interfaces.CrawlerInterface, ch chan<- string) {
	crawler.Start(urlStr, ch)
}

func handlerRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to BPSearch v1.0.0")
}

func handlerSearch(w http.ResponseWriter, r *http.Request) {
	keyString := r.URL.Query().Get("keywords")
	keywords := strings.Split(keyString, " ")
	keysJSON, err := json.Marshal(keywords)
	if err != nil {
		log.Println(err)
	}
	fmt.Fprintf(w, string(keysJSON))
}

func loadAVLFromDisk() *avl.Tree {
	tree := avl.NewWithStringComparator()
	b := avl.Load("out")
	json := avl.Decompress(b)
	err := tree.FromJSON(json)
	if err != nil {
		log.Println(err)
	}

	return tree
}

func saveAVLToDisk(tree *avl.Tree) {
	json, _ := tree.ToJSON()
	b := avl.Compress(json)
	avl.Save("out", b.Bytes())
}
