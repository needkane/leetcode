/*                                                                                                                                                                                                                 
给定一个二叉树，返回它的 后序 遍历。
 
示例:
 
输入: [1,null,2,3]
   1
    \
     2
    /
   3
 
输出: [3,2,1]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？
 
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-postorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main
 
func main() {
 
}
 
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
 
func postorderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil
    }
 
    var ss simpleStack
    var temp, result []int
    ss = ss.Push(root)
    for {
        if ss.IsEmpty() {
            break
        }
        var node *TreeNode
        ss, node = ss.Pop()
        temp = append(temp, node.Val)
        if node.Left != nil {
            ss = ss.Push(node.Left)
        }
        if node.Right != nil {
            ss = ss.Push(node.Right)
        }
 
    }
    for i := len(temp) - 1; i >= 0; i-- {
        result = append(result, temp[i])
    }
    return result
}
