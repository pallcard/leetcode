package main

import (
	"fmt"
	"sort"
)

//leetcode submit region begin(Prohibit modification and deletion)

type ExamRoom struct {
	//SeatArr   []int       //座位
	Seated []int //已坐的位置
	Length int
	//SeatedMap map[int]int //seated:index
}

func Constructor(n int) ExamRoom {
	return ExamRoom{
		Seated: []int{},
		Length: n,
	}
}
func (this *ExamRoom) Seat() int {
	// todo 左右边界需要特殊处理，
	// todo 左边界是直接第一个点到0的距离，右边界是最后一个点到n-1的距离，中间点是 i+1到i的距离
	// 使用库函数 sort.SearchInts([]int, p), 返回插入的数据
	if len(this.Seated) == 0 { //仅有一个座位时
		this.Seated = append(this.Seated, 0)
		return 0
	}
	// 左边界距离
	maxDis := this.Seated[0] - 0
	point := 0

	for i := 1; i < len(this.Seated); i++ { //中间的位置
		dis := (this.Seated[i] - this.Seated[i-1]) / 2
		if dis > maxDis {
			maxDis = dis
			point = this.Seated[i-1] + (this.Seated[i]-this.Seated[i-1])/2
		}
	}

	// 右边界距离
	if (this.Length-1)-this.Seated[len(this.Seated)-1] > maxDis {
		maxDis = (this.Length - 1) - this.Seated[len(this.Seated)-1]
		point = this.Length - 1
	}

	pIndex := sort.SearchInts(this.Seated, point) //二分查找插入的位置
	if pIndex < this.Length {
		this.Seated = append(this.Seated, 0)
		copy(this.Seated[pIndex+1:], this.Seated[pIndex:])
		this.Seated[pIndex] = point
	}
	return point
}

func (this *ExamRoom) Leave(point int) {
	pIndex := sort.SearchInts(this.Seated, point)
	copy(this.Seated[pIndex:], this.Seated[pIndex+1:])
	this.Seated = this.Seated[:len(this.Seated)-1]
}

func (this *ExamRoom) Seat2() int {

	if len(this.Seated) == 0 {
		this.Seated = append(this.Seated, 0)
		//this.SeatArr[0] = 1
		return 0
	} else {
		sort.Ints(this.Seated)
		maxDis := -1
		//第一个座位为空
		if this.Seated[0] != 0 {
			maxDis = (this.Seated[0] - 0) //第一个到-1的距离
		}

		maxIndex := -1
		maxEnd := 0
		for i := 0; i < len(this.Seated); i++ {
			if i == len(this.Seated)-1 { //最后一个到N的距离
				if (this.Length - 1 - this.Seated[i]) > maxDis {
					maxDis = (this.Length - 1 - this.Seated[i])
					maxIndex = i
					maxEnd = this.Length
				}
			} else {
				if (this.Seated[i+1]-this.Seated[i])/2 > maxDis {
					maxDis = (this.Seated[i+1] - this.Seated[i]) / 2
					maxIndex = i
					maxEnd = i + 1
				}
			}
		}

		seat := -1
		if maxEnd == this.Length {
			seat = this.Length - 1
		} else if maxIndex == -1 {
			seat = 0
		} else {
			seat = this.Seated[maxIndex] + maxDis
		}
		this.Seated = append(this.Seated, seat)
		//this.SeatArr[seat] = 1
		return seat
	}

	//
	//pop := heap.Pop(this.Heap).([]int)
	//seat := -1
	//if pop[0] == -1 {
	//	seat = 0
	//} else if pop[1] == this.Length {
	//	seat = this.Length - 1
	//} else {
	//	seat = pop[0] + (pop[1]-pop[0])/2
	//}
	//heap.Push(this.Heap, []int{pop[0], seat})
	//heap.Push(this.Heap, []int{seat, pop[1]})
	//this.StartP[seat] = []int{pop[0], seat}
	//this.EndP[seat] = []int{seat, pop[1]}
	//return seat
}

func (this *ExamRoom) Leave2(p int) {
	index := 0
	for i := 0; i < len(this.Seated); i++ {
		if this.Seated[i] == p {
			index = i
		}
	}

	this.Seated = append(this.Seated[0:index], this.Seated[index+1:]...)
	//this.Seated = append(this.Seated, seat)
}

/**
 * Your ExamRoom object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Seat();
 * obj.Leave(p);
 */
//leetcode submit region end(Prohibit modification and deletion)

func main() {
	ints := sort.SearchInts([]int{1, 2, 3}, 4)
	fmt.Print(ints)

	arr := make([]int, 0)
	arr = append(arr, 1)

	obj := Constructor(8)
	obj.Seat()
	obj.Seat()
	obj.Seat()
	obj.Leave(0)
	obj.Leave(7)
	obj.Seat()
	obj.Seat()
	obj.Seat()
	obj.Seat()
	obj.Seat()
	obj.Seat()
	obj.Seat()
	obj.Seat()
	obj.Seat()
	obj.Leave(0)
}
