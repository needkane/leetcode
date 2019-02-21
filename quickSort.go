package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func quickSort(arr []int, left, right int) {
	if left < right {
		key := arr[left]
		i := left
		j := right
		for {

			for j > i {
				if arr[j] < key {
					arr[i] = arr[j]
					arr[j] = key
					break
				}
				j--
			}
			for i < j {
				if arr[i] > key {
					arr[j] = arr[i]
					arr[i] = key
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

func quickSort2() {

}
func main() {
	var t *testing.T
	arr := []int{41, 24, 76, 11, 45, 64, 21, 69, 19, 36}

	right := len(arr) - 1

	quickSort(arr, 0, right)
	sortArr := []int{11, 19, 21, 24, 36, 41, 45, 64, 69, 76}
	assert.Equal(t, arr, sortArr)
	fmt.Println(arr)
}
