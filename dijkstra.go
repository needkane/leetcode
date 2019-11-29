package main

import (
	"fmt"
	"math"
)

type Graph struct {
	vexs   []int   // 顶点集合
	vexNum int     // 顶点数
	edgNum int     // 边数
	matrix [][]int // 邻接矩阵
}

// 边的结构体
type EdgeData struct {
	start  string // 边的起点
	end    string // 边的终点
	weight int    // 边的权重
}

func main() {}

/*
 * Dijkstra最短路径。
 * 即，统计图(G)中"顶点vs"到其它各个顶点的最短路径。
 *
 * 参数说明：
 *        G -- 图
 *       vs -- 起始顶点(start vertex)。即计算"顶点vs"到其它顶点的最短路径。
 *     prev -- 前驱顶点数组。即，prev[i]的值是"顶点vs"到"顶点i"的最短路径所经历的全部顶点中，位于"顶点i"之前的那个顶点。
 *     dist -- 长度数组。即，dist[i]是"顶点vs"到"顶点i"的最短路径的长度。
 */
func dijkstra(G Graph, vs int, prev []int, dist []int) {
	var i, j, k, min, tmp int
	var flag []int // flag[i]=1表示"顶点vs"到"顶点i"的最短路径已成功获取。

	// 初始化
	for i = 0; i < G.vexNum; i++ {
		flag[i] = 0               // 顶点i的最短路径还没获取到。
		prev[i] = 0               // 顶点i的前驱顶点为0。
		dist[i] = G.matrix[vs][i] // 顶点i的最短路径为"顶点vs"到"顶点i"的权。
	}

	// 对"顶点vs"自身进行初始化
	flag[vs] = 1
	dist[vs] = 0

	// 遍历G.vexnum-1次；每次找出一个顶点的最短路径。
	for i = 1; i < G.vexNum; i++ {
		// 寻找当前最小的路径；
		// 即，在未获取最短路径的顶点中，找到离vs最近的顶点(k)。
		min = math.MaxUint32
		for j = 0; j < G.vexNum; j++ {
			if flag[j] == 0 && dist[j] < min {
				min = dist[j]
				k = j
			}
		}
		// 标记"顶点k"为已经获取到最短路径
		flag[k] = 1

		// 修正当前最短路径和前驱顶点
		// 即，当已经"顶点k的最短路径"之后，更新"未获取最短路径的顶点的最短路径和前驱顶点"。
		for j = 0; j < G.vexNum; j++ {
			if G.matrix[k][j] == math.MaxUint32 {
				tmp = math.MaxUint32
			} else {
				tmp = min + G.matrix[k][j] // 防止溢出

			}
			if flag[j] == 0 && tmp < dist[j] {
				dist[j] = tmp
				prev[j] = k
			}
		}
	}

	// 打印dijkstra最短路径的结果
	fmt.Printf("dijkstra(%c): \n", G.vexs[vs])
	for i = 0; i < G.vexNum; i++ {
		fmt.Printf("  shortest(%c, %c)=%d\n", G.vexs[vs], G.vexs[i], dist[i])
	}
}
