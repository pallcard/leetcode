package main

// leetcode submit region begin(Prohibit modification and deletion)
type LRUCache struct {
	kv       map[int]int //k:v
	k        []int
	capacity int //容量
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		kv:       map[int]int{},
		k:        make([]int, capacity, capacity),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.kv[key]; ok {
		// get了，则把key移动到首位
		keyIndex := 0
		for i := 0; i < len(this.k); i++ {
			if key == this.k[i] {
				keyIndex = i
				continue
			}
		}
		copy(this.k[1:keyIndex+1], this.k[0:keyIndex])
		//for i := keyIndex - 1; i >= 0; i-- {
		//	this.KArr[i+1] = this.KArr[i]
		//}
		this.k[0] = key
		return v
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if len(this.k) == 0 {
		this.k[0] = key
		this.kv[key] = value
		return
	}

	if this.Get(key) != -1 {
		this.kv[key] = value
		return
	}

	delete(this.kv, this.k[len(this.k)-1])
	copy(this.k[1:this.capacity], this.k[0:this.capacity-1])

	this.kv[key] = value
	this.k[0] = key
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
//leetcode submit region end(Prohibit modification and deletion)

func main() {
	lru := Constructor(2)
	lru.Put(2, 1)
	lru.Put(2, 2)
	lru.Get(2)
	lru.Put(1, 1)
	lru.Put(1, 4)
	lru.Get(2)

}
