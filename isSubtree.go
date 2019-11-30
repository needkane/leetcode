package main

import "fmt"

//从一颗二叉树中任意选取N个节点，判断这个N个节点是否为二叉树的独立子树，其中N个节点对应的指针存储在长度为N的数组中
func main() {

	root := new(TreeNode)
	root.Val = 1
	root.Left = new(TreeNode)
	root.Left.Val = 2
	root.Right = new(TreeNode)
	root.Right.Val = 3

	nodes := []*TreeNode{root, root.Left, root.Right}
	//true
	fmt.Println(isSubtree(nodes))

	node := new(TreeNode)
	node.Val = 4
	nodes = append(nodes, node)
	//false
	fmt.Println(isSubtree(nodes))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(nTree []*TreeNode) bool {

	if nTree == nil {
		return false
	}
	// find root
	root := nTree[0]
	for i := 0; i < len(nTree); i++ {
		if isRoot(nTree[i], root) {
			root = nTree[i]
		}
	}
	_, result := isSub(nTree, root)

	return result
}
func isRoot(root *TreeNode, node *TreeNode) bool {
	if root.Left == node || root.Right == node {
		return true
	}
	if root.Left != nil {
		return isRoot(root.Left, node)
	}
	if root.Right != nil {
		return isRoot(root.Right, node)
	}

	return false
}

func isSub(nodes []*TreeNode, root *TreeNode) ([]*TreeNode, bool) {

	if root == nil {
		return nodes, false
	}
	var isInclude, result bool

	nodes, isInclude = include(nodes, root)
	if len(nodes) == 0 && isInclude {
		return nodes, true
	}
	if root.Left != nil {
		nodes, result = isSub(nodes, root.Left)
		if result {
			return nodes, result
		}
	}
	if root.Right != nil {
		nodes, result = isSub(nodes, root.Right)
		if result {
			return nodes, result
		}
	}
	return nodes, false
}

func include(nodes []*TreeNode, node *TreeNode) ([]*TreeNode, bool) {

	for i := 0; i < len(nodes); i++ {
		if nodes[i] == node {
			nodes = sliceDelEle(nodes, i)
			return nodes, true
		}
	}
	return nodes, false
}

func sliceDelEle(nodes []*TreeNode, temp int) []*TreeNode {
	if temp == 0 {
		nodes = nodes[temp+1:]
	} else if temp == len(nodes)-1 {
		nodes = nodes[:temp]
	} else {
		nodes = append(nodes[:temp], nodes[temp+1:]...)
	}

	return nodes
}
