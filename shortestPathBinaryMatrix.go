/*
在一个 N × M 的网格中，每个单元格有两种状态：空（0）或者阻塞（1）。

一条从左上角到右下角、长度为 k 的畅通路径，由满足下述条件的单元格 C_1, C_2, ..., C_k 组成：

相邻单元格 C_i 和 C_{i+1} 在8个方向之一上连通（此时，C_i 和 C_{i+1} 不同且共享边或角）
C_1 位于 (0, 0)（即，值为 grid[0][0]）
C_k 位于 (N-1, N-1)（即，值为 grid[N-1][N-1]）
如果 C_i 位于 (r, c)，则 grid[r][c] 为空（即，grid[r][c] == 0）
返回这条从左上角到右下角的最短畅通路径的长度。如果不存在这样的路径，返回 -1 。



示例 1：

输入：[[0,1],[1,0]]

输出：2

示例 2：

输入：[[0,0,0],[1,1,0],[1,1,0]]

输出：4



提示：

1 <= grid.length == grid[0].length <= 100
grid[i][j] 为 0 或 1

*/
package main

func main() {

}
func shortestPathBinaryMatrix(grid [][]int) int {
	if grid == nil || len(grid[0]) == 0 || grid[0][0] == 1 {
		return -1
	}
	row := len(grid)
	col := len(grid[0])
	if grid[row-1][col-1] != 0 {
		return -1
	}

	var shortestPaths = make([][]int, row) //(0,0) to (i,j)
	var rQ, cQ []int                       //queue for store j and j
	var r, c int
	for i := 0; i < row; i++ {
		shortestPaths[i] = make([]int, col)
	}
	shortestPaths[0][0] = 1
	rQ = append(rQ, 0)
	cQ = append(cQ, 0)
	for {
		if len(rQ) == 0 {
			break
		}
		r = rQ[0]
		c = cQ[0]
		rQ = rQ[1:]
		cQ = cQ[1:]
		if r == row-1 && c == col-1 {
			return shortestPaths[r][c]
		}
		walkTo(shortestPaths[r][c], r-1, c-1, grid, shortestPaths, rQ, cQ) //left up
		walkTo(shortestPaths[r][c], r-1, c, grid, shortestPaths, rQ, cQ)   //up
		walkTo(shortestPaths[r][c], r-1, c+1, grid, shortestPaths, rQ, cQ) //right up
		walkTo(shortestPaths[r][c], r, c-1, grid, shortestPaths, rQ, cQ)   //left
		walkTo(shortestPaths[r][c], r, c+1, grid, shortestPaths, rQ, cQ)   //right
		walkTo(shortestPaths[r][c], r+1, c-1, grid, shortestPaths, rQ, cQ) //left down
		walkTo(shortestPaths[r][c], r+1, c, grid, shortestPaths, rQ, cQ)   //down
		walkTo(shortestPaths[r][c], r+1, c+1, grid, shortestPaths, rQ, cQ) //right down

	}
	return -1
}
func walkTo(pre, toR, toC int, grid, shortestPaths [][]int, rQ, cQ []int) {
	if toR < 0 || toR == len(grid) || toC < 0 || toC == len(grid[0]) || grid[toR][toC] != 0 || shortestPaths[toR][toC] != 0 {
		return
	}
	shortestPaths[toR][toC] = pre + 1
	rQ = append(rQ, toR)
	cQ = append(cQ, toC)
}
