/*
给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。

示例:

输入: "25525511135"
输出: ["255.255.11.135", "255.255.111.35"]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/restore-ip-addresses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "25525511135"
	fmt.Println(restoreIpAddresses(str))
	str = "0000"
	fmt.Println(restoreIpAddresses(str))
}
func restoreIpAddresses(s string) []string {
	lS := len(s)
	if lS < 4 {
		return nil
	}
	var result []string
	for i := 1; i < 4; i++ {
		if !valid(s[:i]) || lS-i < 3 {
			break
		}
		for j := i + 1; j < i+4; j++ {
			if !valid(s[i:j]) || lS-j < 2 {
				break
			}

			for k := j + 1; k < lS; k++ {
				if !valid(s[j:k]) {
					break
				}
				if valid(s[k:]) {
					ip := s[:i] + "." + s[i:j] + "." + s[j:k] + "." + s[k:]
					result = append(result, ip)
				}
			}
		}
	}
	return result
}

/*
   Check if the current segment is valid :
   1. less or equal to 255
   2. the first character could be '0'
   only if the segment is equal to '0'
*/
func valid(segment string) bool {
	lSeg := len(segment)
	if lSeg < 1 || lSeg > 3 {
		return false
	}
	if lSeg > 1 && segment[0] == '0' {
		return false
	}
	val, err := strconv.Atoi(segment)
	if err != nil {
		return false
	}
	if val > 255 {
		return false
	}
	return true
}
