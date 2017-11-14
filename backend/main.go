package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"./config"
)

func main() {
	fmt.Printf("BPSearch v1.0.0\n\n")
	start := time.Now()

	// load settings
	// settings := config.Load()

	// load keywords from disk
	// tree := avl.LoadFromDisk()

	// Start crawling and indexing in background
	// ch := make(chan string)
	// go indexer.Start(tree, ch)
	// crawler.Start("https://jeremywho.com", ch)

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
	keyString := r.URL.Query().Get("keywords")
	keywords := strings.Split(keyString, " ")
	keysJSON, err := json.Marshal(keywords)
	if err != nil {
		log.Println(err)
	}
	fmt.Fprintf(w, string(keysJSON))
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
