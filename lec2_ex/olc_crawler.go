package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string, reply chan fetchReply)
}

type fetchReply struct {
	sendUrl string
	body    string
	urls    []string
	err     error
}

type urlMap struct {
	m     map[string]int
	mutex sync.Mutex
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	var wg sync.WaitGroup
	urlDict := urlMap{}
	urlDict.m = make(map[string]int)
	reply := make(chan fetchReply)
	urlDict.m[url] = int(1)
	//	i := 0
	wg.Add(1)
	go fetcher.Fetch(url, reply)

	go func() {
		wg.Wait()
		close(reply)
	}()

	for result := range reply {
		if result.err != nil {
			fmt.Println(result.err)
			wg.Done()
		} else {
			fmt.Printf("found: %s %q\n", result.sendUrl, result.body)
			for _, u := range result.urls {
				//				fmt.Printf("%d: %v\n", i, urlDict.m)
				//				i++
				if _, ok := urlDict.m[u]; !ok {
					urlDict.m[u] = int(1)
					wg.Add(1)
					go fetcher.Fetch(u, reply)
				}
			}
			wg.Done()
		}

	}
}

func main() {
	Crawl("https://golang.org/", fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string, reply chan fetchReply) {
	// if initialization; condition {}
	if res, ok := f[url]; ok {
		reply <- fetchReply{url, res.body, res.urls, nil}
	} else {
		reply <- fetchReply{url, "", nil, fmt.Errorf("not found: %s", url)}
	}
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
