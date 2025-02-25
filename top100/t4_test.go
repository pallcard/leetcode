package top100

import (
	"fmt"
	"testing"
)

func Test_findMedianSortedArrays(t *testing.T) {
	res := findMedianSortedArrays([]int{1, 2}, []int{-1, 3})
	fmt.Println(res)
}
