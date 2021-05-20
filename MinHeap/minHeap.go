package MinHeap

type minHeap struct { // 定义 minHeap 结构体
	k    int   // k 容量
	heap []int // heap 数组
}

func createMinHeap(k int, nums []int) *minHeap { // 创建 minHeap
	heap := &minHeap{k: k, heap: []int{}}
	for _, n := range nums { // 把 nums 的数字放进去初始化
		heap.add(n)
	}
	return heap // 返回出初始化好的 heap
}

func (this *minHeap) add(num int) { // minHeap 绑定 add 方法
	if len(this.heap) < this.k { // heap 数组长度还不够 k
		this.heap = append(this.heap, num) // 将num加入heap数组
		this.up(len(this.heap) - 1)        // 执行上浮，上浮到合适的位置
	} else if num > this.heap[0] { // 如果num比堆顶数字要大
		this.heap[0] = num // 堆顶 换人
		this.down(0)       // 执行下沉，下沉到合适的位置
	} // 其他情况 不加入
}

func (this *minHeap) up(i int) { // 位置i上的元素，上浮到合适位置
	for i > 0 { // 上浮到索引0就停止上浮，0是堆顶位置
		parent := (i - 1) >> 1                // 找到父节点在heap数组中的位置
		if this.heap[parent] > this.heap[i] { // 如果父节点的数字比插入的数字大
			this.heap[parent], this.heap[i] = this.heap[i], this.heap[parent] // 交换
			i = parent                                                        // 更新 i
		} else { // 父比自己小，满足最小堆的性质，break
			break
		}
	}
}

func (this *minHeap) down(i int) { // 下沉到合适的位置
	for 2*i+1 < len(this.heap) { // 左子节点的索引如果已经越界，终止下沉
		child := 2*i + 1 // 左子节点在heap数组中的位置
		if child+1 < len(this.heap) && this.heap[child+1] < this.heap[child] {
			child++ // 如果右子节点存在且值更小，则用它，去比较
		}
		if this.heap[i] > this.heap[child] { // 如果插入的数字比子节点都大
			this.heap[child], this.heap[i] = this.heap[i], this.heap[child] // 交换
			i = child                                                       // 更新 i
		} else { // 子比自己大，满足最小堆的属性，break
			break
		}
	}
}

type KthLargest struct { // KthLargest 结构体
	heap *minHeap
}

func Constructor(k int, nums []int) KthLargest { // 创建 KthLargest
	return KthLargest{heap: createMinHeap(k, nums)}
}

func (this *KthLargest) Add(val int) int { // 执行加入操作，返回第 k 大数字
	this.heap.add(val)
	return this.heap.heap[0] // 堆顶即第 k 大数字
}
