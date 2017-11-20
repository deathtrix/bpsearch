package main

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sort"
	"strings"
	"time"

	_ "net/http/pprof"

	avl "./avltree"
	"./config"
	"./gmaj"
	"./gmaj/gmajpb"
	"github.com/sajari/fuzzy"
)

// Scores struct
type Scores map[string]float64

// Keymap struct
type Keymap map[string]Scores

// P2P/DHT configuration
var nodeConfig = struct {
	id         string
	addr       string
	parentAddr string
	debug      bool // TODO: remove
	pprofAddr  string
}{id: "", addr: ":9988", parentAddr: "", debug: false, pprofAddr: ":9999"}
var node *gmaj.Node

func main() {
	fmt.Printf("BPSearch v1.0.6\n\n")
	start := time.Now()

	// load settings
	// settings := config.Load()

	// load keywords from disk
	// tree := avl.LoadFromDisk()

	node = initNode()

	// keywordStore, _ := tree.Get("soy")
	// keyword, _ := keywordStore.(map[string]interface{})
	// score := keyword["http://www.intermod.ro"].(float64)
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
	go initHTTP()

	// wait for Ctrl-C
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	// shutdown
	log.Println("shutting down")
	node.Shutdown()
	elapsed := time.Since(start)
	fmt.Printf("Time: %s", elapsed)
}

// Initialize the HTTP server
func initHTTP() {
	fmt.Println("Listening on http://localhost:3333/")
	http.HandleFunc("/", handlerRoot)
	http.HandleFunc("/search/", handlerSearch)
	http.HandleFunc("/config-save/", handlerConfigSave)
	http.HandleFunc("/config-load/", handlerConfigLoad)
	http.ListenAndServe(":3333", nil)
}

// Initialize local Chord Node
func initNode() *gmaj.Node {
	// add parent to node
	var parent *gmajpb.Node
	if nodeConfig.parentAddr != "" {
		conn, err := gmaj.Dial(nodeConfig.parentAddr)
		if err != nil {
			log.Fatalf("dialing parent %v failed: %v", nodeConfig.parentAddr, err)
		}

		client := gmajpb.NewGMajClient(conn)
		id, err := client.GetID(context.Background(), &gmajpb.GetIDRequest{})
		_ = conn.Close()
		if err != nil {
			log.Fatalf("getting parent ID failed: %v", err)
		}

		parent = &gmajpb.Node{Id: id.Id, Addr: nodeConfig.parentAddr}
		log.Printf("attaching to %v", gmaj.IDToString(parent.Id))
	}

	var opts []gmaj.NodeOption
	opts = append(opts, gmaj.WithAddress(nodeConfig.addr))

	// add ID to node
	if nodeConfig.id != "" {
		id, err := gmaj.NewID(nodeConfig.id)
		if err != nil {
			log.Fatalf("parsing ID failed: %v", err)
		}
		opts = append(opts, gmaj.WithID(id))
	}

	// create node
	node, err := gmaj.NewNode(parent, opts...)
	if err != nil {
		log.Fatalf("faild to instantiate node: %v", err)
	}

	log.Printf("%+v", node)

	// TODO: remove
	if nodeConfig.debug {
		// gmaj.Put(node, "aaa", []byte("val123"))
		// node.PutKeyVal(context.Background(), &gmajpb.KeyVal{Key: "bbb", Val: []byte("alex")})

		go func() {
			for range time.Tick(5 * time.Second) {
				log.Println(node)
				log.Println(node.DatastoreString())
				val, _ := gmaj.Get(node, "bbb")
				log.Println(string(val))
			}
		}()
	}

	return node
}

// Handle root route
func handlerRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to BPSearch v1.0.6")
}

// Handle search route
func handlerSearch(w http.ResponseWriter, r *http.Request) {
	// load keywords from disk
	tree := avl.LoadFromDisk()

	// Symspell initialization to spellcheck keywords
	model := fuzzy.NewModel()
	model.SetThreshold(1) // For testing only, this is not advisable on production
	model.SetDepth(5)
	words := tree.KeysString()
	model.Train(words)
	model.TrainWord("single")

	keyString := r.URL.Query().Get("keywords")
	keywords := strings.Split(keyString, " ")

	// build list of pages that contain the keywords and add scores
	pages := map[string]float64{}
	hack := map[float64]string{}
	pagesKeys := []float64{}
	for _, keyword := range keywords {
		keyword = model.SpellCheck(keyword)
		// Read from tree
		// keywordStore, _ := tree.Get(keyword)
		// urls, _ := keywordStore.(map[string]interface{})

		// Read from DHT
		keywordByte, _ := gmaj.Get(node, keyword)
		var urls Scores
		b := bytes.NewBuffer(keywordByte)
		d := gob.NewDecoder(b)
		err := d.Decode(&urls)
		if err != nil {
			log.Println(err)
			continue
		}
		for k, v := range urls {
			pages[k] += v
		}
	}

	// sort pages by score
	pagesNames := []string{}
	for k, v := range pages {
		hack[v] = k
		pagesKeys = append(pagesKeys, v)
	}
	sort.Float64s(pagesKeys)
	for _, val := range pagesKeys {
		pagesNames = append(pagesNames, hack[val])
	}

	// encode response as JSON
	b, err := json.Marshal(pagesNames)
	if err != nil {
		log.Println(err)
	}

	fmt.Fprintf(w, string(b))
}

// Handle saving config route
func handlerConfigSave(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	settings := config.Load()
	for k, v := range r.Form {
		settings[k] = strings.Join(v, "")
	}

	config.Save(settings)
}

// Handle loading config route
func handlerConfigLoad(w http.ResponseWriter, r *http.Request) {
	settings := config.LoadJSON()
	fmt.Fprintf(w, settings)
}
