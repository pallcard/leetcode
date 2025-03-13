package main

// leetcode submit region begin(Prohibit modification and deletion)
type LRUCache struct {
	KV       map[int]int //k:v
	KArr     []int
	Capacity int //容量
	Count    int //当前数量
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		KV:       map[int]int{},
		KArr:     make([]int, capacity, capacity),
		Capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.KV[key]; ok {
		keyIndex := 0
		for i := this.Count - 1; i >= 0; i-- {
			if key == this.KArr[i] {
				keyIndex = i
				continue
			}
		}
		for i := keyIndex - 1; i >= 0; i-- {
			this.KArr[i+1] = this.KArr[i]
		}
		this.KArr[0] = key
		return v
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if this.Get(key) != -1 {
		this.KV[key] = value
		return
	}
	for i := this.Count - 1; i >= 0; i-- {
		if i == this.Capacity-1 { //容量已满删除最后一个key
			delete(this.KV, this.KArr[i])
			continue
		}
		this.KArr[i+1] = this.KArr[i]
	}
	this.KV[key] = value
	this.KArr[0] = key
	if this.Count != this.Capacity {
		this.Count++
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
