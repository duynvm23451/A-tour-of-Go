// Exercise: Web Crawler / Bài tập: Trình thu thập web
//
// ĐỀ BÀI
// Sửa hàm Crawl để nó lấy các URL SONG SONG mà KHÔNG lấy trùng một URL hai lần.
//
// Gợi ý: có thể dùng map để nhớ các URL đã lấy, nhưng map thuần
// KHÔNG an toàn khi nhiều goroutine cùng dùng. Hãy bảo vệ nó bằng sync.Mutex
// (xem lại bài 08-mutex-counter).
//
// Hai vấn đề phải giải quyết:
//  1. Tránh trùng lặp: nhiều goroutine có thể cùng thấy một URL một lúc.
//  2. Chờ xong: main phải chờ mọi goroutine con kết thúc rồi mới thoát,
//     nếu không chương trình sẽ kết thúc sớm và không in gì.
//
// Chạy:  go run ./06-concurrency/09-exercise-web-crawler
// Kết quả đúng: mỗi URL trong fakeFetcher được in ĐÚNG MỘT LẦN.
package main

import (
	"fmt"
)

type Fetcher interface {
	// Fetch trả về body của URL và danh sách URL tìm thấy trên trang đó.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl duyệt đệ quy các trang bắt đầu từ url, tới độ sâu tối đa depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: lấy các URL song song.
	// TODO: không lấy trùng URL.
	// Cài đặt dưới đây làm SAI cả hai điều trên.
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher là Fetcher trả về kết quả dựng sẵn.
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

// fetcher là fakeFetcher đã nạp sẵn dữ liệu.
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
