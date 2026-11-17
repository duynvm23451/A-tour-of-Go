package main

import "fmt"

type Notifier interface {
	Notify()
}

type Email struct {
	Address string
}

func (e *Email) Notify() {
	fmt.Println("Gửi email tới ", e.Address)
}

type SMS string

func (s SMS) Notify() {
	fmt.Println("Gửi sms tới", s)
}

func describe(n Notifier) {
	fmt.Printf("%v, %T\n", n, n)
}

func main() {
	var n Notifier
	n = &Email{Address: "abc@x.com"}
	describe(n)
	n.Notify()

	n = SMS("0901234567")
	describe(n)
	n.Notify()

	// n = Email{Address: "abc@y.com"}
	// Lỗi vì Notify dùng pointer receiver (*Email), nên chỉ *Email implement Notifier,
	// 04-methods\11-interface-values\scratch\main.go:37:6: cannot use Email{…} (value of struct type Email) as Notifier value in assignment: Email does not implement Notifier (method Notify has pointer receiver)

	ns := []Notifier{
		&Email{Address: "abc@z.com"},
		SMS("0901234568"),
		&Email{Address: "abc@w.com"},
	}

	for i, n := range ns {
		fmt.Printf("%v %T:", i, n)
		n.Notify()
	}
}