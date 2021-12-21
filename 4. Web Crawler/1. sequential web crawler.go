package main

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/html"
)

var urlFetched map[string]bool

func Crawl(url string, depth int) {
	if depth < 0 {
		return
	}

	urls, errs := findLinks(url)
	if errs != nil {
		return
	}

	fmt.Println("Found url :- ", url)
	urlFetched[url] = true

	for _, u := range urls {
		if !urlFetched[u] {
			Crawl(u, depth-1)
		}
	}
	return
}

func main() {
	urlFetched = make(map[string]bool)
	now := time.Now()
	Crawl("https://google.com", 1)
	fmt.Println("Time taken :- ", time.Since(now))
}

func findLinks(url string) ([]string, error) {
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

	return visit(nil, doc), nil
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}

	return links
}
