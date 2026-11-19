package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Logger interface {
	Log(msg string)
}

type ConsoleLogger struct{}

type Server struct {
	Name   string
	Logger Logger // có thể không được gán
}

func (c ConsoleLogger) Log(msg string) {
	fmt.Println("[LOG] " + msg)
}

func (s *Server) Start() {
	fmt.Printf("Server %v đang chạy\n", s.Name)
	if s.Logger != nil {
		s.Logger.Log("some thing")
	}
}

type Circle struct{
	R float64
}

func (c *Circle) Area() float64 {
	if c == nil {
		return 0
	}
	return math.Pi * c.R * c.R
}

func status(s Shape) string {
	if s == nil {
		return "interface nil"
	}
	return fmt.Sprintf("diện tích = %.2f", s.Area())
}

func main() {
	var s Shape

	fmt.Printf("%v, %T\n", s, s)
	fmt.Println(s == nil)

	// s.Area()
	// panic: runtime error: invalid memory address or nil pointer dereference
	// vì s đang có type là nil nên method Area không biết gọi từ type nào nên xảy ra panic

	s1 := &Server{Name: "A", Logger: ConsoleLogger{}} // có logger
	s2 := &Server{Name: "B"}                         // không gán → Logger là interface nil (nil, nil)

	s1.Start()

	s2.Start()

	var c *Circle

	c2 := &Circle{R: 2}

	fmt.Println(status(s))
	fmt.Printf("(%v, %T)\n", s, s)
	fmt.Println(status(c))
	fmt.Printf("(%v, %T)\n", c, c)
	fmt.Println(status(c2))
	fmt.Printf("(%v, %T)\n", c2, c2)

}