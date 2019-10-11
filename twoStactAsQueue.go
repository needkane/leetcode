package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/karalabe/cookiejar.v2/collections/stack"
)

type Queue struct {
	stackPush stack.Stack
	stackPop  stack.Stack
}

func main() {
	var t *testing.T
	q := Queue{}
	for i := 0; i < 5; i++ {
		q.Push(i)
	}
	assert.Equal(t, q.Pop().(int), 0)
	fmt.Println(q.Pop())
	q.Push(6)
	assert.Equal(t, q.Pop().(int), 2)
}

func (q *Queue) Push(i interface{}) {
	q.stackPush.Push(i)
}
func (q *Queue) Pop() (i interface{}) {
	if q.stackPop.Empty() {
		for {
			if q.stackPush.Empty() {
				break
			}
			j := q.stackPush.Pop()
			q.stackPop.Push(j)
		}
	}
	return q.stackPop.Pop()
}
