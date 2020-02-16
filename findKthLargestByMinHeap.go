/*
在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

示例 1:

输入: [3,2,1,5,6,4] 和 k = 2
输出: 5
示例 2:

输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
输出: 4
说明:

你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/kth-largest-element-in-an-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import "fmt"

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	fmt.Println(findKthLargest(nums, k))
	nums = []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k = 4
	fmt.Println(findKthLargest(nums, k))
}

func findKthLargest(nums []int, k int) int {
	lNums := len(nums)
	if lNums < 1 || k < 1 || k > lNums {
		return -1
	}
	//前K个元素原地建小顶堆
	buildMinHeap(nums, k) //nums 0~(k-1),nums[0] is result
	//遍历剩下元素，比堆顶小，跳过；比堆顶大，交换后重新堆化
	for i := k; i < lNums; i++ {
		if nums[i] < nums[0] {
			continue
		}
		nums[0], nums[i] = nums[i], nums[0]
		heapify(nums, k, 0)
	}
	return nums[0]
}

/*
 * 建堆函数
 * 从倒数第一个非叶子节点开始堆化，倒数第一个非叶子节点下标为 K/2-1
 */
func buildMinHeap(arr []int, k int) {
	for i := k/2 - 1; i >= 0; i-- {
		heapify(arr, k, i)
	}
}

/*
 * 堆化函数
 * 父节点下标i，左右子节点的下标分别为 2*i+1 和 2*i+2
 */

func heapify(arr []int, k, i int) {
	minPos := i //assume nums[i] is minValue
	for {
		left := 2*i + 1
		right := 2*i + 2
		if left < k && arr[left] < arr[minPos] {
			minPos = left
		}
		if right < k && arr[right] < arr[minPos] {
			minPos = right
		}
		if minPos == i {
			break
		}

		arr[i], arr[minPos] = arr[minPos], arr[i]
		i = minPos
	}
}
