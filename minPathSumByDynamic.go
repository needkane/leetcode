/*
leetcode 64
给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。

示例:

输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 7
解释: 因为路径 1→3→1→1→1 的总和最小。

*/

package main

import "fmt"

func main() {

	var grid = [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	fmt.Println(minPathSum(grid))
}

func minPathSum(grid [][]int) int {

	row := len(grid)
	col := len(grid[0])
	dp := make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
	}

	dp[0][0] = grid[0][0]
	for i := 1; i < row; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < col; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			dp[i][j] = minInt(dp[i][j-1], dp[i-1][j]) + grid[i][j]
		}
	}
	return dp[row-1][col-1]

}

func minInt(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
