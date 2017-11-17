package main

import (
	"fmt"

	"../../internal/chord"
	"github.com/r-medina/gmaj"
	"github.com/r-medina/gmaj/gmajpb"

	"golang.org/x/net/context"
)

var client chord.ChordClient

func main() {
	conn, _ := gmaj.Dial("localhost:25450")
	client = chord.NewChordClient(conn)

	key := "aaa"
	val := []byte("val123")

	client.PutKeyVal(context.Background(), &gmajpb.KeyVal{Key: key, Val: val})

	resp, _ := client.GetKey(context.Background(), &gmajpb.Key{Key: key})
	fmt.Printf("%s", resp.Val)
}
