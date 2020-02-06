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
	sortS1 := sort(s1)
	lS1 := len(s1)
	lS2 := len(s2)
	if lS2 < lS1 {
		return false
	}
	if lS2 == lS1 {
		sortS2 := sort(s2)
		return sortS2 == sortS1
	}
	var arrFromS2 []string
	for i := 0; i <= lS2-lS1; i++ {
		arrFromS2 = append(arrFromS2, s2[i:i+lS1])
	}
	for _, v := range arrFromS2 {
		sortV := sort(v)
		if sortV == sortS1 {
			return true
		}
	}
	return false
}

func sort(s string) string {
	arr := []rune(s)
	quickSort(arr, 0, len(arr)-1)
	return string(arr)
}

func quickSort(arr []rune, left, right int) {
	if left < right {
		key := arr[left]
		i := left
		j := right
		for {

			for j > i {
				if arr[j] < key {
					arr[i] = arr[j]
					arr[j] = key
					// i++
					break
				}
				j--
			}
			for i < j {
				if arr[i] > key {
					arr[j] = arr[i]
					arr[i] = key
					// j--
					break
				}
				i++
			}
			if i >= j {
				break
			}
		}

		quickSort(arr, left, j-1)
		quickSort(arr, j+1, right)
	}

}
