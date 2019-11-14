package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Node struct {
	Value int
	Next  *Node
}

func main() {
	var t *testing.T
	arr := []int{1, 3, 5, 7, 2}
	head := initList(arr)
	result := listToArr(head)
	fmt.Println(result)
	assert.Equal(t, arr, result)
	reverseResult := listToArr(reverseList(head))
	fmt.Println(reverseResult)
}
func reverseList(head *Node) *Node {
	var Pre *Node
	var Next *Node
	for {
		if head == nil {
			break
		}
		Next = head.Next
		head.Next = Pre
		Pre = head
		head = Next
	}
	return Pre
}
func initList(arr []int) *Node {
	if len(arr) == 0 {
		return nil
	}
	head := new(Node)
	head.Value = arr[0]
	for i := 1; i < len(arr); i++ {
		head.addNode(arr[i])
	}
	return head
}
func (node *Node) addNode(v int) {
	cur := new(Node)
	cur.Value = v
	if node.Next == nil {
		node.Next = cur
	} else {
		node.Next.addNode(v)
	}
}

func listToArr(head *Node) []int {
	result := []int{}
	for {
		if head == nil {
			break
		}
		result = append(result, head.Value)
		head = head.Next
	}
	return result
}
