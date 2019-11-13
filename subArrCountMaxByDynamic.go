/*
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:

输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
*/
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func main() {
	var t *testing.T
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	assert.Equal(t, 6, subArrCountMax(arr))
}

func subArrCountMax(nums []int) (count int) {
	var i int //subArr last index
	if len(nums) == 1 {
		count = nums[0]
		return
	}
	l := len(nums)
	dps := make([]int, l)
	dps[0] = nums[0]
	count = dps[0]
	for i = 1; i < l; i++ {
		dps[i] = maxInt(dps[i-1], 0) + nums[i]
		count = maxInt(count, dps[i])
	}

	return count
}

func maxInt(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
