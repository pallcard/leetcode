package main

import "fmt"

/*
todo
 1.  left, right := 0, len(scores)   ==> [left, right) 左开右闭
 2.  for left < right {}             ==> 与1对应，可取全数据
 3.  scores[mid] >= target -->  right = mid     ====> 保证target在区间内
     scores[mid] < target  -->  left = mid + 1  ====> 比target小的排出区间
--------
todo
 1.  left, right := 0, len(scores)   ==> [left, right] 左闭右闭
 2.  for left <= right {}            ==> 与1对应，可取全数据
 3.  scores[mid] > target  -->  right = mid + 1     ====> 比target大在排出区间
     scores[mid] < target  -->  left = mid + 1     ====> 比target小的排出区间
     scores[mid] = target  -->  目标值
*/

// leetcode submit region begin(Prohibit modification and deletion)
func countTarget2(scores []int, target int) int {
	left, right := 0, len(scores)-1
	for left <= right {
		mid := left + (right-left)/2
		if scores[mid] > target {
			right = mid - 1
		} else if scores[mid] < target {
			left = mid + 1
		} else {
			cnt := 1
			for i := mid - 1; i >= 0; i-- {
				if scores[i] == target {
					cnt++
				} else {
					break
				}
			}

			for i := mid + 1; i < len(scores); i++ {
				if scores[i] == target {
					cnt++
				} else {
					break
				}
			}
			return cnt
		}

	}

	return 0
}

func countTarget(scores []int, target int) int {
	left, right := 0, len(scores)

	for left < right {
		mid := left + (right-left)/2

		if scores[mid] >= target {
			right = mid
		} else if scores[mid] < target {
			left = mid + 1
		}
	}

	cnt := 0
	for ; left < len(scores); left++ {
		if scores[left] == target {
			cnt++
		}
	}

	return cnt
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	cnt := countTarget2([]int{1, 2, 2, 3}, 2)
	fmt.Print(cnt)
}
