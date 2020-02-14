/*
给定一个包含了一些 0 和 1的非空二维数组 grid , 一个 岛屿 是由四个方向 (水平或垂直) 的 1 (代表土地) 构成的组合。你可以假设二维矩阵的四个边缘都被水包围着。

找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为0。)

示例 1:

[[0,0,1,0,0,0,0,1,0,0,0,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,1,1,0,1,0,0,0,0,0,0,0,0],
 [0,1,0,0,1,1,0,0,1,0,1,0,0],
 [0,1,0,0,1,1,0,0,1,1,1,0,0],
 [0,0,0,0,0,0,0,0,0,0,1,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,0,0,0,0,0,0,1,1,0,0,0,0]]
对于上面这个给定矩阵应返回 6。注意答案不应该是11，因为岛屿只能包含水平或垂直的四个方向的‘1’。

示例 2:

[[0,0,0,0,0,0,0,0]]
对于上面这个给定的矩阵, 返回 0。

注意: 给定的矩阵grid 的长度和宽度都不超过 50。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/max-area-of-island
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import "fmt"

func main() {
	grid := [][]int{
		[]int{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		[]int{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		[]int{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}
	fmt.Println(maxAreaOfIsland(grid))
	grid = [][]int{
		[]int{0, 0, 0, 0, 0, 0, 0, 0},
	}
	fmt.Println(maxAreaOfIsland(grid))
}
func maxAreaOfIsland(grid [][]int) int {
	lGrid := len(grid)
	if lGrid < 1 {
		return 0
	}
	hasCount := make([][]bool, lGrid)
	lG0 := len(grid[0])
	for i := 0; i < lGrid; i++ {
		hasCount[i] = make([]bool, lG0)
	}
	var maxArea int
	for i := 0; i < lGrid; i++ {
		for j := 0; j < lG0; j++ {
			tmp := area(grid, hasCount, i, j)
			if tmp > maxArea {
				maxArea = tmp
			}
		}
	}
	return maxArea
}

func area(grid [][]int, hasCount [][]bool, row, col int) int {
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) || hasCount[row][col] || grid[row][col] == 0 {
		return 0
	}
	hasCount[row][col] = true
	countPreRow := area(grid, hasCount, row-1, col)
	countPreCol := area(grid, hasCount, row, col-1)
	countNextRow := area(grid, hasCount, row+1, col)
	countNextCol := area(grid, hasCount, row, col+1)
	return 1 + countPreRow + countPreCol + countNextRow + countNextCol
}
