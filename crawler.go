package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, cache *UrlCache, wg *sync.WaitGroup) {
	defer wg.Done()
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	cache.Set(url, body)
	for _, u := range urls {
		if _, ok := cache.Get(u); !ok {
			wg.Add(1)
			go Crawl(u, depth-1, fetcher, cache, wg)
		}
	}
	return
}

type UrlCache struct {
	urls map[string]string
	mux sync.Mutex
}

func (c *UrlCache) Set(key, value string) {
	c.mux.Lock()
	c.urls[key] = value
	c.mux.Unlock()
}

func (c *UrlCache) Get(key string) (string, bool) {
	defer c.mux.Unlock()
	c.mux.Lock()
	v, ok := c.urls[key]
	return v, ok
}

func main() {
	cache := UrlCache{urls: make(map[string]string)}
	var wg sync.WaitGroup
	wg.Add(1)
	go Crawl("http://golang.org/", 4, fetcher, &cache, &wg)
	wg.Wait()
	for url, body := range cache.urls {
		fmt.Printf("found: %s %q\n", url, body)
	}
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
