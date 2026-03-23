package main

import "fmt"

type stack struct { d []int }

func (s *stack) push(x int) { s.d = append(s.d, x) }
func (s *stack) pop() (int, bool) {
	if len(s.d)==0 { return 0,false }
	x := s.d[len(s.d)-1]
	s.d = s.d[:len(s.d)-1]
	return x, true
}

type queue struct { d []int }

func (q *queue) enq(x int) { q.d = append(q.d, x) }
func (q *queue) deq() (int, bool) {
	if len(q.d)==0 { return 0,false }
	x := q.d[0]
	q.d = q.d[1:]
	return x,true
}

func main() {
	fmt.Println("struct tests")
	all := true

	s := &stack{}
	s.push(10)
	s.push(20)
	v1,ok1 := s.pop()
	v2,ok2 := s.pop()
	_,ok3 := s.pop()
	if ok1 && ok2 && !ok3 && v1==20 && v2==10 {
		fmt.Println("stack ok")
	} else {
		all = false
		fmt.Println("stack fail")
	}

	q := &queue{}
	q.enq(1)
	q.enq(2)
	q.enq(3)
	a,oa := q.deq()
	b,ob := q.deq()
	q.enq(99)
	c,oc := q.deq()
	d,od := q.deq()
	_,oe := q.deq()
	if oa&&ob&&oc&&od&&!oe&&a==1&&b==2&&c==3&&d==99 {
		fmt.Println("queue ok")
	} else {
		all=false
		fmt.Println("queue fail")
	}

	if all { fmt.Println("struct pass") } else { fmt.Println("struct fail") }
}
