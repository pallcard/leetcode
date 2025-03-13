package main

import "fmt"

// leetcode submit region begin(Prohibit modification and deletion)
func findOrder(numCourses int, prerequisites [][]int) []int {
	//todo 后续遍历

	graph := make([][]int, numCourses, numCourses)
	for i := 0; i < numCourses; i++ {
		graph[i] = make([]int, 0)
	}

	for i := 0; i < len(prerequisites); i++ {
		form, to := prerequisites[i][1], prerequisites[i][0]
		graph[form] = append(graph[form], to)
	}

	visited := make([]bool, numCourses, numCourses)
	cycle := false

	var dfs func(course int, pathMap []bool, path *[]int)
	dfs = func(course int, pathMap []bool, path *[]int) {
		if cycle {
			return
		}

		if pathMap[course] {
			cycle = true
			return
		}

		if visited[course] { //todo 有visited限制，path不会加重复
			return
		}

		visited[course] = true
		pathMap[course] = true
		for i := 0; i < len(graph[course]); i++ {
			dfs(graph[course][i], pathMap, path)
		}
		pathMap[course] = false
		*path = append(*path, course)

	}
	pathMap := make([]bool, numCourses, numCourses)
	postorder := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		dfs(i, pathMap, &postorder)
	}

	if cycle {
		return []int{}
	}

	// todo 反转path
	left := 0
	right := len(postorder) - 1
	for left < right {
		postorder[left], postorder[right] = postorder[right], postorder[left]
		left++
		right--
	}
	return postorder
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	res := findOrder(2, [][]int{{1, 0}})
	fmt.Print(res)
}
