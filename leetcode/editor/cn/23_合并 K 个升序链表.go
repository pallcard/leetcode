package main

import (
	"container/heap"
	"fmt"
)

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type PriorityQueue []*ListNode

func (pq *PriorityQueue) Len() int {
	return len(*pq)
}

func (pq *PriorityQueue) Less(i, j int) bool {
	return (*pq)[i].Val < (*pq)[j].Val
}
func (pq *PriorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*ListNode))
}

func (pq *PriorityQueue) Pop() interface{} {
	x := (*pq)[pq.Len()-1]
	*pq = (*pq)[:pq.Len()-1]
	return x
}

// 使用堆
func mergeKLists(lists []*ListNode) *ListNode {
	virtualHead := &ListNode{}
	p := virtualHead

	h := &PriorityQueue{}
	heap.Init(h)

	for _, list := range lists {
		if list == nil {
			continue
		}
		heap.Push(h, list)
	}

	for h.Len() > 0 {
		minP := heap.Pop(h).(*ListNode)
		p.Next = minP
		p = p.Next
		minP = minP.Next
		if minP != nil {
			heap.Push(h, minP)
		}
		p.Next = nil
	}

	return virtualHead.Next
}

// 直接遍历
func mergeKLists2(lists []*ListNode) *ListNode {
	virthalHead := &ListNode{}
	p := virthalHead

	for {
		var minP *ListNode
		minIndex := -1
		for i, list := range lists {
			if list == nil {
				continue
			}
			if minP == nil {
				minP = list
				minIndex = i
			} else {
				if list.Val < minP.Val {
					minP = list
					minIndex = i
				}
			}
		}
		if minP == nil {
			break
		}
		p.Next = minP
		p = p.Next
		lists[minIndex] = lists[minIndex].Next
		p.Next = nil
	}

	return virthalHead.Next
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	list1 := genListByArr([]int{1, 4, 5})
	list2 := genListByArr([]int{1, 3, 4})
	list3 := genListByArr([]int{2, 6})

	lists := mergeKLists2([]*ListNode{list1, list2, list3})
	fmt.Println(lists)
}

func genListByArr(arr []int) *ListNode {
	virtualHead := &ListNode{}
	p := virtualHead
	for _, v := range arr {
		p.Next = &ListNode{
			Val:  v,
			Next: nil,
		}
		p = p.Next
	}
	return virtualHead.Next
}
