package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "net/http/pprof"

	"github.com/r-medina/gmaj"
	"github.com/r-medina/gmaj/gmajpb"
)

var config = struct {
	id         string
	addr       string
	parentAddr string
	debug      bool
	pprofAddr  string
}{id: "", addr: "127.0.0.1:1234", parentAddr: "", debug: true, pprofAddr: ":9999"}

func main() {
	// run pprof
	log.Printf("running pprof server on %s", config.pprofAddr)
	go func() {
		log.Println(http.ListenAndServe(config.pprofAddr, nil))
	}()

	// add parent to node
	var parent *gmajpb.Node
	if config.parentAddr != "" {
		conn, err := gmaj.Dial(config.parentAddr)
		if err != nil {
			log.Fatalf("dialing parent %v failed: %v", config.parentAddr, err)
		}

		client := gmajpb.NewGMajClient(conn)
		id, err := client.GetID(context.Background(), &gmajpb.GetIDRequest{})
		_ = conn.Close()
		if err != nil {
			log.Fatalf("getting parent ID failed: %v", err)
		}

		parent = &gmajpb.Node{Id: id.Id, Addr: config.parentAddr}
		log.Printf("attaching to %v", gmaj.IDToString(parent.Id))
	}

	var opts []gmaj.NodeOption
	opts = append(opts, gmaj.WithAddress(config.addr))

	// add ID to node
	if config.id != "" {
		id, err := gmaj.NewID(config.id)
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

	if config.debug {
		gmaj.Put(node, "aaa", []byte("val123"))
		node.PutKeyVal(context.Background(), &gmajpb.KeyVal{Key: "bbb", Val: []byte("alex")})

		go func() {
			for range time.Tick(5 * time.Second) {
				log.Println(node)
				log.Println(node.DatastoreString())
				val, _ := gmaj.Get(node, "bbb")
				log.Println(string(val))
			}
		}()
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	sig := <-stop
	log.Printf("received signal %v", sig)

	log.Println("shutting down")
	node.Shutdown()
}
