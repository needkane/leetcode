/*
给出集合 [1,2,3,…,n]，其所有元素共有 n! 种排列。

按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：

"123"
"132"
"213"
"231"
"312"
"321"
给定 n 和 k，返回第 k 个排列。

说明：

给定 n 的范围是 [1, 9]。
给定 k 的范围是[1,  n!]。
示例 1:

输入: n = 3, k = 3
输出: "213"
示例 2:

输入: n = 4, k = 9
输出: "2314"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutation-sequence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import "fmt"

func main() {

	n := 3
	k := 3
	fmt.Println(getPermutation(n, k))
	n = 4
	k = 9
	fmt.Println(getPermutation(n, k))
}

func getPermutation(n int, k int) string {
	if n < 1 || n > 9 || k < 1 {
		return ""
	}
	maxK := 1
	for i := 1; i <= n; i++ {
		maxK *= i
	}
	if k > maxK {
		return ""
	}
	if n == 1 {
		if k != 1 {
			return ""
		} else {
			return "1"
		}
	}

	str := "123456789"[:n]
	bytez := []byte(str)
	var resultBytes []byte
	for {
		if len(bytez) == 0 {
			break
		}
		level := k / (maxK / n)
		k = k % (maxK / n)
		if k != 0 {
			resultBytes = append(resultBytes, bytez[level])
			bytez = append(bytez[:level], bytez[level+1:]...)
		} else {
			resultBytes = append(resultBytes, bytez[level-1])
			bytez = append(bytez[:level-1], bytez[level:]...)
			break
		}
		maxK = maxK / n
		n = n - 1
		if n < 1 {
			break
		}
	}
	for i := len(bytez) - 1; i >= 0; i-- {
		resultBytes = append(resultBytes, bytez[i])
	}
	return string(resultBytes)
}
