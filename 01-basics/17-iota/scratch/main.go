package main

import "fmt"

type Level int

const (
	Debug Level = iota + 1
	Info
	Warn
	Error
)

const (
	a = iota * 2 // 0
	b            // 2
	c = 100      // 100
	d            // 100
	e = iota     // 4
)

const (
	Read  = 1 << iota 
	Write             
	Delete
	Admin
)

func main() {
	fmt.Printf("%v\n", Read | Write | Delete)
}