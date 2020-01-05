/*
给定一个整数数组 nums，将该数组升序排列。



示例 1：

输入：[5,2,3,1]
输出：[1,2,3,5]
示例 2：

输入：[5,1,1,2,0,0]
输出：[0,0,1,1,2,5]


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sort-an-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import "fmt"

func main() {

	nums := []int{3, 9, 6, 2, 5}
	nums = sortArray(nums)
	fmt.Println(nums)
}

//default func of leetcode
func sortArray(nums []int) []int {

	if nums == nil {
		return nil
	}
	if len(nums) == 1 {
		return nums
	}
	quickSort(nums, 0, len(nums)-1)

	return nums
}

func quickSort(nums []int, start, end int) {
	if start < end {
		i := start
		j := end
		key := nums[start]
		for {
			for j > i {
				if nums[j] < key {
					nums[i] = nums[j]
					nums[j] = key
					break
				}
				j--
			}
			for i < j {
				if nums[i] > key {
					nums[j] = nums[i]
					nums[i] = key
					break
				}
				i++
			}

			if i >= j {
				break
			}
		}

		quickSort(nums, start, i-1)
		quickSort(nums, i+1, end)

	}
}
