/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。

示例 1：

输入：2
输出：2
解释： 有两种方法可以爬到楼顶。

1. 1 阶 + 1 阶
2. 2 阶
示例 2：

输入：3
输出：3
解释： 有三种方法可以爬到楼顶。

1. 1 阶 + 1 阶 + 1 阶
2. 1 阶 + 2 阶
3. 2 阶 + 1 阶
*/
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func main() {
	var t *testing.T
	assert.Equal(t, 3, climbStairs(3))
	assert.Equal(t, 8, climbStairs(5))
}

func climbStairs(n int) (methodsCount int) {
	if n == 0 || n == 1 {
		methodsCount = n
		return
	}

	var dps = make([]int, n+1)
	dps[0] = 0
	dps[1] = 1
	dps[2] = 2
	for i := 3; i <= n; i++ {
		dps[i] = dps[i-1] + dps[i-2]
	}
	methodsCount = dps[n]
	return
}
