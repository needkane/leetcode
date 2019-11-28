/*
给定一个字符串s，找到其中最长的回文子序列。可以假设s的最大长度为1000。

示例 1:
输入:

"bbbab"
输出:

4
一个可能的最长回文子序列为 "bbbb"。

示例 2:
输入:

"cbbd"
输出:

2
一个可能的最长回文子序列为 "bb"。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-palindromic-subsequence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

func main() {}

func longestPalindromeSubseq(s string) int {
	bytez := []byte(s) //maybe []rune
	l := len(s)
	var dp = make([][]int, l)

	for i := 0; i < l; i++ {
		dp[i] = make([]int, l)
	}
	for i := l - 1; i >= 0; i-- { //end to start
		dp[i][i] = 1 // i to i subSequence
		for j := i + 1; j < l; j++ {
			if bytez[i] == bytez[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = maxInt(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][l-1] //0 to l-1 represent whole s

}

func maxInt(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
