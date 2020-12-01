package latest_x_used

// heap version of lfu

// import "container/heap"

// type LFUItem struct {
// 	Value    int
// 	Key      int
// 	Usecount int
// }

// type LFUHeap []LFUItem

// func (h LFUHeap) Less(i, j int) bool {
// 	return h[i].Usecount > h[j].Usecount
// }

// func (h LFUHeap) Swap(i, j int) {
// 	h[i], h[j] = h[j], h[i]
// }

// func (h LFUHeap) Len() int {
// 	return len(h)
// }

// func (h *LFUHeap) Push(x interface{}) {
// 	// Push and Pop use pointer receivers because they modify the slice's length,
// 	// not just its contents.
// 	*h = append(*h, x.(LFUItem))
// }

// func (h *LFUHeap) Pop() interface{} {
// 	old := *h
// 	n := len(old)
// 	x := old[n-1]
// 	*h = old[0 : n-1]
// 	return x
// }

// type LFUCache struct {
// 	usedMap  map[int]*LFUItem
// 	Capacity int
// 	heap     LFUHeap
// }

// func Constructor(capacity int) LFUCache {
// 	c := LFUCache{}
// 	c.usedMap = make(map[int]*LFUItem)
// 	c.heap = make(LFUHeap, 0, capacity)
// 	c.Capacity = capacity
// 	return c
// }

// func (this *LFUCache) Get(key int) int {
// 	if val, ok := this.usedMap[key]; !ok {
// 		return -1
// 	} else {
// 		this.usedMap[key].Usecount++
// 		heap.Fix(&this.heap, key)
// 		return val.Value
// 	}
// }

// func (this *LFUCache) Put(key int, value int) {
// 	newItem := LFUItem{Key: key, Value: value, Usecount: 1}
// 	if len(this.usedMap) < this.Capacity {
// 		this.usedMap[key] = &newItem
// 		heap.Push(&this.heap, newItem)
// 	} else {
// 		item := heap.Pop(&this.heap).(LFUItem)
// 		delete(this.usedMap, item.Key)
// 		heap.Push(&this.heap, newItem)
// 	}
// }

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
