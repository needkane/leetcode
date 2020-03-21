/*
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-two-sorted-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	head0 := initList("1->2->4")
	head1 := initList("1->3->4")
	result := mergeTwoLists(head0, head1)
	fmt.Println(listToStr(result))
}

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
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	head := new(ListNode)
	for {
		if l1 == nil {
			insertListNode(head, l2)
			break
		}
		if l2 == nil {
			insertListNode(head, l1)
			break
		}
		if l1.Val < l2.Val {
			//head.addListNode(l1.Val)
			addListNode(head, l1.Val)
			l1 = l1.Next
		} else {
			addListNode(head, l2.Val)
			l2 = l2.Next
		}
	}
	return head.Next
}
