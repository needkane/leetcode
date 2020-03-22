/*
2. 两数相加
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s0 := "2->4->3"
	s1 := "5->6->4"
	result := addTwoNumbers(initList(s0), initList(s1))
	fmt.Println(listToStr(result))
}

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func initList(str string) *ListNode {
	arr := strings.Split(str, "->")
	head := new(ListNode)
	head.Val, _ = strconv.Atoi(arr[0])
	for i := 1; i < len(arr); i++ {
		iVal, _ := strconv.Atoi(arr[i])
		//head.addListNode(iVal)
		addListNode(head, iVal)
	}
	return head
}

//for leetcode,func (node *ListNode) addListNode(v int)
func addListNode(node *ListNode, v int) {

	cur := new(ListNode)
	cur.Val = v
	if node.Next == nil {
		node.Next = cur
	} else {
		addListNode(node.Next, v)
	}
}

// for leetcode, (node *ListNode)insertListNode(n *ListNode)
func insertListNode(node *ListNode, n *ListNode) {

	if node.Next == nil {
		node.Next = n
	} else {
		insertListNode(node.Next, n)
	}
}
func listToStr(head *ListNode) string {
	if head == nil {
		return ""
	}
	var result string
	for {
		if head == nil {
			break
		}
		result = result + strconv.Itoa(head.Val) + "->"
		head = head.Next
	}

	return result[:len(result)-2]
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}

	var result = new(ListNode)
	result.Val = 0
	curr := result
	carry := 0
	for {
		if l1 == nil && l2 == nil {
			break
		}
		x := 0
		y := 0
		if l1 != nil {
			x = l1.Val
		}
		if l2 != nil {
			y = l2.Val
		}
		sum := carry + x + y
		carry = sum / 10
		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if carry > 0 {
		curr.Next = &ListNode{Val: carry}
	}
	return result.Next
}
