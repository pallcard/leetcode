package main

import (
	"fmt"
	"math"
)

// leetcode submit region begin(Prohibit modification and deletion)

func openLock2(deadends []string, target string) int {
	// 记录需要跳过的死亡密码
	deads := make(map[string]struct{})
	for _, s := range deadends {
		deads[s] = struct{}{}
	}
	if _, found := deads["0000"]; found {
		return -1
	}

	// 记录已经穷举过的密码，防止走回头路
	visited := make(map[string]struct{})
	q := make([]string, 0)
	// 从起点开始启动广度优先搜索
	step := 0
	q = append(q, "0000")
	visited["0000"] = struct{}{}

	for len(q) > 0 {
		sz := len(q)
		// 将当前队列中的所有节点向周围扩散
		for i := 0; i < sz; i++ {
			cur := q[0]
			q = q[1:]

			// 判断是否到达终点
			if cur == target {
				return step
			}

			// 将一个节点的合法相邻节点加入队列
			for _, neighbor := range getNeighbors(cur) {
				if _, found := visited[neighbor]; !found {
					if _, dead := deads[neighbor]; !dead {
						q = append(q, neighbor)
						visited[neighbor] = struct{}{}
					}
				}
			}
		}
		// 在这里增加步数
		step++
	}
	// 如果穷举完都没找到目标密码，那就是找不到了
	return -1
}

// 将 s[j] 向上拨动一次
func plusOne(s string, j int) string {
	ch := []rune(s)
	if ch[j] == '9' {
		ch[j] = '0'
	} else {
		ch[j]++
	}
	return string(ch)
}

// 将 s[i] 向下拨动一次
func minusOne(s string, j int) string {
	ch := []rune(s)
	if ch[j] == '0' {
		ch[j] = '9'
	} else {
		ch[j]--
	}
	return string(ch)
}

// 将 s 的每一位向上拨动一次或向下拨动一次，8 种相邻密码
func getNeighbors(s string) []string {
	neighbors := make([]string, 0)
	for i := 0; i < 4; i++ {
		neighbors = append(neighbors, plusOne(s, i))
		neighbors = append(neighbors, minusOne(s, i))
	}
	return neighbors
}

func openLock(deadends []string, target string) int {
	// 0000 ---> target

	deadMap := map[string]bool{}
	for _, deadend := range deadends {
		deadMap[deadend] = true
	}
	if deadMap["0000"] {
		return -1
	}

	visited := map[string]bool{"0000": true}
	queue := []string{"0000"}
	step := 0

	for len(queue) > 0 {
		levelCnt := len(queue)
		for i := 0; i < levelCnt; i++ {
			front := queue[0]
			queue = queue[1:] //出队
			if target == front {
				return step
			}
			for _, s := range getNeighbors(front) {
				if !visited[s] && !deadMap[s] {
					queue = append(queue, s)
					visited[s] = true
				}
			}
		}
		step++
	}
	return -1
}

func getNext(cur string) []string {
	next := make([]string, 0)
	for i := 0; i < 4; i++ {
		temp := []byte(cur)
		if temp[i] == '9' {
			temp[i] = '0'
		} else {
			temp[i] = temp[i] + 1
		}
		next = append(next, string(temp))
		temp = []byte(cur)
		if temp[i%4] == '0' {
			temp[i%4] = '9'
		} else {
			temp[i%4] = temp[i%4] - 1
		}
		next = append(next, string(temp))
	}
	return next
}

func backtrace(memoMap map[string]int, current, target []rune,
	step int) int {
	if memo, ok := memoMap[string(current)]; ok {
		return memo
	}
	memoMap[string(current)] = step
	if string(current) == string(target) {
		return step
	}

	minStep := math.MaxInt
	for i := 0; i < 8; i++ {
		if i <= 3 {
			if current[i] == '9' {
				current[i] = '0'
			} else {
				current[i] = current[i] + 1
			}

		} else {
			if current[i] == '0' {
				current[i] = '9'
			} else {
				current[i] = current[i] - 1
			}
		}
		res := backtrace(memoMap, current, target, step+1)
		if res > 0 {
			if res < minStep {
				minStep = res
			}
			return minStep
		}

		if i <= 3 {
			if current[i] == '0' {
				current[i] = '9'
			} else {
				current[i] = current[i] - 1
			}
		} else {
			if current[i] == '9' {
				current[i] = '0'
			} else {
				current[i] = current[i] + 1
			}
		}
	}

	return -1

}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	lock := openLock([]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888")
	fmt.Print(lock)
}
