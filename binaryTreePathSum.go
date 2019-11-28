/*
给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。

说明: 叶子节点是指没有子节点的节点。

示例:
给定如下二叉树，以及目标和 sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
返回:

[
   [5,4,11,2],
   [5,8,4,5]
]
在真实的面试中遇到过这道题

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/path-sum-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

var inner []int
var total []inner

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	inner = append(inner, root.Val)
	if root.Val == sum && root.Left == nil && root.Right == nil {
		var dumpInner = make([]int, len(inner))
		copy(dumpInner, inner)
		total = append(total, inner)
	}

	sumExcludeRoot := sum - root.Val
	if root.Left != nil {
		pathSum(root.Left, sumExcludeRoot)
	}

	if root.Right != nil {
		pathSum(root.Right, sumExcludeRoot)
	}
	inner = inner[:(len(inner) - 1)]
	return total
}
