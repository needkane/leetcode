/*
二叉树的锯齿形层次遍历

给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

例如：

给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7

    1
    2
    3
    4
    5

返回锯齿形层次遍历如下：

[
  [3],
  [20,9],
  [15,7]
]
*/

package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type simpleStack []*TreeNode

func (s simpleStack) Push(v *TreeNode) simpleStack {
	return append(s, v)
}
func (s simpleStack) Pop() (simpleStack, *TreeNode) {
	l := len(s)
	if l == 0 {
		return s, nil
	}
	node := s[l-1]
	return s[:l-1], node
}
func (s simpleStack) IsEmpty() bool {
	if len(s) == 0 {
		return true
	} else {
		return false
	}
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var levels [][]int
	if root == nil {
		return nil
	}
	var stack_odd, stack_even simpleStack
	stack_even = stack_even.Push(root)
	level := 0
	for {
		if stack_odd.IsEmpty() && stack_even.IsEmpty() {
			break
		}
		// start the current level
		levels = append(levels, []int{})
		if level%2 == 0 {
			level_length := len(stack_even)
			for i := 0; i < level_length; i++ {
				var node *TreeNode
				stack_even, node = stack_even.Pop()
				levels[level] = append(levels[level], node.Val)
				if node.Left != nil {
					stack_odd = stack_odd.Push(node.Left)
				}
				if node.Right != nil {
					stack_odd = stack_odd.Push(node.Right)
				}
			}
		} else {
			level_length := len(stack_odd)
			for i := 0; i < level_length; i++ {
				var node *TreeNode
				stack_odd, node = stack_odd.Pop()
				levels[level] = append(levels[level], node.Val)
				if node.Right != nil {
					stack_even = stack_even.Push(node.Right)
				}
				if node.Left != nil {
					stack_even = stack_even.Push(node.Left)
				}
			}
		}
		level++
	}
	return levels
}
func createRoot(datas []int) (root *TreeNode) {

	if len(datas) == 0 || datas[0] == 0 {
		return
	}
	root = *TreeNode{
		Val: datas[0],
	}
	for i = 1; i < len(datas); i++ {

	}
	return
}

func main() {

	root := createRoot()
}
