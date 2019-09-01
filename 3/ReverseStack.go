package main

import "fmt"

type Stack []int

func (s Stack) push(i int) Stack {
	return append(s, i)
}

func (s Stack) pop() (Stack, int) {
	var v int
	v = s[len(s)-1]
	s = s[:len(s)-1]

	return s, v
}

// 取栈底元素
func (s Stack) getTheLast() (Stack, int) {
	s, v := s.pop()
	if len(s) == 0 {
		return s, v
	} else {
		s, l := s.getTheLast()
		s = s.push(v)
		return s, l
	}
}

func reverse(s Stack) Stack {
	if len(s) == 0 {
		return s
	}
	stack, i := s.getTheLast()
	s = reverse(stack)
	return s.push(i)
}

func main() {
	s := make(Stack, 0)
	s = s.push(1)
	s = s.push(2)
	s = s.push(3)
	s = s.push(4)
	fmt.Println(s)

	fmt.Println(reverse(s))
}
