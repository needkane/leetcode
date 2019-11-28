/*
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

*/
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func lengthOfLongestSubstring(s string) int {

	last := make(map[rune]int) //该字符上一次出现的位置
	uniqStart := 0             //距离当前下标不含重复字符的起始下标
	maxLength := 0             //当前最大不重复子串的长度
	//获取最长不重复子串的起始位置和长度（start,maxLength）
	st := []rune(s) //golang utf-8
	for i, ch := range st {

		lastI, ok := last[ch]
		if ok && lastI >= uniqStart {
			uniqStart = lastI + 1
		}
		tmpLen := i + 1 - uniqStart
		if tmpLen > maxLength {
			maxLength = tmpLen
		}
		last[ch] = i
	}

	return maxLength
}

func main() {
	var t *testing.T
	assert.Equal(t, lengthOfLongestSubstring("abcabcbb"), 3)
	assert.Equal(t, lengthOfLongestSubstring("bbbbb"), 1)
	assert.Equal(t, lengthOfLongestSubstring("pwwkew"), 3)
	assert.Equal(t, lengthOfLongestSubstring("dvdf"), 3)
}
