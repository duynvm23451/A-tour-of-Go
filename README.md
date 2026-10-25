# A Tour of Go — bản tái hiện cục bộ

Toàn bộ nội dung của [A Tour of Go](https://go.dev/tour/) được dựng lại thành
**95 chương trình Go chạy được**, kèm **11 bài tập** có đề bài, test tự chấm và
lời giải riêng. Comment song ngữ: thuật ngữ giữ tiếng Anh, giải thích tiếng Việt.

Không cần mạng — các package `golang.org/x/tour/{pic,wc,tree,reader}` đã được
viết lại cục bộ trong [internal/](internal/).

## Cách dùng

Mỗi bài là một thư mục riêng chứa một `main.go` độc lập. Chạy bằng:

```bash
go run ./01-basics/01-packages
go run ./03-moretypes/15-append
```

Với bài tập, sửa `main.go` chỗ đánh dấu `TODO`, rồi tự chấm:

```bash
go test ./02-flowcontrol/08-exercise-loops-and-functions
```

Xem lời giải sau khi đã tự làm:

```bash
cat ./solutions/02-flowcontrol/08-exercise-loops-and-functions/main.go
go run ./solutions/02-flowcontrol/08-exercise-loops-and-functions
```

Vài lệnh hữu ích khác:

```bash
go build ./...            # biên dịch tất cả
go vet ./...              # bắt lỗi thường gặp
gofmt -l .                # liệt kê file chưa đúng định dạng chuẩn
go test ./solutions/...   # toàn bộ lời giải phải PASS
go run -race ./06-concurrency/08-mutex-counter   # phát hiện tranh chấp dữ liệu
```

> `go test ./...` sẽ **FAIL** ở các thư mục bài tập cho tới khi bạn làm xong —
> đó là điều mong muốn, chính là thước đo tiến độ của bạn.

## Lộ trình

### [01-basics/](01-basics/) — Packages, biến, hàm (17 bài)

| # | Bài | Nội dung |
|---|-----|----------|
| 01 | [packages](01-basics/01-packages/main.go) | package main, import path |
| 02 | [imports](01-basics/02-imports/main.go) | import gộp khối |
| 03 | [exported-names](01-basics/03-exported-names/main.go) | chữ hoa = export |
| 04 | [functions](01-basics/04-functions/main.go) | kiểu đứng sau tên |
| 05 | [functions-continued](01-basics/05-functions-continued/main.go) | rút gọn tham số cùng kiểu |
| 06 | [multiple-results](01-basics/06-multiple-results/main.go) | trả về nhiều giá trị |
| 07 | [named-results](01-basics/07-named-results/main.go) | naked return |
| 08 | [variables](01-basics/08-variables/main.go) | `var` |
| 09 | [variables-with-initializers](01-basics/09-variables-with-initializers/main.go) | suy kiểu |
| 10 | [short-variable-declarations](01-basics/10-short-variable-declarations/main.go) | `:=` |
| 11 | [basic-types](01-basics/11-basic-types/main.go) | int, rune, byte, float64... |
| 12 | [zero-values](01-basics/12-zero-values/main.go) | giá trị không |
| 13 | [type-conversions](01-basics/13-type-conversions/main.go) | ép kiểu tường minh |
| 14 | [type-inference](01-basics/14-type-inference/main.go) | untyped constant |
| 15 | [constants](01-basics/15-constants/main.go) | `const` |
| 16 | [numeric-constants](01-basics/16-numeric-constants/main.go) | hằng số độ chính xác cao |
| 17 | [iota](01-basics/17-iota/main.go) | **bổ sung** — enum kiểu Go |

### [02-flowcontrol/](02-flowcontrol/) — Điều khiển luồng (14 bài)

| # | Bài | Nội dung |
|---|-----|----------|
| 01 | [for](02-flowcontrol/01-for/main.go) | vòng lặp duy nhất của Go |
| 02 | [for-continued](02-flowcontrol/02-for-continued/main.go) | bỏ init/post |
| 03 | [for-is-gos-while](02-flowcontrol/03-for-is-gos-while/main.go) | for thay while |
| 04 | [forever](02-flowcontrol/04-forever/main.go) | `for {}`, break, continue |
| 05 | [if](02-flowcontrol/05-if/main.go) | if không ngoặc tròn |
| 06 | [if-with-short-statement](02-flowcontrol/06-if-with-short-statement/main.go) | biến trong scope if |
| 07 | [if-and-else](02-flowcontrol/07-if-and-else/main.go) | early return |
| **08** | [**exercise-loops-and-functions**](02-flowcontrol/08-exercise-loops-and-functions/main.go) | 🧩 **căn bậc hai Newton** |
| 09 | [switch](02-flowcontrol/09-switch/main.go) | không fallthrough mặc định |
| 10 | [switch-evaluation-order](02-flowcontrol/10-switch-evaluation-order/main.go) | thứ tự tính case |
| 11 | [switch-with-no-condition](02-flowcontrol/11-switch-with-no-condition/main.go) | thay if-else dài |
| 12 | [defer](02-flowcontrol/12-defer/main.go) | hoãn thi hành |
| 13 | [defer-multi](02-flowcontrol/13-defer-multi/main.go) | ngăn xếp LIFO |
| 14 | [defer-arguments](02-flowcontrol/14-defer-arguments/main.go) | **bổ sung** — tham số tính khi nào |

### [03-moretypes/](03-moretypes/) — Struct, slice, map (26 bài)

| # | Bài | Nội dung |
|---|-----|----------|
| 01 | [pointers](03-moretypes/01-pointers/main.go) | `&`, `*`, không có pointer arithmetic |
| 02–05 | [structs](03-moretypes/02-structs/main.go), [fields](03-moretypes/03-struct-fields/main.go), [pointers](03-moretypes/04-struct-pointers/main.go), [literals](03-moretypes/05-struct-literals/main.go) | struct |
| 06 | [array](03-moretypes/06-array/main.go) | độ dài là một phần của kiểu |
| 07–10 | [slices](03-moretypes/07-slices/main.go), [pointers](03-moretypes/08-slices-pointers/main.go), [literals](03-moretypes/09-slice-literals/main.go), [bounds](03-moretypes/10-slice-bounds/main.go) | slice và mảng nền |
| 11 | [slice-len-cap](03-moretypes/11-slice-len-cap/main.go) | len vs cap |
| 12 | [nil-slices](03-moretypes/12-nil-slices/main.go) | nil slice dùng được ngay |
| 13 | [making-slices](03-moretypes/13-making-slices/main.go) | `make` |
| 14 | [slices-of-slice](03-moretypes/14-slices-of-slice/main.go) | slice lồng nhau |
| 15 | [append](03-moretypes/15-append/main.go) | khi nào append cấp phát lại |
| 16–17 | [range](03-moretypes/16-range/main.go), [range-continued](03-moretypes/17-range-continued/main.go) | duyệt, `_` |
| **18** | [**exercise-slices**](03-moretypes/18-exercise-slices/main.go) | 🧩 **Pic** |
| 19–22 | [maps](03-moretypes/19-maps/main.go), [literals](03-moretypes/20-map-literals/main.go), [tiếp](03-moretypes/21-map-literals-continued/main.go), [mutating](03-moretypes/22-mutating-maps/main.go) | map, cú pháp comma-ok |
| **23** | [**exercise-maps**](03-moretypes/23-exercise-maps/main.go) | 🧩 **WordCount** |
| 24–25 | [function-values](03-moretypes/24-function-values/main.go), [closures](03-moretypes/25-function-closures/main.go) | hàm là giá trị |
| **26** | [**exercise-fibonacci**](03-moretypes/26-exercise-fibonacci/main.go) | 🧩 **closure Fibonacci** |

### [04-methods/](04-methods/) — Method và interface (25 bài)

Đây là chương quan trọng nhất về tư duy Go.

| # | Bài | Nội dung |
|---|-----|----------|
| 01–03 | [methods](04-methods/01-methods/main.go), [là hàm](04-methods/02-methods-funcs/main.go), [trên kiểu bất kỳ](04-methods/03-methods-continued/main.go) | receiver |
| 04–05 | [pointer receivers](04-methods/04-methods-pointers/main.go), [so với hàm](04-methods/05-methods-pointers-explained/main.go) | sửa được giá trị |
| 06–07 | [indirection](04-methods/06-indirection/main.go), [chiều ngược lại](04-methods/07-indirection-values/main.go) | Go tự chèn `&`/`*` |
| 08 | [chọn receiver](04-methods/08-methods-with-pointer-receivers/main.go) | quy tắc nhất quán |
| 09–10 | [interfaces](04-methods/09-interfaces/main.go), [thỏa mãn ngầm định](04-methods/10-interfaces-are-satisfied-implicitly/main.go) | không có `implements` |
| 11–13 | [interface values](04-methods/11-interface-values/main.go), [chứa nil](04-methods/12-interface-values-with-nil/main.go), [nil thật](04-methods/13-nil-interface-values/main.go) | ⚠️ bẫy `err != nil` kinh điển |
| 14 | [empty interface](04-methods/14-empty-interface/main.go) | `any` |
| 15–16 | [type assertions](04-methods/15-type-assertions/main.go), [type switches](04-methods/16-type-switches/main.go) | lấy lại kiểu cụ thể |
| 17 | [stringer](04-methods/17-stringer/main.go) | `fmt.Stringer` |
| **18** | [**exercise-stringer**](04-methods/18-exercise-stringer/main.go) | 🧩 **IPAddr** |
| 19 | [errors](04-methods/19-errors/main.go) | lỗi là giá trị, không có exception |
| **20** | [**exercise-errors**](04-methods/20-exercise-errors/main.go) | 🧩 **ErrNegativeSqrt** |
| 21 | [reader](04-methods/21-reader/main.go) | `io.Reader` |
| **22** | [**exercise-reader**](04-methods/22-exercise-reader/main.go) | 🧩 **MyReader** |
| **23** | [**exercise-rot13-reader**](04-methods/23-exercise-rot13-reader/main.go) | 🧩 **Reader bọc Reader** |
| 24 | [images](04-methods/24-images/main.go) | `image.Image` |
| **25** | [**exercise-images**](04-methods/25-exercise-images/main.go) | 🧩 **tự cài image.Image** |

### [05-generics/](05-generics/) — Generics (2 bài)

| # | Bài | Nội dung |
|---|-----|----------|
| 01 | [index](05-generics/01-index/main.go) | tham số kiểu, `comparable` |
| 02 | [generic-types](05-generics/02-generic-types/main.go) | `List[T]` |

### [06-concurrency/](06-concurrency/) — Concurrency (11 bài)

| # | Bài | Nội dung |
|---|-----|----------|
| 01 | [goroutines](06-concurrency/01-goroutines/main.go) | `go f()` |
| 02 | [channels](06-concurrency/02-channels/main.go) | `<-`, blocking |
| 03 | [buffered-channels](06-concurrency/03-buffered-channels/main.go) | bộ đệm, deadlock |
| 04 | [range-and-close](06-concurrency/04-range-and-close/main.go) | ai được `close` |
| 05 | [select](06-concurrency/05-select/main.go) | chờ nhiều channel |
| 06 | [default-selection](06-concurrency/06-default-selection/main.go) | non-blocking |
| **07** | [**exercise-equivalent-binary-trees**](06-concurrency/07-exercise-equivalent-binary-trees/main.go) | 🧩 **Walk + Same** |
| 08 | [mutex-counter](06-concurrency/08-mutex-counter/main.go) | `sync.Mutex`, `-race` |
| **09** | [**exercise-web-crawler**](06-concurrency/09-exercise-web-crawler/main.go) | 🧩 **crawl song song, không trùng** |
| 10 | [waitgroup](06-concurrency/10-waitgroup/main.go) | **bổ sung** — `sync.WaitGroup` |
| 11 | [context](06-concurrency/11-context/main.go) | **bổ sung** — hủy và timeout |

## Những chỗ khác bản gốc

- **Bài bổ sung** (không có trong Tour gốc, nhưng cần cho code thật):
  `01-basics/17-iota`, `02-flowcontrol/14-defer-arguments`,
  `06-concurrency/10-waitgroup`, `06-concurrency/11-context`.
- **Bài tập vẽ ảnh**: Tour gốc hiện ảnh ngay trên trình duyệt. Ở đây
  [internal/pic](internal/pic/pic.go) ghi ra file `out.png` ở thư mục gốc và
  in kèm bản xem nhanh bằng ASCII trong terminal.
- **Stub bài tập luôn biên dịch được**: khác Tour gốc (nơi một số đề bài không
  compile cho tới khi bạn viết xong), ở đây stub chạy được nhưng cho kết quả sai,
  để `go test` chấm được ngay từ đầu.

## Sau khi học xong

Tour dạy cú pháp và tư duy, không dạy cách viết một chương trình hoàn chỉnh.
Bước tiếp theo nên là:

1. [Effective Go](https://go.dev/doc/effective_go) và
   [Go Code Review Comments](https://go.dev/wiki/CodeReviewComments) — viết cho
   *đúng chất* Go.
2. `net/http` với `http.ServeMux` của thư viện chuẩn (Go 1.22+ đã hỗ trợ
   method và path params) — **trước khi** đụng tới Gin/Echo/Fiber.
3. `log/slog`, `encoding/json`, `database/sql`, `errors.Is/As`, `errgroup`.
