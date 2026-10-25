// Lời giải: Web Crawler
package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

// visited nhớ các URL đã (hoặc đang) được lấy, có mutex bảo vệ.
//
// Điểm mấu chốt: kiểm tra "đã thấy chưa" và đánh dấu "đã thấy" phải nằm
// TRONG CÙNG một lần giữ khóa. Nếu tách ra hai bước, hai goroutine có thể
// cùng vượt qua bước kiểm tra rồi cùng lấy một URL.
type visited struct {
	mu sync.Mutex
	m  map[string]bool
}

// markSeen trả về true nếu url là LẦN ĐẦU được nhìn thấy.
func (v *visited) markSeen(url string) bool {
	v.mu.Lock()
	defer v.mu.Unlock()

	if v.m[url] {
		return false
	}
	v.m[url] = true
	return true
}

func Crawl(url string, depth int, fetcher Fetcher) {
	v := &visited{m: make(map[string]bool)}
	var wg sync.WaitGroup

	// Hàm đệ quy được khai báo bằng var trước rồi mới gán,
	// vì closure cần tham chiếu tới chính nó.
	var crawl func(url string, depth int)
	crawl = func(url string, depth int) {
		defer wg.Done()

		if depth <= 0 || !v.markSeen(url) {
			return
		}

		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("found: %s %q\n", url, body)

		for _, u := range urls {
			// wg.Add PHẢI gọi ở goroutine cha, TRƯỚC khi go chạy.
			// Gọi bên trong goroutine con thì wg.Wait có thể chạy trước
			// và thoát sớm.
			wg.Add(1)
			go crawl(u, depth-1)
		}
	}

	wg.Add(1)
	go crawl(url, depth)
	wg.Wait()
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

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
