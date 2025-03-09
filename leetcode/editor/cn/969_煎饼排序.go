package main

import "sort"

// leetcode submit region begin(Prohibit modification and deletion)
func pancakeSort(arr []int) []int {
	// todo 有思路按照思路写不要一来就想着怎么优化，暴力解总比做不出好
	// todo 每次把一个数字的位置挪正确，
	// 3，2，4，1
	// 先把4挪到最后 ==> 从4开始翻转【4，2，3，1】，之后在翻转前4个【1，3，2，4】
	// 后面在处理3  ==>  从3开始翻转【3，1，2，4】，之后在翻转前3个【2，1，3，4】
	// 。。。
	indexMap := map[int]int{} // num:index
	target := make([]int, 0, len(arr))
	target = append(target, arr...)
	sort.Slice(target, func(i, j int) bool {
		return target[i] < target[j]
	})
	for i := 0; i < len(arr); i++ {
		indexMap[arr[i]] = i
	}

	res := make([]int, 0)
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == target[i] {
			continue
		}
		targetIndex := indexMap[target[i]]
		res = append(res, targetIndex+1)
		res = append(res, i+1)
		reverse(arr, targetIndex)
		reverse(arr, i)
		for j := 0; j < i; j++ {
			indexMap[arr[j]] = j
		}
	}

	return res
}

func reverse(arr []int, index int) {
	left := 0
	right := index
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	pancakeSort([]int{3, 2, 4, 1})
}
