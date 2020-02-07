/*
给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。

换句话说，第一个字符串的排列之一是第二个字符串的子串。

示例1:

输入: s1 = "ab" s2 = "eidbaooo"
输出: True
解释: s2 包含 s1 的排列之一 ("ba").


示例2:

输入: s1= "ab" s2 = "eidboaoo"
输出: False


注意：

输入的字符串只包含小写字母
两个字符串的长度都在 [1, 10,000] 之间

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutation-in-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import "fmt"

func main() {
	s1 := "ab"
	s2 := "eidbaooo"
	fmt.Println(checkInclusion(s1, s2))
	s1 = "ab"
	s2 = "eidboaoo"
	fmt.Println(checkInclusion(s1, s2))
	s1 = "adc"
	s2 = "dcda"
	fmt.Println(checkInclusion(s1, s2))
}
func checkInclusion(s1 string, s2 string) bool {
	lS1 := len(s1)
	lS2 := len(s2)
	if lS2 < lS1 {
		return false
	}
	var charAppearTimesS1 [26]int
	for i := 0; i < lS1; i++ {
		charAppearTimesS1[[]byte(s1)[i]-97]++
	}
	for i := 0; i <= lS2-lS1; i++ {
		var charAppearTimesS2 [26]int
		for j := 0; j < lS1; j++ {
			charAppearTimesS2[[]byte(s2)[i+j]-97]++
		}
		if matches(charAppearTimesS1, charAppearTimesS2) {
			return true
		}
	}
	return false
}

func matches(arr0, arr1 [26]int) bool {
	for i := 0; i < len(arr0); i++ {
		if arr0[i] != arr1[i] {
			return false
		}
	}
	return true
}
