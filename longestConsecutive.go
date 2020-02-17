/*
给定一个未排序的整数数组，找出最长连续序列的长度。

要求算法的时间复杂度为 O(n)。

示例:

输入: [100, 4, 200, 1, 3, 2]
输出: 4
解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-consecutive-sequence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import "fmt"

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(nums))
}

func longestConsecutive(nums []int) int {
	lNums := len(nums)
	if lNums < 1 {
		return 0
	}
	if lNums == 1 {
		return 1
	}
	buildMaxHeap(nums, lNums)
	for i := lNums - 1; i >= 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		heaprify(nums, i, 0)
	}
	var maxLength = 1
	var tmpLength = 1
	for i := 1; i < lNums; i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		if nums[i] == nums[i-1]+1 {
			tmpLength += 1
		} else {
			tmpLength = 1
		}
		if tmpLength > maxLength {
			maxLength = tmpLength
		}
	}

	return maxLength

}

func buildMaxHeap(arr []int, k int) {
	for i := k / 2; i >= 0; i-- {
		heaprify(arr, k, i)
	}
}

func heaprify(arr []int, k, i int) {
	maxPos := i //assume arr[i] is maxValue
	for {
		left := 2*i + 1
		right := 2*i + 2
		if left < k && arr[left] > arr[maxPos] {
			maxPos = left
		}
		if right < k && arr[right] > arr[maxPos] {
			maxPos = right
		}
		if maxPos == i {
			break
		}

		arr[i], arr[maxPos] = arr[maxPos], arr[i]
		i = maxPos
	}
}
