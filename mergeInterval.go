/*
给出一个区间的集合，请合并所有重叠的区间。

示例 1:

输入: [[1,3],[2,6],[8,10],[15,18]]
输出: [[1,6],[8,10],[15,18]]
解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2:

输入: [[1,4],[4,5]]
输出: [[1,5]]
解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-intervals
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import "fmt"

func main() {
	intervals := [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}
	fmt.Println(merge(intervals))

	intervals = [][]int{
		{1, 4},
		{4, 5},
	}
	fmt.Println(merge(intervals))
	intervals = [][]int{
		{1, 4},
		{0, 4},
	}
	fmt.Println(merge(intervals))
	intervals = [][]int{
		{1, 4},
		{0, 0},
	}
	fmt.Println(merge(intervals))
	intervals = [][]int{
		{2, 3},
		{4, 5},
		{6, 7},
		{8, 9},
		{1, 10},
	}
	fmt.Println(merge(intervals))
}

func merge(intervals [][]int) [][]int {
	lI := len(intervals)
	if lI <= 1 {
		return intervals
	}
	quickSort(intervals, 0, lI-1)
	var result [][]int
	start := intervals[0][0]
	end := intervals[0][1]
	for i := 1; i < lI; i++ {
		if intervals[i-1][1] < intervals[i][0] && intervals[i][0] > end {
			result = append(result, []int{start, end})
			start = intervals[i][0]
			end = intervals[i][1]

		} else {
			if end < intervals[i][1] {
				end = intervals[i][1]
			}
		}
		if i == lI-1 {
			result = append(result, []int{start, end})
		}
	}

	return result
}

func quickSort(intervals [][]int, left, right int) {
	if left < right {
		key := intervals[left]
		i := left
		j := right
		for {

			for j > i {
				if intervals[j][0] < key[0] {
					intervals[i] = intervals[j]
					intervals[j] = key
					i++
					break
				}
				j--
			}
			for i < j {
				if intervals[i][0] > key[0] {
					intervals[j] = intervals[i]
					intervals[i] = key
					j--
					break
				}
				i++
			}
			if i >= j {
				break
			}
		}
		quickSort(intervals, left, j-1)
		quickSort(intervals, j+1, right)
	}
}
