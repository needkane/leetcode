/*
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。



示例：

给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import "fmt"

func main() {

	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
	nums = []int{-2, 0, 0, 2, 2}
	fmt.Println(threeSum(nums))
	nums = []int{-2, 0, 1, 1, 2}
	fmt.Println(threeSum(nums))
	nums = []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	lN := len(nums)
	if len(nums) < 3 {
		return nil
	}
	var ans [][]int
	quickSort(nums, 0, len(nums)-1)
	for i := 0; i < lN-1; i++ {
		if nums[i] > 0 {
			break
		}
		//redundant
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		L := i + 1
		R := lN - 1
		for L < R {

			sum := nums[i] + nums[L] + nums[R]
			if sum == 0 {
				arr := []int{nums[i], nums[L], nums[R]}
				ans = append(ans, arr)
				for L < R && nums[L] == nums[L+1] {
					L++
				}
				for L < R && nums[R] == nums[R-1] {
					R--
				}
				L++
				R--
			} else if sum > 0 {
				R--
			} else if sum < 0 {
				L++
			}
		}
	}
	return ans
}

func quickSort(arr []int, left, right int) {
	if left < right {
		i := left
		j := right
		key := arr[left]
		for {
			for j > i {
				if arr[j] < key {
					arr[i] = arr[j]
					arr[j] = key
					break
				}
				j--
			}
			for i < j {
				if arr[i] > key {
					arr[j] = arr[i]
					arr[i] = key
					break
				}
				i++
			}
			if i >= j {
				break
			}
		}
		quickSort(arr, left, j)
		quickSort(arr, j+1, right)
	}

}
