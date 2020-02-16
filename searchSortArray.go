/*
假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。

搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。

你可以假设数组中不存在重复的元素。

你的算法时间复杂度必须是 O(log n) 级别。

示例 1:

输入: nums = [4,5,6,7,0,1,2], target = 0
输出: 4
示例 2:

输入: nums = [4,5,6,7,0,1,2], target = 3
输出: -1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/search-in-rotated-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import "fmt"

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(search(nums, 0))
	fmt.Println(search(nums, 3))
}

func search(nums []int, target int) int {
	lNums := len(nums)
	if lNums < 1 {
		return -1
	}
	if lNums == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	left := 0
	right := lNums - 1
	mid := right / 2
	for left <= right {
		if nums[mid] == target {
			return mid
		}
		if nums[left] <= nums[mid] { //left asc
			if nums[left] <= target && nums[mid] > target {
				right = mid - 1 //update right index,narrow the range
			} else {
				left = mid + 1
			}
		} else { //right asc
			if nums[right] >= target && nums[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		mid = left + (right-left)/2
	}
	return -1

}
