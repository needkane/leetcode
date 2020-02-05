/*
序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。

请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。

示例:

你可以将以下二叉树：

    1
   / \
  2   3
     / \
    4   5

序列化为 "[1,2,3,null,null,4,5]"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "[1,2,3,null,null,4,5]"
	root := deserialize(str)
	result := serialize(root)
	fmt.Printf("serialize %v\n ", str == result)

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deserialize(data string) *TreeNode {
	l := len(data)
	if data == "" || l < 3 || data[0] != '[' || data[l-1] != ']' {
		return nil
	}
	arr := strings.Split(data[1:l-1], ",")
	rootVal, _ := strconv.Atoi(arr[0])
	root := new(TreeNode)
	root.Val = rootVal
	q := []*TreeNode{root}
	i := 0
	for {
		i++
		if i >= len(arr) {
			break
		}
		node := q[0] //BFS
		q = q[1:]    //remove
		if arr[i] != "null" {
			val, _ := strconv.Atoi(arr[i])
			child := new(TreeNode)
			child.Val = val
			node.Left = child
			q = append(q, child)
		}
		i++
		if i >= len(arr) {
			break
		}
		if arr[i] != "null" {
			val, _ := strconv.Atoi(arr[i])
			child := new(TreeNode)
			child.Val = val
			node.Right = child
			q = append(q, child)
		}
	}

	return root
}

func serialize(root *TreeNode) string {
	if root == nil || root.Val == 0 {
		return ""
	}
	q := []*TreeNode{root}
	data := ""
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node != nil {
			q = append(q, node.Left, node.Right)
			data = data + strconv.Itoa(node.Val) + ","
		} else {
			data += "null,"
		}
	}
	data = trimRightNulls(data, "null,")
	data = "[" + data[:len(data)-1] + "]" //trim right comma
	return data
}

func trimRightNulls(data, target string) string {
	lTarget := len(target)
	for len(data) > lTarget && data[len(data)-lTarget:] == target {

		data = data[:len(data)-lTarget]
	}
	return data
}
