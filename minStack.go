/*
Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

push(x) -- Push element x onto stack.
pop() -- Removes the element on top of the stack.
top() -- Get the top element.
getMin() -- Retrieve the minimum element in the stack.


Example:

MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> Returns -3.
minStack.pop();
minStack.top();      --> Returns 0.
minStack.getMin();   --> Returns -2.
*/
package main

import (
	"fmt"
	"math"

	"gopkg.in/karalabe/cookiejar.v2/collections/stack"
)

const blockSize = 4096

// Last in, first out data structure.
type Stack struct {
	size     int
	capacity int
	offset   int

	blocks [][]interface{}
	active []interface{}
}

// Creates a new, empty stack.
func New() *Stack {
	result := new(Stack)
	result.active = make([]interface{}, blockSize)
	result.blocks = [][]interface{}{result.active}
	result.capacity = blockSize
	return result
}

// Pushes a value onto the stack, expanding it if necessary.
func (s *Stack) Push(data interface{}) {
	if s.size == s.capacity {
		s.active = make([]interface{}, blockSize)
		s.blocks = append(s.blocks, s.active)
		s.capacity += blockSize
		s.offset = 0
	} else if s.offset == blockSize {
		s.active = s.blocks[s.size/blockSize]
		s.offset = 0
	}
	s.active[s.offset] = data
	s.offset++
	s.size++
}

// Pops a value off the stack and returns it. Currently no shrinking is done.
func (s *Stack) Pop() (res interface{}) {
	s.size--
	s.offset--
	if s.offset < 0 {
		s.offset = blockSize - 1
		s.active = s.blocks[s.size/blockSize]
	}
	res, s.active[s.offset] = s.active[s.offset], nil
	return
}

// Returns the value currently on the top of the stack. No bounds are checked.
func (s *Stack) Top() interface{} {
	if s.offset > 0 {
		return s.active[s.offset-1]
	} else {
		return s.blocks[(s.size-1)/blockSize][blockSize-1]
	}
}

// Checks whether the stack is empty or not.
func (s *Stack) Empty() bool {
	return s.size == 0
}

// Returns the number of elements in the stack.
func (s *Stack) Size() int {
	return s.size
}

// Resets the stack, effectively clearing its contents.
func (s *Stack) Reset() {
	*s = *New()
}

type MinStack struct {
	min_val int
}

var st = stack.New()

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{math.MaxInt32}
}

func (this *MinStack) Push(x int) {
	if x <= this.min_val {
		st.Push(this.min_val)
		this.min_val = x
	}
	st.Push(x)
}

func (this *MinStack) Pop() {
	i := st.Pop()
	if i.(int) == this.min_val {
		this.min_val = st.Pop().(int)
	}
}

func (this *MinStack) Top() int {
	return st.Top().(int)
}

func (this *MinStack) GetMin() int {
	return this.min_val
}

func main() {

	mst := Constructor()
	mst.Push(-2)
	mst.Push(0)
	mst.Push(-3)
	fmt.Println(mst.GetMin())
	mst.Pop()
	fmt.Println(mst.Top())
	fmt.Println(mst.GetMin())

	mst2 := Constructor()
	mst2.Push(-1)
	fmt.Println(mst2.Top())
	fmt.Println(mst2.GetMin())

}

/**
 * Your MinStack object will be instantiated and called as such:
  * obj := Constructor();
   * obj.Push(x);
    * obj.Pop();
	 * param_3 := obj.Top();
	  * param_4 := obj.GetMin();
*/
