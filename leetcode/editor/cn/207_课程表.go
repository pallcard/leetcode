package main

import "fmt"

// leetcode submit region begin(Prohibit modification and deletion)
func canFinish(numCourses int, prerequisites [][]int) bool {
	// req[1] -> req[0]
	// todo 图的dfs遍历

	// 图转化为邻接表
	graph := make([][]int, numCourses, numCourses)
	for i := 0; i < numCourses; i++ {
		graph[i] = make([]int, 0)
	}
	for i := 0; i < len(prerequisites); i++ {
		form, to := prerequisites[i][1], prerequisites[i][0]
		graph[form] = append(graph[form], to)
	}

	cycle := false
	visited := make([]bool, numCourses, numCourses)

	var dfs func(graph [][]int, course int, path []bool)
	dfs = func(graph [][]int, course int, path []bool) {
		//if graph[]
		if cycle {
			return
		}

		if path[course] {
			cycle = true
			return
		}

		if visited[course] {
			return
		}

		visited[course] = true // 能走到course这个节点，后续无需再走
		path[course] = true
		for i := 0; i < len(graph[course]); i++ {
			dfs(graph, graph[course][i], path)
		}
		path[course] = false
	}

	path := make([]bool, numCourses, numCourses)

	for i := 0; i < numCourses; i++ {
		//path := make([]bool, numCourses, numCourses)
		dfs(graph, i, path)
	}
	return !cycle

}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	res := canFinish(2, [][]int{{1, 0}})
	fmt.Print(res)
}
