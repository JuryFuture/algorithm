package main

import "fmt"

type Stack struct {
	data []int
	min  []int
	size int
}

func (s *Stack) push(i int) {
	if s.size == 0 {
		s.data = append(s.data, i)
		s.min = append(s.min, i)
	} else {
		if i <= s.min[s.size-1] {
			s.data = append(s.data, i)
			s.min = append(s.min, i)
		} else {
			s.data = append(s.data, i)
			s.min = append(s.min, s.min[s.size-1])
		}
	}
	s.size++
}

func (s *Stack) pop() int {
	if s.size == 0 {
		panic("栈已空")
	}
	s.size--
	v := s.data[s.size]

	s.data = s.data[:s.size]
	s.min = s.min[:s.size]
	return v
}

func (s *Stack) getMin() int {
	index := s.size - 1
	return s.min[index]
}

func main() {
	stack := &Stack{make([]int, 0), make([]int, 0), 0}
	stack.push(1)
	stack.push(2)
	stack.push(3)
	stack.push(0)
	stack.push(-1)

	fmt.Println(stack)
	fmt.Println(stack.getMin())
	fmt.Println(stack.pop())

	fmt.Println(stack)
	fmt.Println(stack.getMin())
	fmt.Println(stack.pop())

	fmt.Println(stack)
	fmt.Println(stack.getMin())
	fmt.Println(stack.pop())

	fmt.Println(stack)
	fmt.Println(stack.getMin())
	fmt.Println(stack.pop())

	fmt.Println(stack)
	fmt.Println(stack.getMin())
	fmt.Println(stack.pop())
}
