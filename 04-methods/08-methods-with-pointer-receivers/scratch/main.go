package main

import "fmt"

type Counter struct {
	Name  string
	Value int
}

func (c *Counter) Get() int { return c.Value }
func (c *Counter) Inc()     { c.Value++ }

// func (c Counter) Reset()   { c.Value = 0 } // có lỗi gì ở đây? // value sẽ không được reset

func (c *Counter) Reset() {
	c.Value = 0
}

type Stack []int

func (s Stack) PushCopy(x int) {
	s = append(s, x)
}

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}

func (s *Stack) Pop() int{
	result := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return result
}

func main() {
	counter := Counter{
		Name:  "Cars",
		Value: 2,
	}

	counter.Inc()
	fmt.Println(counter.Get())

	counter.Inc()
	fmt.Println(counter.Get())

	counter.Inc()
	fmt.Println(counter.Get())

	counter.Reset()
	fmt.Println(counter.Get())

	var s Stack

	s.PushCopy(1)
	fmt.Println(s) // {}

	s.Push(2)
	fmt.Println(s) // {2}

	s.Push(3)
	fmt.Println(s) // {2, 3}

	s.Pop()
	fmt.Println(s) // {2}

	s.Pop()
	fmt.Println(s) // {}

}