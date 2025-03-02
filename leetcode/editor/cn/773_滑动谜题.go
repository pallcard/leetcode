package main

import (
	"fmt"
	"strings"
)

// leetcode submit region begin(Prohibit modification and deletion)
func slidingPuzzle2(board [][]int) int {

	visited := map[string]bool{}
	queue := make([][][]int, 0)
	queue = append(queue, board)
	visited[board2Str(board)] = true

	target := [][]int{{1, 2, 3}, {4, 5, 0}}

	step := 0
	for len(queue) > 0 {

		levelCnt := len(queue)
		for i := 0; i < levelCnt; i++ {
			front := queue[0]
			queue = queue[1:]

			if board2Str(front) == board2Str(target) {
				return step
			}

			for _, b := range getNext773(front) {
				if !visited[board2Str(b)] {
					queue = append(queue, b)
				}
			}

		}
		step++
	}
	return step

}

func board2Str(board [][]int) string {
	str := ""
	for _, row := range board {
		for _, num := range row {
			str += fmt.Sprintf("%d", num)
		}
	}
	return str
}

func getNext773(board [][]int) [][][]int {

	indexI := 0
	indexJ := 0
out:
	for i, row := range board {
		for j, num := range row {
			if num == 0 {
				indexI = i
				indexJ = j
				break out
			}
		}
	}

	res := make([][][]int, 0, 3)

	if indexI >= 1 {
		tempBoard := copyArr(board)
		tempBoard[indexI][indexJ], tempBoard[indexI-1][indexJ] = tempBoard[indexI-1][indexJ], tempBoard[indexI][indexJ]
		res = append(res, tempBoard)
	}
	if indexJ >= 1 {
		tempBoard := copyArr(board)
		tempBoard[indexI][indexJ], tempBoard[indexI][indexJ-1] = tempBoard[indexI][indexJ-1], tempBoard[indexI][indexJ]
		res = append(res, tempBoard)
	}

	if indexI+1 < len(board) {
		tempBoard := copyArr(board)
		tempBoard[indexI][indexJ], tempBoard[indexI+1][indexJ] = tempBoard[indexI+1][indexJ], tempBoard[indexI][indexJ]
		res = append(res, tempBoard)
	}

	if indexJ+1 < len(board[0]) {
		tempBoard := copyArr(board)
		tempBoard[indexI][indexJ], tempBoard[indexI][indexJ+1] = tempBoard[indexI][indexJ+1], tempBoard[indexI][indexJ]
		res = append(res, tempBoard)
	}

	return res
}

func slidingPuzzle(board [][]int) int {
	visited := map[string]bool{}

	boardStr := board2Str(board)
	queue := make([]string, 0)
	queue = append(queue, boardStr)
	visited[boardStr] = true

	target := "123450"

	step := 0
	for len(queue) > 0 {

		levelCnt := len(queue)
		for i := 0; i < levelCnt; i++ {
			front := queue[0]
			queue = queue[1:]

			if front == target {
				return step
			}

			for _, b := range getNext773_3(front) {
				if !visited[b] {
					queue = append(queue, b)
					visited[b] = true
				}
			}

		}
		step++
	}
	return -1
}

func getNext773_3(board string) []string {

	index := strings.Index(board, "0")
	// 0 1 2
	// 3 4 5
	swapIndex := [][]int{
		{1, 3},
		{0, 2, 4},
		{1, 5},
		{0, 4},
		{1, 3, 5},
		{2, 4},
	}

	res := make([]string, 0, 3)
	for _, si := range swapIndex[index] {
		tempBoard := []byte(board)
		// i,j  i-1,j
		tempBoard[index], tempBoard[si] = tempBoard[si], tempBoard[index]
		res = append(res, string(tempBoard))
	}
	return res

}

func getNext773_2(board string) []string {

	index0 := 0
	for i, num := range board {
		if num == '0' {
			index0 = i
			break
		}
	}

	// board 为2*3
	indexI := 0
	indexJ := index0
	if index0 >= 3 {
		indexI = 1
		indexJ = indexJ - 3
	}

	res := make([]string, 0, 3)

	if indexI >= 1 {
		tempBoard := []byte(board)
		// i,j  i-1,j
		tempBoard[indexJ+indexI*3], tempBoard[indexJ+(indexI-1)*3] = tempBoard[indexJ+(indexI-1)*3], tempBoard[indexJ+indexI*3]
		res = append(res, string(tempBoard))
	}
	if indexJ >= 1 {
		tempBoard := []byte(board)
		// i,j i,j-1
		tempBoard[indexJ+indexI*3], tempBoard[(indexJ-1)+indexI*3] = tempBoard[(indexJ-1)+indexI*3], tempBoard[indexJ+indexI*3]
		res = append(res, string(tempBoard))
	}

	if indexI+1 < 2 {
		tempBoard := []byte(board)
		// i,j  i+1,j
		tempBoard[indexJ+indexI*3], tempBoard[indexJ+(indexI+1)*3] = tempBoard[indexJ+(indexI+1)*3], tempBoard[indexJ+indexI*3]
		res = append(res, string(tempBoard))
	}

	if indexJ+1 < 3 {
		tempBoard := []byte(board)
		// i,j  i,j+1
		tempBoard[indexJ+indexI*3], tempBoard[(indexJ+1)+indexI*3] = tempBoard[(indexJ+1)+indexI*3], tempBoard[indexJ+indexI*3]
		res = append(res, string(tempBoard))
	}

	return res
}

func copyArr(board [][]int) [][]int {
	arr := make([][]int, 0, 2)
	for i, row := range board {
		arr = append(arr, []int{})
		for _, num := range row {
			arr[i] = append(arr[i], num)
		}
	}
	return arr
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	res := slidingPuzzle([][]int{{1, 2, 3}, {5, 4, 0}})
	fmt.Println(res)
}
