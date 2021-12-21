package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"

	"golang.org/x/net/html"
)

type result struct {
	url   string
	urls  []string
	err   error
	depth int
}

var urlFetched1 map[string]bool

func CrawlConcurrent(url string, depth int) {

	ch := make(chan *result)

	fetch := func(url string, depth int) {
		urls, errs := findLinks1(url)
		ch <- &result{
			url:   url,
			urls:  urls,
			err:   errs,
			depth: depth,
		}
	}

	go fetch(url, depth)
	urlFetched1[url] = true

	for fetching := 1; fetching > 0; fetching-- {
		res := <-ch

		if res.err != nil {
			continue
		}

		fmt.Println("Found url :- ", res.url)
		if res.depth > 0 {
			for _, u := range res.urls {
				if !urlFetched1[u] {
					fetching++
					go fetch(u, res.depth-1)
					urlFetched1[u] = true
				}
			}
		}

	}

	close(ch)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	urlFetched1 = make(map[string]bool)
	now := time.Now()
	CrawlConcurrent("https://google.com", 1)

	fmt.Println("Time taken :- ", time.Since(now))
}

func findLinks1(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s : %s ", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	return visit1(nil, doc), nil
}

func visit1(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit1(links, c)
	}

	return links
}
