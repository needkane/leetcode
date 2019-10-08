package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func main() {
	var t *testing.T
	origin := []int{41, 24, 76, 11, 45, 64, 21, 69, 19, 36}

	arr := mergeSort(origin)
	sortArr := []int{11, 19, 21, 24, 36, 41, 45, 64, 69, 76}
	fmt.Println(arr)
	assert.Equal(t, arr, sortArr)

}
func merge(left, right []int) (result []int) {
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return
}

func mergeSort(slice []int) (result []int) {
	l := len(slice)
	if l > 1 {
		num := l / 2
		left := mergeSort(slice[:num])
		right := mergeSort(slice[num:])
		return merge(left, right)
	} else {
		return slice
	}
}
