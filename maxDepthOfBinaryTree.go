/*
给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明: 叶子节点是指没有子节点的节点。

示例：
给定二叉树 [ 3 , 9 , 20 , null , null , 15 , 7  ]，
		3
     9    20
	    15  7

返回它的最大深度 3 。
*/
package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

type BinaryTree struct {
	Left  *BinaryTree
	Right *BinaryTree
	Value int
}

func main() {
	var t *testing.T
	arr := []int{3, 9, 20, -1, -1, 15, 7} //use -1 instead of null
	//	bt := initBinaryTree(arr)
	assert.Equal(t, 3, maxDepthByArr(arr))
	arr = []int{3, 9, 20, 1, 1, 15, 7, 8} //use -1 instead of null
	assert.Equal(t, 4, maxDepthByArr(arr))
	arr = []int{3, 9, 20, 1, 1, 15, 7, 8, -1, -1, -1, -1, -1, -1, -1, 9} //use -1 instead of null
	assert.Equal(t, 5, maxDepthByArr(arr))
}

func initBinaryTree(arr []int) *BinaryTree {
	l := len(arr)
	if l < 1 {
		return nil
	}
	root := new(BinaryTree)
	root.Value = arr[0]
	if l == 2 {
		root.Left = new(BinaryTree)
		root.Left.Value = arr[1]
	} else if l == 3 {
		root.Left = new(BinaryTree)
		root.Left.Value = arr[1]
		root.Right = new(BinaryTree)
		root.Right.Value = arr[2]
	} else {

		root.Left = new(BinaryTree)
		root.Left.Value = arr[1]
		root.Right = new(BinaryTree)
		root.Right.Value = arr[2]
	}
	return root
}

func maxDepthByArr(arr []int) int {
	var i int
	for i = len(arr) - 1; i >= 0; i-- {
		if arr[i] != -1 {
			break
		}
	}
	return int(math.Log2(float64(i+1)) + 1)
}
func maxDepth() {

}
