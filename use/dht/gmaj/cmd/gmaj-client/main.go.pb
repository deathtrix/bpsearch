package main

import (
	"fmt"

	"github.com/r-medina/gmaj"
	"github.com/r-medina/gmaj/gmajpb"

	"golang.org/x/net/context"
)

var client gmajpb.GMajClient

func main() {
	conn, _ := gmaj.Dial("localhost:1234")
	client = gmajpb.NewGMajClient(conn)

	key := "ddd"
	val := []byte("value")

	client.Put(context.Background(), &gmajpb.PutRequest{Key: key, Value: val})

	resp, _ := client.Get(context.Background(), &gmajpb.GetRequest{Key: key})
	fmt.Printf("%s", resp.Value)
}
