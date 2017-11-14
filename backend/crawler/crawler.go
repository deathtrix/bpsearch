package crawler

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"strings"

	"golang.org/x/net/html"
)

// Start crawling
func Start(urlStr string, ch chan<- string) {
	_, err := url.ParseRequestURI(urlStr)
	if err != nil {
		log.Println(err)
		return
	}
	urlProcessor := make(chan string)
	done := make(chan bool)

	go processURL(urlProcessor, done, ch)
	urlProcessor <- urlStr

	<-done
	fmt.Println("Done")
}

func processURL(urlProcessor chan string, done chan bool, ch chan<- string) {
	visited := make(map[string]bool)
	for {
		select {
		case url := <-urlProcessor:
			if _, ok := visited[url]; ok {
				continue
			} else {
				visited[url] = true
				go exploreURL(url, urlProcessor, ch)
			}
		case <-time.After(3 * time.Second):
			fmt.Printf("\nExplored %d pages\n", len(visited))
			done <- true
		}
	}
}

func exploreURL(urlStr string, urlProcessor chan string, ch chan<- string) {
	fmt.Printf("Visiting %s.\n", urlStr)

	resp, err := http.Get(urlStr)
	if err != nil {
		log.Println(err)
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
						_, err := url.ParseRequestURI(a.Val)
						if err != nil {
							continue
						}

						// if link is within jeremywho.com
						if strings.HasPrefix(a.Val, "http://www.intermod.ro") {
							urlProcessor <- a.Val
							ch <- a.Val
						}

						// crawl every link in page (external links also)
						// if strings.HasPrefix(a.Val, "http") {
						// 	urlProcessor <- a.Val
						// } else {
						// 	urlProcessor <- "http://intermod.ro" + a.Val // TODO: optimize concatenation
						// }
					}
				}
			}
		}
	}
}
