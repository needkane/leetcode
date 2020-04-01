/*
合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。

示例:

输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-k-sorted-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lists := []*ListNode{initList("1->4->5"), initList("1->3->4"), initList("2->6")}
	result := mergeKLists(lists)
	fmt.Println(listToStr(result))
	lists = []*ListNode{initList("2"), initList(""), initList("-1")}
	result = mergeKLists(lists)
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

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) < 1 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	result := new(ListNode)
	curr := result
	for {
		if lists[0] == nil {
			if len(lists) >= 2 {
				lists = lists[1:]
				continue
			} else {
				break
			}
		}
		if len(lists) <= 1 {
			if lists[0] != nil {
				curr.Next = lists[0]
			}
			break
		}
		var front = lists[0].Val
		j := 0
		for i := 1; i < len(lists); i++ {
			ele := lists[i]
			if ele == nil {
				lists = append(lists[:i], lists[i+1:]...)
				continue
			}
			iValue := ele.Val
			if iValue < front {
				front = iValue
				j = i
			}
		}
		lists[j] = lists[j].Next
		curr.Next = &ListNode{Val: front}
		curr = curr.Next
	}
	return result.Next

}
