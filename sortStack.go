package main

import (
	"fmt"

	"gopkg.in/karalabe/cookiejar.v2/collections/stack"
)

func main() {
	var originStack = stack.New()
	var sortStack = stack.New()
	originStack.Push(1)
	originStack.Push(6)
	originStack.Push(7)
	originStack.Push(3)
	sortStack.Push(originStack.Pop())
	for {
		if originStack.Size() == 0 {
			break
		}
		var tmp = originStack.Pop()
		for {
			if !sortStack.Empty() && tmp.(int) > sortStack.Top().(int) {
				originStack.Push(sortStack.Pop())
			} else {
				break
			}
		}
		sortStack.Push(tmp)
	}

	for {
		if sortStack.Empty() {
			break
		}
		fmt.Println(sortStack.Pop())
	}
}
