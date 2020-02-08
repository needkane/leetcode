/*
给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

示例 1:

输入: num1 = "2", num2 = "3"
输出: "6"
示例 2:

输入: num1 = "123", num2 = "456"
输出: "56088"
说明：

num1 和 num2 的长度小于110。
num1 和 num2 只包含数字 0-9。
num1 和 num2 均不以零开头，除非是数字 0 本身。
不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/multiply-strings
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	num1 := "2"
	num2 := "3"
	fmt.Println(multiply(num1, num2))
	num1 = "123"
	num2 = "456"
	fmt.Println(multiply(num1, num2))

}
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	lNum1 := len(num1)
	lNum2 := len(num2)
	arr := make([]int, lNum1+lNum2)
	var high int
	for i := lNum1 - 1; i >= 0; i-- {
		index := lNum1 - 1 - i
		num1Val := int([]byte(num1)[i] - 48)
		for j := len(num2) - 1; j >= 0; j-- {
			num2Val := int([]byte(num2)[j] - 48)
			tmp := num1Val * num2Val
			if tmp > 9 {
				high = tmp / 10
				low := tmp % 10
				arr[index] += low
				if high > 0 {
					arr[index+1] += high
					high = 0
				}
			} else {
				arr[index] += tmp
			}
			if arr[index] > 9 {
				high = arr[index] / 10
				low := arr[index] % 10
				arr[index] = low
				if high > 0 {
					arr[index+1] += high
					high = 0
				}
			}
			index++
		}
	}
	var result string
	removePrefixZero := true
	for i := len(arr) - 1; i >= 0; i-- {
		if removePrefixZero {
			if arr[i] == 0 {
				continue
			} else {
				removePrefixZero = false
			}
		}
		result += strconv.Itoa(arr[i])
	}
	return result
}
