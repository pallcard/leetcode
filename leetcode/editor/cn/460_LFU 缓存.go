package main

// leetcode submit region begin(Prohibit modification and deletion)
type LFUCache struct {
	kv       map[int]int //k:v
	k        []int       //k
	kf       map[int]int //k:f
	capacity int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		kv:       map[int]int{},
		k:        make([]int, capacity, capacity),
		kf:       map[int]int{},
		capacity: capacity,
	}
}

func (this *LFUCache) Get(key int) int {
	if _, ok := this.kv[key]; !ok {
		return -1
	}

	this.handleKey(key)
	return this.kv[key]
}

func (this *LFUCache) handleKey(key int) {
	this.kf[key]++

	mIndex := 0
	keyIndex := 0
	for i := 0; i < len(this.k); i++ { //根据频率找到需要插入的位置mIndex
		if this.kf[this.k[i]] <= this.kf[key] {
			mIndex = i
			break
		}
	}
	for i := mIndex; i < len(this.k); i++ { //找到key的位置keyIndex
		if this.k[i] == key {
			keyIndex = i
			break
		}
	}
	// 把[mIndex,keyIndex)整体位移一位
	copy(this.k[mIndex+1:keyIndex+1], this.k[mIndex:keyIndex])
	this.k[mIndex] = key
}

func (this *LFUCache) Put(key int, value int) {
	if len(this.kv) == 0 {
		this.kv[key] = value
		this.kf[key] = 1
		this.k[0] = key
		return
	}

	// 缓存已满 && 插入不存在的key， 删除过期key
	_, ok := this.kf[key]
	if !ok && len(this.kf) >= this.capacity {
		dKey := this.k[len(this.k)-1]
		delete(this.kv, dKey)
		delete(this.kf, dKey)
	}
	// 把元素插入到key的尾部
	if len(this.kf) < this.capacity {
		this.k[len(this.kf)] = key
	}
	this.kv[key] = value
	this.handleKey(key)
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
//leetcode submit region end(Prohibit modification and deletion)

func main() {
	lfu := Constructor(2)
	lfu.Put(1, 1)
	lfu.Put(2, 2)
	lfu.Get(1)
	lfu.Put(3, 3)
	lfu.Get(2)
}
