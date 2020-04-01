/*
合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。

示例:

输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6
*/
package main

import (
	"container/list"
	"fmt"
)

func main() {
	//var t *testing.T
	expectList := list.New()
	expectList.PushBack(1)
	expectList.PushBack(1)
	expectList.PushBack(2)
	expectList.PushBack(3)
	expectList.PushBack(4)
	expectList.PushBack(4)
	expectList.PushBack(5)
	expectList.PushBack(6)
	fmt.Println(expectList)

	list0 := list.New()
	list0.PushBack(1)
	list0.PushBack(4)
	list0.PushBack(5)
	list1 := list.New()
	list0.PushBack(1)
	list0.PushBack(3)
	list0.PushBack(4)
	list2 := list.New()
	list0.PushBack(2)
	list0.PushBack(16)

	var lists = []*list.List{list0, list1, list2}
	result := mergeSortListByGreedy(lists)
	fmt.Println(result)
}

func mergeSortListByGreedy(lists []*list.List) *list.List {
	if len(lists) == 0 {
		return nil
	}
	result := list.New()
	for {
		if len(lists) == 1 {
			break
		}
		var front = lists[0].Front().Value.(int)
		j := 0
		for i := 1; i < len(lists); i++ {
			ele := lists[i].Front()
			if ele == nil {
				lists = append(lists[:i], lists[i+1:]...)
				continue
			}
			iValue := ele.Value.(int)
			if iValue < front {
				front = iValue
				j = i
			}
			lists[j].Remove(ele)
		}
		result.PushBack(front)
	}
	return result
}
