package main

// leetcode submit region begin(Prohibit modification and deletion)
func allPathsSourceTarget(graph [][]int) [][]int {
	n := len(graph) // 节点数
	res := make([][]int, 0)

	var dfs func(trace []int, cur int)

	dfs = func(trace []int, cur int) {
		if cur == n-1 {
			temp := make([]int, len(trace), len(trace))
			copy(temp, trace)
			res = append(res, temp)
		}

		for i := 0; i < len(graph[cur]); i++ {
			trace = append(trace, graph[cur][i])
			dfs(trace, graph[cur][i])
			trace = trace[:len(trace)-1]
		}

	}

	dfs([]int{0}, 0)
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
