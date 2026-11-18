package main

import "fmt"

type Greeter interface {
	Greet()
}

type Person struct {
	Name string
}

func (p *Person) Greet() {
	if p == nil {
		fmt.Println("Xin chào, người lạ!")
	} else {
		fmt.Println("Xin chào, ", p.Name)
	}
}

func describe(g Greeter) {
	fmt.Printf("%v %T\n", g, g)
}

type MyError struct {
	Msg string
}

func (e *MyError) Error() string {
	return e.Msg
}

func check(n int) error {
	if n < 0 {
		return &MyError{"Số âm"}
	}
	return nil
	// sửa như vậy thì nếu để var e *MyError trả về là (nil, *main.MyError) sẽ khác nil
}

func main() {
	var p *Person
	var g Greeter
	g = p
	describe(g)
	g.Greet()

	fmt.Println("p == nil:", p == nil)
	fmt.Println("g == nil:", g == nil)
	g = &Person{Name: "An"}

	describe(g)
	g.Greet()

	var g2 Greeter
	fmt.Println("g2 == nil: ", g2 == nil)

	// interface chỉ nill khi cả value và type đều nil, tuy nhiên g có value là nil tuy nhiên type là *main.Person nên khác nil
	// p có type là *main.Person không phải là interface khi chưa gán giá trị gì chính là nil
	// g2 là interface chưa gán gì nên là nil, nil == nil

	err := check(5)
	if err == nil {
		fmt.Println("Không lỗi")
	} else {
		fmt.Println("Có lỗi")
	}
}