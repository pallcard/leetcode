package main

import "fmt"

// leetcode submit region begin(Prohibit modification and deletion)

func maxAreaOfIsland(grid [][]int) int {
	visited := make([][]bool, len(grid), len(grid))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(grid[i]), len(grid[i]))
	}

	maxArea := 0
	curCount := 0

	direct := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	var dfs func(curX, curY int)
	dfs = func(curX, curY int) {
		if visited[curX][curY] {
			return
		}
		visited[curX][curY] = true
		curCount++ //未访问过才能加1

		for i := 0; i < 4; i++ {
			nextX := curX + direct[i][0]
			nextY := curY + direct[i][1]
			if nextX >= 0 && nextX < len(grid) &&
				nextY >= 0 && nextY < len(grid[0]) &&
				grid[nextX][nextY] == 1 {

				dfs(nextX, nextY)
			}
		}

	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				curCount = 0
				dfs(i, j)
				if curCount > maxArea {
					maxArea = curCount
				}
			}
		}
	}

	return maxArea
}

func maxAreaOfIsland2(grid [][]int) int {
	visited := make([][]bool, len(grid), len(grid))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(grid[i]), len(grid[i]))
	}

	maxArea := 0

	direct := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	var dfs func(curX, curY int, curArea int)
	dfs = func(curX, curY int, curArea int) {
		if visited[curX][curY] {
			return
		}
		visited[curX][curY] = true

		if curArea > maxArea {
			maxArea = curArea
		}

		for i := 0; i < 4; i++ {
			nextX := curX + direct[i][0]
			nextY := curY + direct[i][1]
			if nextX >= 0 && nextX < len(grid) &&
				nextY >= 0 && nextY < len(grid[0]) &&
				grid[nextX][nextY] == 1 {
				dfs(nextX, nextY, curArea+1) //todo 写法错误，这样回溯的时候面积掉了
			}
		}

	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				dfs(i, j, 1)
			}
		}
	}

	return maxArea
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	area := maxAreaOfIsland([][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}})
	fmt.Print(area)
}
