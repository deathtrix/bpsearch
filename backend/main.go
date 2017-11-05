package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	avl "./avltree"
	"./crawler"
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
	tree := loadAVLFromDisk()
	if tree == nil {
		log.Println("Error: Empty tree")
	}

	// Start crawling in background
	crawler.Start("http://jeremywho.com")

	// Run the HTTP server
	// http.HandleFunc("/", handlerRoot)
	// http.HandleFunc("/search/", handlerSearch)
	// http.ListenAndServe(":3333", nil)

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

	// Create AVL tree
	astr := []string{"a", "bb", "ccc"}
	tree.Put("1", astr)
	tree.Put("2", "fdsfb")
	tree.Put("3", "cfdsf")
	tree.Put("4", "dfsf")
	tree.Put("5", "fse")
	tree.Put("6", "sssf")
	fmt.Println(tree)

	// Get AVL by key
	val, _ := tree.Get("1")
	fmt.Println(val)

	// convert to JSON
	// json, _ := tree.ToJSON()
	// fmt.Printf("%+v\n", string(json))

	// save AVL to disk
	// b := avl.Compress(json)
	// avl.Save("out", b.Bytes())

	// load AVL from disk
	// b := avl.Load("out")
	// json := avl.Decompress(b)
	// err := tree.FromJSON(json)
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(tree)
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
