package crawler

import (
	"fmt"
	"net/http"
	"time"

	"strings"

	"golang.org/x/net/html"
)

// Start crawling
func Start(url string) {
	urlProcessor := make(chan string)
	done := make(chan bool)

	go processURL(urlProcessor, done)
	urlProcessor <- url

	<-done
	fmt.Println("Done")
}

func processURL(urlProcessor chan string, done chan bool) {
	visited := make(map[string]bool)
	for {
		select {
		case url := <-urlProcessor:
			if _, ok := visited[url]; ok {
				continue
			} else {
				visited[url] = true
				go exploreURL(url, urlProcessor)
			}
		case <-time.After(1 * time.Second):
			fmt.Printf("Explored %d pages\n", len(visited))
			done <- true
		}
	}
}

func exploreURL(url string, urlProcessor chan string) {
	fmt.Printf("Visiting %s.\n", url)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	z := html.NewTokenizer(resp.Body)

	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			return
		}

		if tt == html.StartTagToken {
			t := z.Token()

			if t.Data == "a" {
				for _, a := range t.Attr {
					if a.Key == "href" {

						// if link is within jeremywho.com
						// if strings.HasPrefix(a.Val, "https://jeremywho.com") {
						if strings.HasPrefix(a.Val, "http") {
							urlProcessor <- a.Val
						} else {
							urlProcessor <- "http://intermod.ro" + a.Val // TODO: optimize concatenation
						}
						// }
					}
				}
			}
		}
	}
}
