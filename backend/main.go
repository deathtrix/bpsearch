package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	avl "./avltree"
)

type crawl interface {
	Start()
}

type store interface {
	Insert()
	Search()
}

type index interface {
	Get()
	Put()
}

func main() {
	fmt.Printf("BPSearch v1.0.0\n\n")

	// load keywords AVL from disk

	// Start crawling in background
	// go crawler.Start("http://intermod.ro")

	// Run the HTTP server
	http.HandleFunc("/", handlerRoot)
	http.HandleFunc("/search/", handlerSearch)
	http.ListenAndServe(":3333", nil)

	// testAVL()
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

func testAVL() {
	tree := avl.NewWithStringComparator()

	tree.Put("1", "xfdsf")
	tree.Put("2", "fdsfb")
	tree.Put("1", "fdsfa")
	tree.Put("3", "cfdsf")
	tree.Put("4", "dfsf")
	tree.Put("5", "fse")
	tree.Put("6", "sssf")

	fmt.Println(tree)

	val, _ := tree.Get("4")
	fmt.Println(val)

	json, _ := tree.ToJSON()
	fmt.Printf("%+v\n", string(json))

	b := avl.Compress(json)
	avl.Save("out", b.Bytes())

	s := avl.Decompress(b.Bytes())
	fmt.Println(string(s))
}
