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

/*************************
 * Crawl Context
 *************************/

type CrawlContext struct {
	mutex     sync.Mutex     // protect url slice
	done_urls []string       // store fetched urls(no matter succeeds or not)
	wg        sync.WaitGroup // wait for all crawls
}

func (context *CrawlContext) IsUrlDone(url string) bool {
	context.mutex.Lock()
	defer context.mutex.Unlock()

	for _, done_url := range context.done_urls {
		if done_url == url {
			return true
		}
	}
	return false
}

func (context *CrawlContext) UrlPush(url string) {
	context.mutex.Lock()
	defer context.mutex.Unlock()

	context.done_urls = append(context.done_urls, url)
	return
}

func (context *CrawlContext) JobWait() {
	context.wg.Wait()
}

func (context *CrawlContext) JobAdd() {
	context.wg.Add(1)
}

func (context *CrawlContext) JobDone() {
	context.wg.Done()
}

var crawl_context CrawlContext

// Crawl function
func DoCrawl(url string, depth int, fetcher Fetcher) {

	defer crawl_context.JobDone()

	if depth <= 0 {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	crawl_context.UrlPush(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		if !crawl_context.IsUrlDone(u) {
			crawl_context.JobAdd()
			go DoCrawl(u, depth-1, fetcher)
		}
	}
	return
}

func Crawl(url string, depth int, fetcher Fetcher) {
	crawl_context.JobAdd()
	go DoCrawl(url, depth, fetcher)
	crawl_context.JobWait()
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
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
