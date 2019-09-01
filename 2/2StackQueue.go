package main

import "fmt"
import "errors"

type Stack []int

func (s Stack) push(i int) Stack {
	return append(s, i)
}

func (s Stack) pop() (Stack, int) {
	v := s[len(s)-1]
	return s[:len(s)-1], v
}

type Queue struct {
	sPush Stack
	sPop  Stack
}

func (q *Queue) add(i int) {
	q.sPush = q.sPush.push(i)
}

func (q *Queue) poll() (int, error) {
	// sPop 为空时将sPush压入sPop
	if len(q.sPop) == 0 {
		len := len(q.sPush)
		for {
			if len == 0 {
				break
			}

			var v int
			q.sPush, v = q.sPush.pop()
			q.sPop = q.sPop.push(v)

			len--
		}
	}

	if len(q.sPop) == 0 {
		//panic("队列已空")
		return 0, errors.New("队列已空")
	}

	var v int
	q.sPop, v = q.sPop.pop()
	return v, nil
}

func main() {
	q := &Queue{make([]int, 0), make([]int, 0)}
	q.add(1)
	q.add(2)
	q.add(3)

	fmt.Println(q)

	/*defer func() {
		if r := recover();r!= nil {
			fmt.Println("recover")
			fmt.Println(r)
		}
	}()*/

	for {
		v, e := q.poll()
		if e != nil {
			fmt.Println("error")
			fmt.Println(e)
			break
		} else {
			fmt.Println(v)
		}
	}
}
