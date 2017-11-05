package crawler

import (
	"fmt"
	"net/http"
	"time"
	"unicode"

	"strings"

	"../indexer"
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
		case <-time.After(3 * time.Second):
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

	var pageText string

	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			pageText = stripSpaces(pageText)
			indexer.Start(pageText)
			return
		}

		if tt == html.TextToken {
			t := z.Token()
			pageText = pageText + t.String()
		}

		if tt == html.StartTagToken {
			t := z.Token()

			if t.Data == "a" {
				for _, a := range t.Attr {
					if a.Key == "href" {

						// if link is within jeremywho.com
						if strings.HasPrefix(a.Val, "http://jeremywho.com") {
							urlProcessor <- a.Val
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

func stripSpaces(str string) string {
	var s string
	r := false
	for _, el := range str {
		if !unicode.IsSpace(el) {
			s = s + string(el)
			r = true
		} else {
			if r {
				s = s + string(el)
			}
			r = false
		}
	}
	return s

	// return strings.Map(func(r rune) rune {
	// 	if unicode.IsSpace(r) {
	// 		return -1
	// 	}
	// 	return r
	// }, str)
}
