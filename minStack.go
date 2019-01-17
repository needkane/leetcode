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

type MinStack struct {
}

var min_val = math.MaxInt32
var st = stack.New()

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	if x <= min_val {
		st.Push(min_val)
		min_val = x
	}
	st.Push(x)
}

func (this *MinStack) Pop() {
	i := st.Pop()
	if i.(int) == min_val {
		min_val = st.Pop().(int)
	}
}

func (this *MinStack) Top() int {
	return st.Top().(int)
}

func (this *MinStack) GetMin() int {
	return min_val
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

}

/**
 * Your MinStack object will be instantiated and called as such:
  * obj := Constructor();
   * obj.Push(x);
    * obj.Pop();
	 * param_3 := obj.Top();
	  * param_4 := obj.GetMin();
*/
